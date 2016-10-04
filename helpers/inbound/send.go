package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sendgrid/rest"
)

func main() {
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
	response, err := rest.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
