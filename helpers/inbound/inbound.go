package inbound

import (
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

type ParsedEmail struct {
	Headers     map[string]string
	Body        map[string]string
	Attachments map[string][]byte
	rawRequest  *http.Request
}

func Parse(request *http.Request) (*ParsedEmail, error) {
	result := ParsedEmail{
		Headers:     make(map[string]string),
		Body:        make(map[string]string),
		Attachments: make(map[string][]byte),
		rawRequest:  request,
	}
	err := result.parse()
	return &result, err
}

func (email *ParsedEmail) parse() error {
	err := email.rawRequest.ParseMultipartForm(0)
	if err != nil {
		return err
	}
	emails := email.rawRequest.MultipartForm.Value["email"]
	headers := email.rawRequest.MultipartForm.Value["headers"]
	if len(headers) > 0 {
		email.parseHeaders(headers[0])
	}
	if len(emails) > 0 {
		email.parseRawEmail(emails[0])
	}

	return nil
}

func (email *ParsedEmail) parseRawEmail(rawEmail string) error {
	sections := strings.SplitN(rawEmail, "\n\n", 2)
	email.parseHeaders(sections[0])
	raw, err := parseMultipart(strings.NewReader(sections[1]), email.Headers["Content-Type"])
	if err != nil {
		return err
	}

	for {
		emailPart, err := raw.NextPart()
		if err == io.EOF {
			return nil
		}
		rawEmailBody, err := parseMultipart(emailPart, emailPart.Header.Get("Content-Type"))
		if err != nil {
			return err
		}
		if rawEmailBody != nil {
			for {
				emailBodyPart, err := rawEmailBody.NextPart()
				if err == io.EOF {
					break
				}
				header := emailBodyPart.Header.Get("Content-Type")
				b, err := ioutil.ReadAll(emailPart)
				if err != nil {
					return err
				}
				email.Body[header] = string(b)
			}

		} else if emailPart.FileName() != "" {
			b, err := ioutil.ReadAll(emailPart)
			if err != nil {
				return err
			}
			email.Attachments[emailPart.FileName()] = b
		} else {
			header := emailPart.Header.Get("Content-Type")
			b, err := ioutil.ReadAll(emailPart)
			if err != nil {
				return err
			}
			email.Body[header] = string(b)
		}
	}
}

func parseMultipart(body io.Reader, contentType string) (*multipart.Reader, error) {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		return multipart.NewReader(body, params["boundary"]), nil
	}
	return nil, nil
}

func (email *ParsedEmail) parseHeaders(headers string) {
	splitHeaders := strings.Split(strings.TrimSpace(headers), "\n")
	for _, header := range splitHeaders {
		splitHeader := strings.SplitN(header, ": ", 2)
		email.Headers[splitHeader[0]] = splitHeader[1]
	}
}
