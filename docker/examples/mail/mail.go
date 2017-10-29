package main

import (
  "os"
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  "log"
)

func SendSampleHelloMail() {
  apiKey := os.Getenv("SENDGRID_API_KEY")
  host := "http://localhost:4010"
  request := sendgrid.GetRequest(apiKey, "/v3/mail/send", host)
  request.Method = "POST"
  request.Body = []byte(` {
  "personalizations": [
    {
      "to": [
        {
          "email": "test@example.com"
        }
      ],
      "subject": "Sending with SendGrid is Fun"
    }
  ],
  "from": {
    "email": "test@example.com"
  },
  "subject": "Hello SendGrid",
  "content": [
    {
      "type": "text/plain",
      "value": "and easy to do anywhere, even with Go"
    }
  ]
}`)
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
