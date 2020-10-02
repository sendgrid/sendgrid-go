
### Personalization (with helper) - Sending Two Different Emails to Two Different Groups of Recipients from two different From emails

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
  // create new *SGMailV3
  m := mail.NewV3Mail()

  from := mail.NewEmail("Example Sender 1", "defaultSender@example.com")
  content := mail.NewContent("text/html", "<p>Personalizations are awesome!</p>")

  m.SetFrom(from)
  m.AddContent(content)
  
  // create new *Personalization(s)
  personalization1 := mail.NewPersonalization()
  personalization2 := mail.NewPersonalization()
  
  // populate `personalization1` with data 
  //this email will be sent from Example Sender 1
  p1_to := mail.NewEmail("Example User 1", "test1@example.com")
  
  personalization1.AddTos(p1_to)
  personalization1.Subject = "Having fun learning about personalizations?"

  // populate `personalization2` with data 
  //this email will be sent from Example Sender 2
  p2_from :=mail.NewEmail("Example Sender 2", "sender2@example.com")
  p2_to := mail.NewEmail("Example User 1", "test1@example.com")
  
  personalization2.AddFrom(p2_from)
  personalization2.AddTos(p2_to)
  personalization2.Subject = "Personalizations are fun!"
  
  // add `personalization1` and `personalization2` to `m`
  m.AddPersonalizations(personalization1, personalization2)

  request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
  request.Method = "POST"
  request.Body = mail.GetRequestBody(m)
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