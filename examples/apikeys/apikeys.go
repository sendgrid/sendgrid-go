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
# Create API keys #
# POST /api_keys #

request := sendgrid.GetRequest(apiKey, "/api_keys", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "name": "My API Key", 
  "scopes": [
    "mail.send", 
    "alerts.create", 
    "alerts.read"
  ]
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
# Retrieve all API Keys belonging to the authenticated user #
# GET /api_keys #

request := sendgrid.GetRequest(apiKey, "/api_keys", host, "v3")
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
# Update the name & scopes of an API Key #
# PUT /api_keys/{api_key_id} #

request := sendgrid.GetRequest(apiKey, "/api_keys/{api_key_id}", host, "v3")
request.Method = "PUT"
request.RequestBody = []byte(` {
  "name": "A New Hope", 
  "scopes": [
    "user.profile.read", 
    "user.profile.update"
  ]
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
# Update API keys #
# PATCH /api_keys/{api_key_id} #

request := sendgrid.GetRequest(apiKey, "/api_keys/{api_key_id}", host, "v3")
request.Method = "PATCH"
request.RequestBody = []byte(` {
  "name": "A New Hope"
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
# Retrieve an existing API Key #
# GET /api_keys/{api_key_id} #

request := sendgrid.GetRequest(apiKey, "/api_keys/{api_key_id}", host, "v3")
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
# Delete API keys #
# DELETE /api_keys/{api_key_id} #

request := sendgrid.GetRequest(apiKey, "/api_keys/{api_key_id}", host, "v3")
request.Method = "DELETE"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

