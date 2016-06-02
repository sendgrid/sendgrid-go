package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve email statistics by client type.
// GET /clients/stats

func Retrieveemailstatisticsbyclienttype() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/clients/stats", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["aggregated_by"] = "day"
queryParams["start_date"] = "2016-01-01"
queryParams["end_date"] = "2016-04-01"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve stats by a specific client type.
// GET /clients/{client_type}/stats

func Retrievestatsbyaspecificclienttype() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/clients/{client_type}/stats", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["aggregated_by"] = "day"
queryParams["start_date"] = "2016-01-01"
queryParams["end_date"] = "2016-04-01"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

func main() {
    // add your function calls here
}
