package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve all recent access attempts
// GET /access_settings/activity

func Retrieveallrecentaccessattempts() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/activity", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["limit"] = "1"
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
// Add one or more IPs to the whitelist
// POST /access_settings/whitelist

func AddoneormoreIPstothewhitelist() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/whitelist", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "ips": [
    {
      "ip": "192.168.1.1"
    }, 
    {
      "ip": "192.*.*.*"
    }, 
    {
      "ip": "192.168.1.3/32"
    }
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
}

///////////////////////////////////////////////////
// Retrieve a list of currently whitelisted IPs
// GET /access_settings/whitelist

func RetrievealistofcurrentlywhitelistedIPs() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/whitelist", host, "v3")
  request.Method = "GET"
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
// Remove one or more IPs from the whitelist
// DELETE /access_settings/whitelist

func RemoveoneormoreIPsfromthewhitelist() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/whitelist", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` {
  "ids": [
    1, 
    2, 
    3
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
}

///////////////////////////////////////////////////
// Retrieve a specific whitelisted IP
// GET /access_settings/whitelist/{rule_id}

func RetrieveaspecificwhitelistedIP() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/whitelist/{rule_id}", host, "v3")
  request.Method = "GET"
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
// Remove a specific IP from the whitelist
// DELETE /access_settings/whitelist/{rule_id}

func RemoveaspecificIPfromthewhitelist() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/access_settings/whitelist/{rule_id}", host, "v3")
  request.Method = "DELETE"
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