### Personalization (without helper) - Sending Two Different Emails to Two Different Groups of Recipients from two different From email address

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/sendgrid/sendgrid-go"
)

func main() {
  request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
  request.Method = "POST"
  request.Body = []byte(`{
  "personalizations": [{
      "to": [{
          "email": "recipient1@example.com"
      }],
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }, {
      "to": [{
          "email": "recipient2@example.com"
      }],
      "from": {
          "email": "sender2@example.com"
      },
      "subject": "YOUR OTHER SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "defaultSender@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p>Personalizations are awesome!</p>"
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
```