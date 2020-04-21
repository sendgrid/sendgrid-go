package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/sendgrid/rest"
)

type configuration struct {
	Endpoint string `json:"endpoint"`
	Port     string `json:"port"`
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

func indexHandler(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "%s", "Hello World")
	if err != nil {
		log.Fatal(err)
	}
}

func getBoundary(value string, contentType string) (string, *strings.Reader) {
	_, params, _ := mime.ParseMediaType(contentType)
	body := strings.NewReader(value)
	return params["boundary"], body
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

func inboundHandler(response http.ResponseWriter, request *http.Request) {
	reader := parseMultipart(request.Body, request.Header.Get("Content-Type"))
	if reader != nil {
		parsedEmail := make(map[string]string)
		emailHeader := make(map[string]string)
		binaryFiles := make(map[string][]byte)
		parsedRawEmail := make(map[string]string)
		rawFiles := make(map[string]string)
		for {
			requestPart, err := reader.NextPart()
			if err == io.EOF {
				// We have finished parsing, do something with the parsed data
				printMap(parsedEmail, "")
				printMap(emailHeader, "e")

				for key, value := range binaryFiles {
					fmt.Println("bKey:", key, " bValue:", value)
				}

				printMap(parsedRawEmail, "r")
				printMap(rawFiles, "rf")
				// Twilio SendGrid needs a 200 OK response to stop POSTing
				response.WriteHeader(http.StatusOK)
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			body := readBody(requestPart)

			if requestPart.FileName() == "" {
				if requestPart.FormName() == "email" {
					handleRawEmail(requestPart, parsedRawEmail, rawFiles)
				} else if requestPart.FormName() == "headers" {
					handleHeaders(body, emailHeader)
				} else {
					parsedEmail[requestPart.FormName()] = string(body)
				}
			} else {
				// attachment with binary data
				binaryFiles[requestPart.FileName()] = body
			}
		}
	}
}
func handleHeaders(value []byte, emailHeader map[string]string) {
	s := strings.Split(string(value), "\n")
	var a []string
	for _, v := range s {
		t := strings.Split(string(v), ": ")
		a = append(a, t...)
	}
	for i := 0; i < len(a)-1; i += 2 {
		emailHeader[a[i]] = a[i+1]
	}
}

func printMap(inputMap map[string]string, prefix string) {
	for key, value := range inputMap {
		fmt.Println(prefix, "Key:", key, " ", prefix, "Value:", value)
	}
}
func handleRawEmail(email *multipart.Part, parsedRawEmail map[string]string, rawFiles map[string]string) {
	raw := parseMultipart(email, email.Header.Get("Content-Type"))
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
				parsedRawEmail[header] = string(readBody(emailBodyPart))
			}

		} else {
			rawFiles[emailPart.FileName()] = string(readBody(emailPart))
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		// Test Sender
		path := os.Args[1]
		host := os.Args[2]
		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal("Check your Filepath. ", err)
		}
		Headers := map[string]string{
			"User-Agent":   "Twilio-SendGrid-Test",
			"Content-Type": "multipart/form-data; boundary=xYzZY",
		}
		method := rest.Post
		request := rest.Request{
			Method:  method,
			BaseURL: host,
			Headers: Headers,
			Body:    file,
		}
		_, err = rest.Send(request)
		if err != nil {
			log.Fatal("Check your Filepath. ", err)
		}
	} else {
		conf := loadConfig("./conf.json")
		http.HandleFunc("/", indexHandler)
		http.HandleFunc(conf.Endpoint, inboundHandler)
		port := os.Getenv("PORT")
		if port == "" {
			port = conf.Port
		}
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalln("Error")
		}
	}
}
