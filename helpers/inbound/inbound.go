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
	"strings"
)

type inbound_parse struct {
	Test1 string
	Test2 string
}

type Configuration struct {
	Endpoint string `json:"endpoint"`
	Port     string `json:"port"`
}

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}
	var conf Configuration
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return conf
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello World")
}

func inboundHandler(w http.ResponseWriter, r *http.Request) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		fmt.Println("================MESSAGE RECEIVED===============")
		mr := multipart.NewReader(r.Body, params["boundary"])
		parsedEmail := make(map[string]string)
		emailHeader := make(map[string]string)
		for {
			p, err := mr.NextPart()
			// We have found an attachment
			if err == nil && p.FileName() != "" {
				fmt.Println("FileName: ", p.FileName())
				contentType := p.Header.Get("Content-Type")
				fmt.Println("Content-Type: ", contentType)
				contents, err := ioutil.ReadAll(p)
				if err != nil {
					log.Fatal(err)
				}
				// Binary file contents
				fmt.Println("Contents: ", contents)
				// Only works with text files, this is just for testing
				fmt.Println("Contents Decoded: ", string(contents))
			}
			if err == io.EOF {
				// We have finished parsing
				for key, value := range parsedEmail {
					fmt.Println("Key:", key, " Value:", value)
				}
				for key, value := range emailHeader {
					fmt.Println("eKey:", key, " eValue:", value)
				}
				w.WriteHeader(http.StatusOK)
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
				header = header[11 :]
				f := strings.Split(header, "=")
				parsedEmail[f[1][1:len(f[1])-11]] = f[2][1:len(f[2])-1]
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
			delete(parsedEmail, "headers")
		}
	}
}

func main() {
	conf := LoadConfig("./conf.json")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc(conf.Endpoint, inboundHandler)
	http.ListenAndServe(conf.Port, nil)
}
