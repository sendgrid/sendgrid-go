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
# Retrieve all categories #
# GET /categories #

request := sendgrid.GetRequest(apiKey, "/categories", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["category"] = "test_string"
queryParams["limit"] = "1"
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

##################################################
# Retrieve Email Statistics for Categories #
# GET /categories/stats #

request := sendgrid.GetRequest(apiKey, "/categories/stats", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["end_date"] = "2016-04-01"
queryParams["aggregated_by"] = "day"
queryParams["limit"] = "1"
queryParams["offset"] = "1"
queryParams["start_date"] = "2016-01-01"
queryParams["categories"] = "test_string"
request.QueryParams = queryParams
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Retrieve sums of email stats for each category [Needs: Stats object defined, has category ID?] #
# GET /categories/stats/sums #

request := sendgrid.GetRequest(apiKey, "/categories/stats/sums", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["end_date"] = "2016-04-01"
queryParams["aggregated_by"] = "day"
queryParams["limit"] = "1"
queryParams["sort_by_metric"] = "test_string"
queryParams["offset"] = "1"
queryParams["start_date"] = "2016-01-01"
queryParams["sort_by_direction"] = "asc"
request.QueryParams = queryParams
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

