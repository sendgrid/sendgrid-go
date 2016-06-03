package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Create a Group
// POST /asm/groups

func CreateaGroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups", host, "v3")
  request.Method = "POST"
  request.Body = []byte(` {
  "description": "A group description", 
  "is_default": false, 
  "name": "A group name"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Retrieve all suppression groups associated with the user.
// GET /asm/groups

func Retrieveallsuppressiongroupsassociatedwiththeuser() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Update a suppression group.
// PATCH /asm/groups/{group_id}

func Updateasuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}", host, "v3")
  request.Method = "PATCH"
  request.Body = []byte(` {
  "description": "Suggestions for items our users might like.", 
  "id": 103, 
  "name": "Item Suggestions"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Get information on a single suppression group.
// GET /asm/groups/{group_id}

func Getinformationonasinglesuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Delete a suppression group.
// DELETE /asm/groups/{group_id}

func Deleteasuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Add suppressions to a suppression group
// POST /asm/groups/{group_id}/suppressions

func Addsuppressionstoasuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}/suppressions", host, "v3")
  request.Method = "POST"
  request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com", 
    "test2@example.com"
  ]
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Retrieve all suppressions for a suppression group
// GET /asm/groups/{group_id}/suppressions

func Retrieveallsuppressionsforasuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}/suppressions", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Delete a suppression from a suppression group
// DELETE /asm/groups/{group_id}/suppressions/{email}

func Deleteasuppressionfromasuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/groups/{group_id}/suppressions/{email}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Add recipient addresses to the global suppression group.
// POST /asm/suppressions/global

func Addrecipientaddressestotheglobalsuppressiongroup() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/suppressions/global", host, "v3")
  request.Method = "POST"
  request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com", 
    "test2@example.com"
  ]
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Retrieve a Global Suppression
// GET /asm/suppressions/global/{email}

func RetrieveaGlobalSuppression() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/suppressions/global/{email}", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

///////////////////////////////////////////////////
// Delete a Global Suppression
// DELETE /asm/suppressions/global/{email}

func DeleteaGlobalSuppression() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/asm/suppressions/global/{email}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

func main() {
    // add your function calls here
}
