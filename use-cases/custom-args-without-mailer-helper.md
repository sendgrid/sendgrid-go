## CustomArgs - Without Mail Helper Class

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
  request.Body = []byte(` {
    "personalizations": [
      {
        "to": [
            {
                "email": "test@example.com"
            }
        ],
        "subject": "CustomArgs can be fun",
        "substitutions": {
          "-name-": "Example User",
          "-city-": "Denver"
        }, 
        "custom_args": {
          "user_id": "343", 
          "batch_id": "3"
        }
      }
    ],
    "from": {
        "email": "test@example.com"
    },
    "content": [
        {
            "type": "text/html",
            "value": "<html>\n<head>\n\t<title></title>\n</head>\n<body>\nHello -name-,\n<br /><br/>\nI'm glad you are trying out the CustomArgs feature!\n<br /><br/>\nI hope you are having a great day in -city- :)\n<br /><br/>\n</body>\n</html>"
        }
    ], 
    "custom_args": {
      "campaign": "welcome",
      "weekday": "morning"
    } 
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