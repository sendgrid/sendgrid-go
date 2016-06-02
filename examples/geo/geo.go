package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

apiKey := "SENDGRID_APIKEY"
host = "https://api.sendgrid.com"

##################################################
# Retrieve email statistics by country and state/province. #
# GET /geo/stats #

request := sendgrid.GetRequest(apiKey, "/geo/stats", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["end_date"] = "2016-04-01"
queryParams["country"] = "US"
queryParams["aggregated_by"] = "day"
queryParams["limit"] = "1"
queryParams["offset"] = "1"
queryParams["start_date"] = "2016-01-01"
request.QueryParams = queryParams
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

