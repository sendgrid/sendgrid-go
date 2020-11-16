package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// GetMessages : Filter all messages
// GET /messages
func GetMessages() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/messages", host)
	request.Method = "GET"

	filterKey := "to_email"
	filterOperator := url.QueryEscape("=")
	filterValue := "testing@sendgrid.net"
	filterValue = url.QueryEscape(fmt.Sprintf("\"%s\"", filterValue))

	queryParams := make(map[string]string)
	queryParams["query"] = fmt.Sprintf("%s%s%s", filterKey, filterOperator, filterValue)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams

	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetMessagesByID : Filter messages by message ID
// GET /messages/{msg_id}
func GetMessagesByID() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/messages/{msg_id}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// RequestCSV : Request a CSV
// POST /messages/download
func RequestCSV() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/messages/download", host)
	request.Method = "POST"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DownloadCSV : Download CSV
// GET /messages/download/{download_uuid}
func DownloadCSV() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/messages/download/{download_uuid}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// add your function calls here
}
