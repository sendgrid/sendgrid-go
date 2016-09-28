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
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				fmt.Printf("%s\n", "End of File")
				w.WriteHeader(http.StatusOK)
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			slurp, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}
			header := p.Header.Get("Content-Disposition")
			header = header[17 : len(header)-1]
			fmt.Printf("%q: %q\n", header, slurp)
		}
	}
}

func main() {
	conf := LoadConfig("./conf.json")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc(conf.Endpoint, inboundHandler)
	http.ListenAndServe(conf.Port, nil)
}
