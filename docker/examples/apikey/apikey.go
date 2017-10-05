package main

import (
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  "log"
)

func GetAPIKey() {
  apiKey := "JUST_A_TEST_KEY"
  host := "http://localhost:4010"
  request := sendgrid.GetRequest(apiKey, "/v3/api_keys", host)
  request.Method = "GET"

  response, err := sendgrid.API(request)
  if err != nil {
    log.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}

func main() {
  // add your function calls here
}
