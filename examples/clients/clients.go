package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

// GetEmailStatsByClientType Retrieve email statistics by client type.
// GET /clients/stats
func GetEmailStatsByClientType() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/clients/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
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

// GetStatsByClientType Retrieve stats by a specific client type.
// GET /clients/{client_type}/stats
func GetStatsByClientType() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/clients/{client_type}/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
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
