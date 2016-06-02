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
# Create Subuser #
# POST /subusers #

request := sendgrid.GetRequest(apiKey, "/subusers", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "email": "John@example.com", 
  "ips": [
    "1.1.1.1", 
    "2.2.2.2"
  ], 
  "password": "johns_password", 
  "username": "John@example.com"
}`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# List all Subusers #
# GET /subusers #

request := sendgrid.GetRequest(apiKey, "/subusers", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["username"] = "test_string"
queryParams["limit"] = "0"
queryParams["offset"] = "0"
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
# Retrieve Subuser Reputations #
# GET /subusers/reputations #

request := sendgrid.GetRequest(apiKey, "/subusers/reputations", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["usernames"] = "test_string"
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
# Retrieve email statistics for your subusers. #
# GET /subusers/stats #

request := sendgrid.GetRequest(apiKey, "/subusers/stats", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["end_date"] = "2016-04-01"
queryParams["aggregated_by"] = "day"
queryParams["limit"] = "1"
queryParams["offset"] = "1"
queryParams["start_date"] = "2016-01-01"
queryParams["subusers"] = "test_string"
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
# Retrieve monthly stats for all subusers #
# GET /subusers/stats/monthly #

request := sendgrid.GetRequest(apiKey, "/subusers/stats/monthly", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["subuser"] = "test_string"
queryParams["limit"] = "1"
queryParams["sort_by_metric"] = "test_string"
queryParams["offset"] = "1"
queryParams["date"] = "test_string"
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

##################################################
#  Retrieve the totals for each email statistic metric for all subusers. #
# GET /subusers/stats/sums #

request := sendgrid.GetRequest(apiKey, "/subusers/stats/sums", host, "v3")
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

##################################################
# Enable/disable a subuser #
# PATCH /subusers/{subuser_name} #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}", host, "v3")
request.Method = "PATCH"
request.RequestBody = []byte(` {
  "disabled": false
}`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Delete a subuser #
# DELETE /subusers/{subuser_name} #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}", host, "v3")
request.Method = "DELETE"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Update IPs assigned to a subuser #
# PUT /subusers/{subuser_name}/ips #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/ips", host, "v3")
request.Method = "PUT"
request.RequestBody = []byte(` [
  "127.0.0.1"
]`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Update Monitor Settings for a subuser #
# PUT /subusers/{subuser_name}/monitor #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/monitor", host, "v3")
request.Method = "PUT"
request.RequestBody = []byte(` {
  "email": "example@example.com", 
  "frequency": 500
}`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Create monitor settings #
# POST /subusers/{subuser_name}/monitor #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/monitor", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "email": "example@example.com", 
  "frequency": 50000
}`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Retrieve monitor settings for a subuser #
# GET /subusers/{subuser_name}/monitor #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/monitor", host, "v3")
request.Method = "GET"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Delete monitor settings #
# DELETE /subusers/{subuser_name}/monitor #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/monitor", host, "v3")
request.Method = "DELETE"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Retrieve the monthly email statistics for a single subuser #
# GET /subusers/{subuser_name}/stats/monthly #

request := sendgrid.GetRequest(apiKey, "/subusers/{subuser_name}/stats/monthly", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["date"] = "test_string"
queryParams["sort_by_direction"] = "asc"
queryParams["limit"] = "0"
queryParams["sort_by_metric"] = "test_string"
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

