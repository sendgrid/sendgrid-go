This documentation provides examples for specific use cases. Please [open an issue](https://github.com/sendgrid/sendgrid-go/issues) or make a pull request for any use cases you would like us to document here. Thank you!

# Table of Contents

* [Transactional Templates](#transactional_templates)
* [CustomArgs](#customargs)
* [Personalizations](#personalizations)
* [Substitutions](#substitutions)
* [Sections](#sections)
* [Attachments](#attachments)
* [Integrations - Slack Events](#slack-events)

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

<a name="slack-events"></a>
# Integrations - Slack Events

This tutorial will go over how to setup a simple integration with Slack that allows an email to be triggered based on some defined event. You'll need the following to complete this tutorial.

* SendGrid API Key
* Slack account & Valid Slack Workspace
* ngrok (if you do not have an publicly accessible web endpoint)

The following example will explore the basic setup of an endpoint for the [Slack Events API](https://api.slack.com/events-api) that will listen for a specific [event type](https://api.slack.com/events), and then react by using SendGrid's mail send to send a customized message.

We will be listening for a [message event](https://api.slack.com/events/message) with the text `sendgrid|test email body.|from@example.com|to@example.com`, which will trigger an email sent from a specified address to another specified address.

### Credential Storage

It's not a good practice to directly insert API keys (or any authorization information) into your script. As an alternative, please save your credentials in their respective `.env` files, with the following command:
```
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
source ./sendgrid.env
```

### Slack Setup

Once you have a functional Slack workspace running, you'll need to:

1. Create a Slack [app](https://api.slack.com/apps), and then click on it for further configuration.

2. On the sidebar under Features select *Event Subscriptions*.

3. Toggle **On** for Enable Events. We will insert a Request URL later.

4. Under Subscribe to Workspace Events, click **Add Workspace Event** and add the [message.channels](https://api.slack.com/events/message.channels) event type.

Once you've setup your new app the following the steps above, you can now install it to any given workspace. At that point you're done! No Slack API key/OAuth setup required - all authentication/scope issues are taken care of internally when you install an app to a workspace from the UI.

### Setting up a Webhook

Now that you have an installed Slack app configured with **Event Subscriptions** turned **On**, we need to setup a server and a web endpoint that accepts POST requests. Once we have a functional endpoint, we will insert the endpoint URL into the field *Request URL* on the **Events Subscriptions** page.

Let's setup a basic server that will run on a specified port:

```go
//this import contains all dependencies used in the example.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
	"os"
	"strings"
)

//starts server on given port on localhost.
func startServer(port string) error {
	portNum := fmt.Sprintf(":%v", port)
	fmt.Println("Currently listening on localhost @ port ", portNum)
	err := http.ListenAndServe(portNum, nil)
	return err
}
```

Before you can save a new endpoint URL into the **Request URL** field, you'll be asked to have your endpoint return an HTTP response containing the value of `challenge` for verification purposes.

We'll need to model Slack's Event API JSON responses as well as the verification JSON object as Go structs in order to properly unmarshal them, and decode raw data we'll be receiving from the endpoint.

```go
//SlackPostVerification contains attributes that Slack needs to verify URL works.
type SlackPostVerification struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

//Event encapsulates actual data we want from that posted to the webhook.
type Event struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Ts      string `json:"event_ts"`
	Type    string `json:"type"`
	User    string `json:"user"`
}

//SlackMessageEvent represents JSON structure of a Slack message event.
type SlackMessageEvent struct {
	SlackPostVerification
	Event Event `json:"event"`
}
```

With these models, we can now build our endpoint. We'll want it to:

1. Only accept POST requests.

2. Decode raw JSON data into useful structs.

```go
//slackEventEndpoint accepts POST requests, which will contain slack events.
func slackEventEndpoint(w http.ResponseWriter, req *http.Request) {
	// only accept POST requests
	if req.Method != http.MethodPost {
		w.WriteHeader(405)
	} else {

		// translate body to JSON
		decoder := json.NewDecoder(req.Body)
		var slackMessage SlackMessageEvent
		if err := decoder.Decode(&slackMessage); err != nil {
			log.Fatal(err)
		}

		// return challenge as part of response.
		var response SlackPostVerification
		response.Challenge = slackMessage.Challenge
		if resp, err := json.Marshal(response); err != nil {
			log.Fatal(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
		}
		// parse consumed message - send Email based on params.
	}
}
```

With this functional endpoint, it's time to associate it with our server. We'll do this in the `main` function.

```go
func main() {
	http.HandleFunc("/slack_events", slackEventEndpoint) // <your_base_url>/post_events will be our endpoint.
	err := startServer("8080")
	if err != nil {
		log.Println(err)
	}
}
```

Now when our server runs, `/post_events` can receive incoming POST requests.

With all these pieces in place, it's time to insert our request URL into Slack. Since we are assuming a development environment, we will be using [ngrok](https://ngrok.com/) to expose a locally available endpoint.

Please follow installation [instructions](https://ngrok.com/download), and continue this tutorial once you are able to successfully run `./ngrok http <port number>`. Running this command will create a publicly accessible URL that tunnels to your localhost at the specified port, e.g. `http://a123bcde.ngrok.io`.

Just insert `<your_base_url>/post_events` into Slack, and make sure it verifies in the UI (`challenge` value was accepted). Event data should now be sent to your endpoint when a message is submitted on a channel in your Slack workspace.

### Sending Mail based on Events

Now that we've laid out the infrastructure (Slack setup, server, endpoint), we can start handling incoming Slack events. In this case we are listening for the [message.channels](https://api.slack.com/events/message.channels) event type, and in the previous section we created Go structs to help us decode this event data.

We need to now add code that will:

1. Handle a `message.channels` event and parse out individual text components. We expect the format of a command to be `sendgrid|test email body.|from@example.com|to@example.com`.

2. Build a mail object using the SendGrid-Go API library.

3. Send the email.

```go
//slackEventEndpoint accepts POST requests, which will contain slack events.
func slackEventEndpoint(w http.ResponseWriter, req *http.Request) {
	// only accept POST requests
	if req.Method != http.MethodPost {
		w.WriteHeader(405)
	} else {

		// translate body to JSON
		decoder := json.NewDecoder(req.Body)
		var slackMessage SlackMessageEvent
		if err := decoder.Decode(&slackMessage); err != nil {
			log.Fatal(err)
		}

		// return challenge as part of response.
		var response SlackPostVerification
		response.Challenge = slackMessage.Challenge
		if resp, err := json.Marshal(response); err != nil {
			log.Fatal(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
		}
		// parse consumed message - send Email based on params.
		eventParsed := parseMessageEvent(slackMessage.Event)
		if eventParsed[0] == "sendgrid" {
			if len(eventParsed) == 4 {
				m := convertEventMail(eventParsed)
				sendSGMessage(m)
			}
		}
	}
}

//parseMessageEvent reads the text field of `Event`, and splits it into
// ["<cmd>", "<body>", "<from_address>", "<to_address>"]
func parseMessageEvent(e Event) []string {
    eventParsed := strings.Split(e.Text, "|")
    return eventParsed
}

//convertEventMail takes relevant info from Event and populates a mail object.
func convertEventMail(eventParsed []string) *mail.SGMailV3 {
  m := mail.NewV3Mail()
  from := mail.NewEmail("From Slack Channel", eventParsed[2])
  content := mail.NewContent("text/html", eventParsed[1])
  to := mail.NewEmail("", eventParsed[3])
  p := mail.NewPersonalization()

  m.SetFrom(from)
  m.AddContent(content)
  p.AddTos(to)
  p.Subject = "Message sent from Slack!"
  m.AddPersonalizations(p)
  return m
}

//sendSGMessage wraps SendGrid library
func sendSGMessage(m *mail.SGMailV3) {
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
  return
}
```

### Conclusion

Now when we submit a message of the format `sendgrid|test email body.|from@example.com|to@example.com` in any channel on our Slack workspace, our endpoint will process the event data and then send an email using SendGrid based on these text arguments.

We covered a basic integration - but there's many different [event types](https://api.slack.com/events/api) that Slack's API offers. The possibilities are plentiful!
