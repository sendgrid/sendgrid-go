package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve all blocks
// GET /suppression/blocks

func Retrieveallblocks() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/blocks", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["start_time"] = "1"
queryParams["limit"] = "1"
queryParams["end_time"] = "1"
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
// Delete blocks
// DELETE /suppression/blocks

func Deleteblocks() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/blocks", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` {
  "delete_all": false, 
  "emails": [
    "example1@example.com", 
    "example2@example.com"
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
// Retrieve a specific block
// GET /suppression/blocks/{email}

func Retrieveaspecificblock() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/blocks/{email}", host, "v3")
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
// Delete a specific block
// DELETE /suppression/blocks/{email}

func Deleteaspecificblock() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/blocks/{email}", host, "v3")
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

///////////////////////////////////////////////////
// Retrieve all bounces
// GET /suppression/bounces

func Retrieveallbounces() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/bounces", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["start_time"] = "0"
queryParams["end_time"] = "0"
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
// Delete bounces
// DELETE /suppression/bounces

func Deletebounces() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/bounces", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` {
  "delete_all": true, 
  "emails": [
    "example@example.com", 
    "example2@example.com"
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
// Retrieve a Bounce
// GET /suppression/bounces/{email}

func RetrieveaBounce() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/bounces/{email}", host, "v3")
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
// Delete a bounce
// DELETE /suppression/bounces/{email}

func Deleteabounce() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/bounces/{email}", host, "v3")
  request.Method = "DELETE"
  queryParams := make(map[string]string)
  queryParams["email_address"] = "example@example.com"
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
// Retrieve all invalid emails
// GET /suppression/invalid_emails

func Retrieveallinvalidemails() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/invalid_emails", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["start_time"] = "1"
queryParams["limit"] = "1"
queryParams["end_time"] = "1"
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
// Delete invalid emails
// DELETE /suppression/invalid_emails

func Deleteinvalidemails() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/invalid_emails", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` {
  "delete_all": false, 
  "emails": [
    "example1@example.com", 
    "example2@example.com"
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
// Retrieve a specific invalid email
// GET /suppression/invalid_emails/{email}

func Retrieveaspecificinvalidemail() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/invalid_emails/{email}", host, "v3")
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
// Delete a specific invalid email
// DELETE /suppression/invalid_emails/{email}

func Deleteaspecificinvalidemail() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/invalid_emails/{email}", host, "v3")
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

///////////////////////////////////////////////////
// Retrieve a specific spam report
// GET /suppression/spam_report/{email}

func Retrieveaspecificspamreport() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/spam_report/{email}", host, "v3")
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
// Delete a specific spam report
// DELETE /suppression/spam_report/{email}

func Deleteaspecificspamreport() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/spam_report/{email}", host, "v3")
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

///////////////////////////////////////////////////
// Retrieve all spam reports
// GET /suppression/spam_reports

func Retrieveallspamreports() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/spam_reports", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["start_time"] = "1"
queryParams["limit"] = "1"
queryParams["end_time"] = "1"
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
// Delete spam reports
// DELETE /suppression/spam_reports

func Deletespamreports() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/spam_reports", host, "v3")
  request.Method = "DELETE"
  request.RequestBody = []byte(` {
  "delete_all": false, 
  "emails": [
    "example1@example.com", 
    "example2@example.com"
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
// Retrieve all global suppressions
// GET /suppression/unsubscribes

func Retrieveallglobalsuppressions() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/suppression/unsubscribes", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["start_time"] = "1"
queryParams["limit"] = "1"
queryParams["end_time"] = "1"
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

