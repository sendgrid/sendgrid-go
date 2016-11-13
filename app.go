package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go/helpers/inbound"
)

type configuration struct {
	Endpoint string `json:"endpoint"`
	Port     string `json:"port"`
}

func main() {
	if len(os.Args) < 2 {
		conf := loadConfig("./conf.json")
		http.HandleFunc("/", inbound.IndexHandler)
		http.HandleFunc(conf.Endpoint, inbound.InboundHandler)
		port := os.Getenv("PORT")
		if port == "" {
			port = conf.Port
		}
		log.Println("Listening And Serving On Port", port)
		http.ListenAndServe(":"+port, nil)
	} else {
		inbound.TestSender()
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
