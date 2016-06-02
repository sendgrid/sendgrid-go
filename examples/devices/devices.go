package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve email statistics by device type.
// GET /devices/stats

func Retrieveemailstatisticsbydevicetype() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/devices/stats", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["aggregated_by"] = "day"
queryParams["limit"] = "1"
queryParams["start_date"] = "2016-01-01"
queryParams["end_date"] = "2016-04-01"
queryParams["offset"] = "1"
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