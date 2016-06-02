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
# Create a Campaign #
# POST /campaigns #

request := sendgrid.GetRequest(apiKey, "/campaigns", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "categories": [
    "spring line"
  ], 
  "custom_unsubscribe_url": "", 
  "html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>", 
  "ip_pool": "marketing", 
  "list_ids": [
    110, 
    124
  ], 
  "plain_content": "Check out our spring line!", 
  "segment_ids": [
    110
  ], 
  "sender_id": 124451, 
  "subject": "New Products for Spring!", 
  "suppression_group_id": 42, 
  "title": "March Newsletter"
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
# Retrieve all Campaigns #
# GET /campaigns #

request := sendgrid.GetRequest(apiKey, "/campaigns", host, "v3")
request.Method = "GET"
queryParams := make(map[string]string)
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
# Update a Campaign #
# PATCH /campaigns/{campaign_id} #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}", host, "v3")
request.Method = "PATCH"
request.RequestBody = []byte(` {
  "categories": [
    "summer line"
  ], 
  "html_content": "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>", 
  "plain_content": "Check out our summer line!", 
  "subject": "New Products for Summer!", 
  "title": "May Newsletter"
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
# Retrieve a single campaign #
# GET /campaigns/{campaign_id} #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}", host, "v3")
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
# Delete a Campaign #
# DELETE /campaigns/{campaign_id} #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}", host, "v3")
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
# Update a Scheduled Campaign #
# PATCH /campaigns/{campaign_id}/schedules #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules", host, "v3")
request.Method = "PATCH"
request.RequestBody = []byte(` {
  "send_at": 1489451436
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
# Schedule a Campaign #
# POST /campaigns/{campaign_id}/schedules #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "send_at": 1489771528
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
# View Scheduled Time of a Campaign #
# GET /campaigns/{campaign_id}/schedules #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules", host, "v3")
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
# Unschedule a Scheduled Campaign #
# DELETE /campaigns/{campaign_id}/schedules #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules", host, "v3")
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
# Send a Campaign #
# POST /campaigns/{campaign_id}/schedules/now #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules/now", host, "v3")
request.Method = "POST"
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

##################################################
# Send a Test Campaign #
# POST /campaigns/{campaign_id}/schedules/test #

request := sendgrid.GetRequest(apiKey, "/campaigns/{campaign_id}/schedules/test", host, "v3")
request.Method = "POST"
request.RequestBody = []byte(` {
  "to": "your.email@example.com"
}`)
response, err := sendgrid.API(request)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(response.StatusCode)
  fmt.Println(response.ResponseBody)
  fmt.Println(response.ResponseHeaders)
}

