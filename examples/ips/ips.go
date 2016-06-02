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
# Retrieve all IP addresses #
# GET /ips #

request := sendgrid.GetRequest(apiKey, "/ips", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
queryParams["subuser"] = "test_string"
queryParams["ip"] = "test_string"
queryParams["limit"] = "1"
queryParams["exclude_whitelabels"] = "true"
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
# Retrieve all assigned IPs #
# GET /ips/assigned #

request := sendgrid.GetRequest(apiKey, "/ips/assigned", host, "v3")
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
# Create an IP pool. #
# POST /ips/pools #

request := sendgrid.GetRequest(apiKey, "/ips/pools", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "name": "marketing"
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
# Retrieve all IP pools. #
# GET /ips/pools #

request := sendgrid.GetRequest(apiKey, "/ips/pools", host, "v3")
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
# Update an IP pools name. #
# PUT /ips/pools/{pool_name} #

request := sendgrid.GetRequest(apiKey, "/ips/pools/{pool_name}", host, "v3")
request.Method = "PUT"
request.RequestBody = []byte(` {
  "name": "new_pool_name"
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
# Retrieve all IPs in a specified pool. #
# GET /ips/pools/{pool_name} #

request := sendgrid.GetRequest(apiKey, "/ips/pools/{pool_name}", host, "v3")
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
# Delete an IP pool. #
# DELETE /ips/pools/{pool_name} #

request := sendgrid.GetRequest(apiKey, "/ips/pools/{pool_name}", host, "v3")
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
# Add an IP address to a pool #
# POST /ips/pools/{pool_name}/ips #

request := sendgrid.GetRequest(apiKey, "/ips/pools/{pool_name}/ips", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "ip": "0.0.0.0"
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
# Remove an IP address from a pool. #
# DELETE /ips/pools/{pool_name}/ips/{ip} #

request := sendgrid.GetRequest(apiKey, "/ips/pools/{pool_name}/ips/{ip}", host, "v3")
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
# Add an IP to warmup #
# POST /ips/warmup #

request := sendgrid.GetRequest(apiKey, "/ips/warmup", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "ip": "0.0.0.0"
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
# Retrieve all IPs currently in warmup #
# GET /ips/warmup #

request := sendgrid.GetRequest(apiKey, "/ips/warmup", host, "v3")
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
# Retrieve warmup status for a specific IP address #
# GET /ips/warmup/{ip_address} #

request := sendgrid.GetRequest(apiKey, "/ips/warmup/{ip_address}", host, "v3")
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
# Remove an IP from warmup #
# DELETE /ips/warmup/{ip_address} #

request := sendgrid.GetRequest(apiKey, "/ips/warmup/{ip_address}", host, "v3")
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
# Retrieve all IP pools an IP address belongs to #
# GET /ips/{ip_address} #

request := sendgrid.GetRequest(apiKey, "/ips/{ip_address}", host, "v3")
request.Method = "GET"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

