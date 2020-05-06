package inbound

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

type configuration struct {
	Endpoint string `json:"endpoint"`
	Port     string `json:"port"`
}

type ParsedEmail struct {
	Headers			map[string]string
	Body			map[string]string
	Attachments		map[string][]byte
	rawRequest		*http.Request
}

func Parse(response http.ResponseWriter, request *http.Request) *ParsedEmail {
	result := ParsedEmail{
		Headers:		make(map[string]string),
		Body:		    make(map[string]string),
		Attachments: 	make(map[string][]byte),
		rawRequest:     request,
	}
	result.parse(response)
	return &result
}

func (email *ParsedEmail) parse(resp http.ResponseWriter) {
	err := email.rawRequest.ParseMultipartForm(0)
	if err != nil {
		log.Fatal(err)
	}
	emails := email.rawRequest.MultipartForm.Value["email"]
	headers := email.rawRequest.MultipartForm.Value["headers"]
	if len(headers) > 0  {
		email.parseHeaders(headers[0])
	}
	if len(emails) > 0 {
		email.parseRawEmail(emails[0])
	}
	resp.WriteHeader(http.StatusOK)
}

func (email *ParsedEmail) parseRawEmail(rawEmail string) {
	sections := strings.SplitN(rawEmail, "\n\n", 2)
	email.parseHeaders(sections[0])
	raw := parseMultipart(strings.NewReader(sections[1]), email.Headers["Content-Type"])
	for {
		emailPart, err := raw.NextPart()
		if err == io.EOF {
			return
		}
		rawEmailBody := parseMultipart(emailPart, emailPart.Header.Get("Content-Type"))
		if rawEmailBody != nil {
			for {
				emailBodyPart, err := rawEmailBody.NextPart()
				if err == io.EOF {
					break
				}
				header := emailBodyPart.Header.Get("Content-Type")
				email.Body[header] = string(readBody(emailBodyPart))
			}

		} else if emailPart.FileName() != "" {
			email.Attachments[emailPart.FileName()] = readBody(emailPart)
		} else {
			header := emailPart.Header.Get("Content-Type")
			email.Body[header] = string(readBody(emailPart))
		}
	}
}

func parseMultipart(body io.Reader, contentType string) *multipart.Reader {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		return multipart.NewReader(body, params["boundary"])
	}
	return nil
}

func readBody(body io.Reader) []byte {
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	return raw
}

func (email *ParsedEmail) parseHeaders(headers string) {
	splitHeaders := strings.Split(strings.TrimSpace(headers), "\n")
	for _, header := range splitHeaders {
		splitHeader := strings.SplitN(header, ": ", 2)
		email.Headers[splitHeader[0]] = splitHeader[1]
	}
}

func loadConfig(path string) configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}
	var conf configuration
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}
	return conf
}
