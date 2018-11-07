## Attachments - Without Mail Helper Class
 
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
                     "email": "test1@example.com"
                 }
             ],
             "subject": "Attachments - Demystified!"
             }
         }
     ],
     "from": {
         "email": "test@example.com"
     },
     "content": [
         {
             "type": "text/html",
             "value": "<p>Sending different attachments.</p>"
         }
     ], 
     "attachments": [
        {
          "content": "SGVsbG8gV29ybGQh", 
          "content_id": "testing_1", 
          "disposition": "attachment", 
          "filename": "testing.txt", 
          "type": "txt"
        },
        {
          "content": "BASE64 encoded content block here", 
          "content_id": "testing_2", 
          "disposition": "attachment", 
          "filename": "testing.jpg", 
          "type": "jpg"
        },
        {
          "content": "BASE64 encoded content block here", 
          "content_id": "testing_3", 
          "disposition": "attachment", 
          "filename": "testing.pdf", 
          "type": "pdf"
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