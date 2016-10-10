package main

import (
	"bufio"
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
	fmt.Fprintf(response, "%s", "Hello World")
}

func getBoundary(value string, contentType string) (string, *strings.Reader) {
	body := strings.NewReader(value)
	bodySplit := strings.Split(string(value), contentType)
	scanner := bufio.NewScanner(strings.NewReader(bodySplit[1]))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		break
	}
	boundary := lines[0][9:]
	return boundary, body
}

func inboundHandler(response http.ResponseWriter, request *http.Request) {
	mediaType, params, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(request.Body, params["boundary"])
		parsedEmail := make(map[string]string)
		emailHeader := make(map[string]string)
		binaryFiles := make(map[string][]byte)
		parsedRawEmail := make(map[string]string)
		rawFiles := make(map[string]string)
		for {
			p, err := mr.NextPart()
			// We have found an attachment with binary data
			if err == nil && p.FileName() != "" {
				contents, err := ioutil.ReadAll(p)
				if err != nil {
					log.Fatal(err)
				}
				binaryFiles[p.FileName()] = contents
			}
			if err == io.EOF {
				// We have finished parsing, do something with the parsed data
				for key, value := range parsedEmail {
					fmt.Println("Key:", key, " Value:", value)
				}
				for key, value := range emailHeader {
					fmt.Println("eKey:", key, " eValue:", value)
				}
				for key, value := range binaryFiles {
					fmt.Println("bKey:", key, " bValue:", value)
				}
				for key, value := range parsedRawEmail {
					fmt.Println("rKey:", key, " rValue:", value)
				}
				for key, value := range rawFiles {
					fmt.Println("rfKey:", key, " rfValue:", value)
				}
				// SendGrid needs a 200 OK response to stop POSTing
				response.WriteHeader(http.StatusOK)
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			value, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}
			header := p.Header.Get("Content-Disposition")
			if strings.Contains(header, "filename") != true {
				header = header[17 : len(header)-1]
				parsedEmail[header] = string(value)
			} else {
				header = header[11:]
				f := strings.Split(header, "=")
				parsedEmail[f[1][1:len(f[1])-11]] = f[2][1 : len(f[2])-1]
			}
			if header == "headers" {
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
			// Since we have parsed the headers, we can delete the original
			delete(parsedEmail, "headers")

			// We have a raw message
			if header == "email" {
				boundary, body := getBoundary(string(value), "Content-Type: multipart/mixed; ")
				raw := multipart.NewReader(body, boundary)
				for {
					next, err := raw.NextPart()
					if err == io.EOF {
						// We have finished parsing
						break
					}
					value, err := ioutil.ReadAll(next)
					if err != nil {
						log.Fatal(err)
					}
					header := next.Header.Get("Content-Type")

					// Parse the headers
					if strings.Contains(header, "multipart/alternative") {
						boundary, body := getBoundary(string(value), "Content-Type: multipart/alternative; ")
						raw := multipart.NewReader(body, boundary)
						for {
							next, err := raw.NextPart()
							if err == io.EOF {
								// We have finished parsing
								break
							}
							value, err := ioutil.ReadAll(next)
							if err != nil {
								log.Fatal(err)
							}
							header := next.Header.Get("Content-Type")
							parsedRawEmail[header] = string(value)
						}
					} else {
						// It's a base64 encoded attachment
						rawFiles[header] = string(value)
					}
				}
			}
			// Since we've parsed this header, we can delete the original
			delete(parsedEmail, "email")
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
		Headers := make(map[string]string)
		Headers["User-Agent"] = "SendGrid-Test"
		Headers["Content-Type"] = "multipart/form-data; boundary=xYzZY"
		method := rest.Post
		request := rest.Request{
			Method:  method,
			BaseURL: host,
			Headers: Headers,
			Body:    file,
		}
		_, err = rest.API(request)
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
		http.ListenAndServe(port, nil)
	}
}
