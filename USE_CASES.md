This documentation provides examples for specific use cases. Please [open an issue](https://github.com/sendgrid/sendgrid-go/issues) or make a pull request for any use cases you would like us to document here. Thank you!

# Table of Contents

* [Transactional Templates](#transactional_templates)
* [CustomArgs](#customargs)
* [Personalizations](#personalizations)
* [Substitutions](#substitutions)
* [Sections](#sections)
* [Attachments](#attachments)
* [Deployments](#deployments)

<a name="transactional_templates"></a>
# Transactional Templates

For this example, we assume you have created a [transactional template](https://sendgrid.com/docs/User_Guide/Transactional_Templates/index.html). Following is the template content we used for testing.

Template ID (replace with your own):

```text
13b8f94f-bcae-4ec6-b752-70d6cb59f932
```

Email Subject:

```text
<%subject%>
```

Template Body:

```html
<html>
<head>
  <title></title>
</head>
<body>
Hello -name-,
<br /><br/>
I'm glad you are trying out the template feature!
<br /><br/>
<%body%>
<br /><br/>
I hope you are having a great day in -city- :)
<br /><br/>
</body>
</html>
```

## With Mail Helper Class

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
  from := mail.NewEmail("Example User", "test@example.com")
  subject := "I'm replacing the subject tag"
  to := mail.NewEmail("Example User", "test@example.com")
  content := mail.NewContent("text/html", "I'm replacing the <strong>body tag</strong>")
  m := mail.NewV3MailInit(from, subject, to, content)
  m.Personalizations[0].SetSubstitution("-name-", "Example User")
  m.Personalizations[0].SetSubstitution("-city-", "Denver")
  m.SetTemplateID("13b8f94f-bcae-4ec6-b752-70d6cb59f932")

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

## Without Mail Helper Class

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
            "subject": "I'm replacing the subject tag",
            "substitutions": {
              "-name-": "Example User",
              "-city-": "Denver"
            },
        }
    ],
    "from": {
        "email": "test@example.com"
    },
    "content": [
        {
            "type": "text/html",
            "value": "I'm replacing the <strong>body tag</strong>"
        }
    ],
    "template_id": "13b8f94f-bcae-4ec6-b752-70d6cb59f932"
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

<a name="customargs"></a>
# CustomArgs

## With Mail Helper Class

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
  from := mail.NewEmail("Example User", "test@example.com")
  subject := "CustomArgs can be fun"
  to := mail.NewEmail("Example User", "test@example.com")
  content := mail.NewContent("text/html", "<html>\n<head>\n\t<title></title>\n</head>\n<body>\nHello -name-,\n<br /><br/>\nI'm glad you are trying out the CustomArgs feature!\n<br /><br/>\nI hope you are having a great day in -city- :)\n<br /><br/>\n</body>\n</html>")
  m := mail.NewV3MailInit(from, subject, to, content)
  m.Personalizations[0].SetSubstitution("-name-", "Example User")
  m.Personalizations[0].SetSubstitution("-city-", "Denver")
  m.Personalizations[0].SetCustomArg("user_id", "343")
  m.Personalizations[0].SetCustomArg("batch_id", "3")

  m.SetCustomArg("campaign", "welcome")
  m.SetCustomArg("weekday", "morning")

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

## Without Mail Helper Class

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

<a name="personalizations"></a>
# Personalizations

## With Mail Helper Class

### Sending a Single Email to a Single Recipient

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
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  
  // populate `personalization` with data
  to := mail.NewEmail("Example User", "test1@example.com")
  
  personalization.AddTos(to)
  personalization.SetSubstitution("%fname%", "recipient")
  personalization.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization.Subject = "Having fun learning about personalizations?"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

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

### Sending a Single Email to a Single Recipient with a CC

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
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  
  // populate `personalization` with data
  to := mail.NewEmail("Example User", "test1@example.com")
  cc1 := mail.NewEmail("Example CC", "test2@example.com")

  personalization.AddTos(to)
  personalization.AddCCs(cc1)
  personalization.SetSubstitution("%fname%", "recipient")
  personalization.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization.Subject = "Having fun learning about personalizations?"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

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

### Sending a Single Email to a Single Recipient with a CC and a BCC

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
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  
  // populate `personalization` with data
  to := mail.NewEmail("Example User", "test1@example.com")
  cc1 := mail.NewEmail("Example CC", "test2@example.com")
  bcc1 := mail.NewEmail("Example BCC", "test3@example.com")
  
  personalization.AddTos(to)
  personalization.AddCCs(cc1)
  personalization.AddBCCs(bcc1)
  personalization.SetSubstitution("%fname%", "recipient")
  personalization.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization.Subject = "Having fun learning about personalizations?"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

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

### Sending a Single Email to Multiple Recipients

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
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  
  // populate `personalization` with data
  to1 := mail.NewEmail("Example User 1", "test1@example.com")
  to2 := mail.NewEmail("Example User 2", "test2@example.com")
  to3 := mail.NewEmail("Example User 3", "test3@example.com")
  
  personalization.AddTos(to1, to2, to3)
  personalization.SetSubstitution("%fname%", "recipient")
  personalization.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization.Subject = "Having fun learning about personalizations?"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

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

### Sending a Single Email to a Single Recipient with Multiple CCs/BCCs

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
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  
  // populate `personalization` with data
  to := mail.NewEmail("Example User 1", "test1@example.com")
  cc1 := mail.NewEmail("Example User 2", "test2@example.com")
  cc2 := mail.NewEmail("Example User 3", "test3@example.com")
  cc3 := mail.NewEmail("Example User 3", "test4@example.com")
  
  personalization.AddTos(to)
  personalization.AddCCs(cc1, cc2, cc3)
  personalization.SetSubstitution("%fname%", "recipient")
  personalization.SetSubstitution("%CustomerID%", "CUSTOMER ID GOES HERE")
  personalization.Subject = "Having fun learning about personalizations?"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

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

### Sending Two Different Emails to Two Different Groups of Recipients

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

## Without Mail Helper Class

### Sending A Single Email to a Single Recipient

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
          "email": "test1@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending a Single Email to a Single Recipient With a CC

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
      "cc": [{
          "email": "recipient2@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending a Single Email to a Single Recipient With a CC

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
      "cc": [{
          "email": "recipient2@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending a Single Email to a Single Recipient With a CC and a BCC

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
      "cc": [{
          "email": "recipient2@example.com"
      }],
      "bcc": [{
          "email": "recipient3@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      }
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending the same Email to Multiple Recipients

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
      }, {
          "email": "recipient2@example.com"
      }, {
          "email": "recipient3@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending a Single Email to a Single Recipient with Multiple CCs/BCCs

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
      "cc": [{
          "email": "recipient2@example.com"
      }, {
          "email": "recipient3@example.com"
      }, {
          "email": "recipient4@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

### Sending Two Different Emails to Two Different Groups of Recipients

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
      "cc": [{
          "email": "recipient2@example.com"
      }, {
          "email": "recipient3@example.com"
      }, {
          "email": "recipient4@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient",
          "%CustomerID%": "CUSTOMER ID GOES HERE"
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }, {
      "to": [{
          "email": "recipient5@example.com"
      }],
      "cc": [{
          "email": "recipient6@example.com"
      }, {
          "email": "recipient7@example.com"
      }, {
          "email": "recipient8@example.com"
      }],
      "substitutions": {
          "%fname%": "recipient2",
          "%CustomerID%": 55
      },
      "subject": "YOUR SUBJECT LINE GOES HERE"
  }],
  "from": {
    "email": "test@example.com"
  },
  "content": [
      {
          "type": "text/html",
          "value": "<p> %fname% : %CustomerID% - Personalizations are awesome!</p>"
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

<a name="substitutions"></a>
# Substitutions
 
## With Mail Helper Class
 
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
  from := mail.NewEmail("Example User", "test@example.com")
  subject := "Substitutions can be fun"
  to := mail.NewEmail("Example User", "test@example.com")
  content := mail.NewContent("text/html", "<html>\n<head>\n\t<title></title>\n</head>\n<body>\nHello -name-,\n<br /><br/>\nI'm glad you are trying out the Substitutions feature!\n<br /><br/>\nI hope you are having a great day in -city- :)\n<br /><br/>\n</body>\n</html>")
  m := mail.NewV3MailInit(from, subject, to, content)
  m.Personalizations[0].SetSubstitution("-name-", "Example User")
  m.Personalizations[0].SetSubstitution("-city-", "Denver")
  m.Personalizations[0].SetSubstitution("-user_id-", "343")
  m.Personalizations[0].SetCustomArg("user_id", "-user_id-")
  m.Personalizations[0].SetCustomArg("city", "-city-")
 
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
 
## Without Mail Helper Class
 
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
             "subject": "Substitutions can be fun",
             "substitutions": {
               "-name-": "Example User",
               "-city-": "Denver",
               "-user_id-": "343"
             }, 
             "custom_args": {
               "user_id": "-user_id-", 
               "city": "-city-"
             }
         }
     ],
     "from": {
         "email": "test@example.com"
     },
     "content": [
         {
             "type": "text/html",
             "value": "<html>\n<head>\n\t<title></title>\n</head>\n<body>\nHello -name-,\n<br /><br/>\nI'm glad you are trying out the Substitutions feature!\n<br /><br/>\nI hope you are having a great day in -city- :)\n<br /><br/>\n</body>\n</html>"
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
 
<a name="sections"></a>
# Sections
 
## With Mail Helper Class
 
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
  from := mail.NewEmail("Example User", "test@example.com")
  subject := "Sections can be fun"
  to := mail.NewEmail("Example User", "test@example.com")
  content := mail.NewContent("text/html", "<html>\n<head>\n\t<title></title>\n</head>\n<body>\n-wel-\n<br /><br/>\nI'm glad you are trying out the Sections feature!\n<br /><br/>\n-gday-\n<br /><br/>\n</body>\n</html>")
  m := mail.NewV3MailInit(from, subject, to, content)
  m.Personalizations[0].SetSubstitution("-name-", "Example User")
  m.Personalizations[0].SetSubstitution("-city-", "Denver")
  m.Personalizations[0].SetSubstitution("-wel-", "-welcome-")
  m.Personalizations[0].SetSubstitution("-gday-", "-great_day-")
 
  m.AddSection("-welcome-", "Hello -name-,")
  m.AddSection("-great_day-", "I hope you are having a great day in -city- :)")
 
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
 
## Without Mail Helper Class
 
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
             "subject": "Sections can be fun",
             "substitutions": {
               "-name-": "Example User",
               "-city-": "Denver",
               "-wel-": "-welcome-",
               "-gday-": "-great_day-"
             }
         }
     ],
     "from": {
         "email": "test@example.com"
     },
     "content": [
         {
             "type": "text/html",
             "value": "<html>\n<head>\n\t<title></title>\n</head>\n<body>\n-wel-\n<br /><br/>\nI'm glad you are trying out the Sections feature!\n<br /><br/>\n-gday-\n<br /><br/>\n</body>\n</html>"
         }
     ], 
   "sections": {
     "section": {
       "-welcome-": "Hello -name-,", 
       "-great_day-": "I hope you are having a great day in -city- :)"
     }
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
<a name="attachments"></a>
# Attachments
 
## With Mail Helper Class

```go
package main

import (
  "fmt"
  "log"
  "os"
  "encoding/base64"
  "io/ioutil"
  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
  // create new *SGMailV3
  m := mail.NewV3Mail()

  from := mail.NewEmail("test", "test@example.com")
  content := mail.NewContent("text/html", "<p>Sending different attachments.</p>")
  to := mail.NewEmail("Example User", "test1@example.com")

  m.SetFrom(from)
  m.AddContent(content)
  
  // create new *Personalization
  personalization := mail.NewPersonalization()
  personalization.AddTos(to)
  personalization.Subject = "Attachments - Demystified!"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)
  
  // read/attach .txt file
  a_txt := mail.NewAttachment()
  dat, err := ioutil.ReadFile("testing.txt")
  if err != nil {
    fmt.Println(err)
  }
  encoded := base64.StdEncoding.EncodeToString([]byte(dat))
  a_txt.SetContent(encoded)
  a_txt.SetType("text/plain")
  a_txt.SetFilename("testing.txt")
  a_txt.SetDisposition("attachment")
  a_txt.SetContentID("Test Document")
  
  // read/attach .pdf file
  a_pdf := mail.NewAttachment()
  dat, err = ioutil.ReadFile("testing.pdf")
  if err != nil {
    fmt.Println(err)
  }
  encoded = base64.StdEncoding.EncodeToString([]byte(dat))
  a_pdf.SetContent(encoded)
  a_pdf.SetType("application/pdf")
  a_pdf.SetFilename("testing.pdf")
  a_pdf.SetDisposition("attachment")
  a_pdf.SetContentID("Test Attachment")

  // read/attach .jpg file
  a_jpg := mail.NewAttachment()
  dat, err = ioutil.ReadFile("testing.jpg")
  if err != nil {
    fmt.Println(err)
  }
  encoded = base64.StdEncoding.EncodeToString([]byte(dat))
  a_jpg.SetContent(encoded)
  a_jpg.SetType("image/jpeg")
  a_jpg.SetFilename("testing.jpg")
  a_jpg.SetDisposition("attachment")
  a_jpg.SetContentID("Test Attachment")
  
  // add `a_txt`, `a_pdf` and `a_jpg` to `m`
  m.AddAttachment(a_txt)
  m.AddAttachment(a_pdf)
  m.AddAttachment(a_jpg)  
  
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
 
## Without Mail Helper Class
 
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
<a name="deployments"></a>
# Deployments
## Google App Engine

### Setup
#### Create a project
In order to deploy your application to Google App Engine, you must first create a Cloud Platform project and App Engine application using the Cloud Platform Console.

To do this, go to [App Engine](https://console.cloud.google.com/projectselector/appengine/create?lang=go&st=true&_ga=2.6200051.-1972623632.1507938170). 

From here, you may either select an existing project or create a new one. Once you have created a project, select a region. After Google App Engine initializes the backend services for your app, you are good to go.

#### Install the SDK
If you have already configured the SDK you may skip this step, otherwise you will need to install the SDK to deploy your app to Google App Engine.

Follow the steps to [Download the SDK](https://cloud.google.com/appengine/docs/standard/go/download).
Once you have the ```gcloud``` tool installed, make sure it is configured to use the project you created earlier.
If you are having trouble with the ```gcloud``` tool, take a look at the [```gcloud``` reference](https://cloud.google.com/sdk/gcloud/reference/).

### Example app
The following is a simple Go app which allows you to send a "Hello World" type email to someone with a simple POST request.
```go
package hellosendgrid

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/sendgrid/rest"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

type EmailSettings struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func init() {
	http.HandleFunc("/send", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// Make sure it is a POST request
	if r.Method == "POST" {
		// Grab request body
		settings := EmailSettings{}
		err := json.NewDecoder(r.Body).Decode(&settings)
		if err != nil {
			http.Error(w, "Failed to decode body", http.StatusInternalServerError)
			return
		}
		//Send the email
		response, err := SendEmail(settings, ctx)
		// Make sure there were no errors sending the email
		if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {
			http.Error(w, response.Body, response.StatusCode)
			return
		}
		w.WriteHeader(response.StatusCode)
		w.Write([]byte("Successfully sent email"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SendEmail(settings EmailSettings, ctx context.Context) (status *rest.Response, err error) {
	// NOTE!
	// Google app engine does not allow the use of the net/http client. Therefore, you must use urlfetch to make HTTP requests in app engine.
	// https://cloud.google.com/appengine/docs/standard/go/issue-requests
	sendgrid.DefaultClient.HTTPClient = urlfetch.Client(ctx)

	from := mail.NewEmail("Hello User", "test@example.com")
	subject := "Sending with SendGrid is Fun!"
	to := mail.NewEmail(settings.Name, settings.Email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	return response, err
}
```
#### App settings
In order to run you app on Google App Engine, you must have a ```app.yaml``` file in the application's root directory.
This file contains information about your application code.
For more information about the ```app.yaml``` file, please view the [```app.yaml``` reference](https://cloud.google.com/appengine/docs/standard/go/config/appref).
```yaml
runtime: go
api_version: go1

env_variables:
  SENDGRID_API_KEY: 'YOUR_API_KEY'

handlers:
- url: /.*
  script: _go_app
```
**NOTE**
Your API Key should be kept secret. If your ```app.yaml``` file contains your key, make sure to not check it into source control.

### Testing your app
The Google Cloud SDK has a tool for testing your app locally so that you can see how your app works before you actually deploy it.
If you installed the SDK correctly, you should have a tool called ```dev_appserver.py```. To run your app locally, just call ```dev_appserver.py app.yaml``` from the directory where your app is located. 
You should see something like the following in your terminal after running the command.
```
INFO     2017-10-15 17:47:30,338 devappserver2.py:105] Skipping SDK update check.
INFO     2017-10-15 17:47:30,362 api_server.py:300] Starting API server at: http://localhost:45941
INFO     2017-10-15 17:47:30,513 dispatcher.py:251] Starting module "default" running at: http://localhost:8080
INFO     2017-10-15 17:47:30,513 admin_server.py:116] Starting admin server at: http://localhost:8000
```
At this point, your app is running at ```localhost:8080``` and ready to test!
Let's test the app with ```curl```. To test this app, simply send a POST request to ```localhost:8080/send``` with a JSON body in the following format. 
```json 
{
  "email": "test@example.com",
  "name": "Example User"
}
```
The email/name field is the address and name of the recipient.
With ```curl```, the command would look like the following,

```curl -X POST -d '{"email": "test@example", "name": "Example User"}' localhost:8080/send```

If you see ```Successfully sent email``` in your terminal, the app sent the email successfully! 

If you have any trouble running the ```dev_appserver.py``` command, take a look at the [development server reference](https://cloud.google.com/appengine/docs/standard/go/tools/using-local-server).

### Deployment
After successfully testing the app, you are now ready to deploy!
To deploy, run the following command from the root of your app directory, ```gcloud app deploy app.yaml```. After running the command you should see something like the following.
```
Services to deploy:

descriptor:      [/path/to/yourApp/app.yaml]
source:          [/path/to/yourApp]
target project:  [YourProjectId]
target service:  [default]
target version:  [VersionId]
target url:      [https://YourProjectId.appspot.com]


Do you want to continue (Y/n)?
```
Keep note of the target url, that is where the app will be located.

Once you continue, the app will be uploaded and deployed to Google App Engine!

### Final notes
* Google App Engine does not allow the use of the ```net/http``` package, you must use [urlfetch](https://cloud.google.com/appengine/docs/standard/python/issue-requests) to issue HTTP(S) requests.
* In general it is bad practice to store keys in your code base. The best way would be to use [Cloud Key Management Service](https://cloud.google.com/kms/) which is a cloud-hosted key management service for Google Cloud Platform.