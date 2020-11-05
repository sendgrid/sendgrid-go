package main

import (
	"fmt"
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
)

// Retrieveemailstatisticsbybrowser : Retrieve email statistics by browser.
// GET /browsers/stats
func Retrieveemailstatisticsbybrowser() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/browsers/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["browsers"] = "test_string"
	queryParams["limit"] = "test_string"
	queryParams["offset"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
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

func main() {
	// add your function calls here
}
