package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Returns a list of all partner settings.
// GET /partner_settings

func Returnsalistofallpartnersettings() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/partner_settings", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
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
}

///////////////////////////////////////////////////
// Updates New Relic partner settings.
// PATCH /partner_settings/new_relic

func UpdatesNewRelicpartnersettings() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/partner_settings/new_relic", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enable_subuser_statistics": true, 
  "enabled": true, 
  "license_key": ""
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
// Returns all New Relic partner settings.
// GET /partner_settings/new_relic

func ReturnsallNewRelicpartnersettings() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/partner_settings/new_relic", host, "v3")
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

func main() {
    // add your function calls here
}