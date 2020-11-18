
### Personalization (with helper) - Sending Two Different Emails to Two Different Groups of Recipients

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

  from := mail.NewEmail("test", "test@example.com")
  content := mail.NewContent("text/html", "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>")

  m.SetFrom(from)
  m.AddContent(content)
  
  // create new *Personalization(s)
  personalization1 := mail.NewPersonalization()
  personalization2 := mail.NewPersonalization()
  
  // populate `personalization1` with data
  p1_to := mail.NewEmail("Example User 1", "test1@example.com")
  p1_cc1 := mail.NewEmail("Example User 2", "test2@example.com")
  p1_cc2 := mail.NewEmail("Example User 3", "test3@example.com")
  p1_cc3 := mail.NewEmail("Example User 3", "test4@example.com")
  
  personalization1.AddTos(p1_to)
  personalization1.AddCCs(p1_cc1, p1_cc2, p1_cc3)
  personalization1.SetSubstitution("%fname%", "recipient")
  personalization1.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization1.Subject = "Having fun learning about personalizations?"

  // populate `personalization2` with data
  p2_to := mail.NewEmail("Example User 1", "test1@example.com")
  p2_cc1 := mail.NewEmail("Example User 2", "test2@example.com")
  p2_cc2 := mail.NewEmail("Example User 3", "test3@example.com")
  p2_cc3 := mail.NewEmail("Example User 3", "test4@example.com")
  
  personalization2.AddTos(p2_to)
  personalization2.AddCCs(p2_cc1, p2_cc2, p2_cc3)
  personalization2.SetSubstitution("%fname%", "recipient2")
  personalization2.SetSubstitution("%CustomerID%", "55")
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