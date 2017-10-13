This documentation provides examples for specific use cases. Please [open an issue](https://github.com/sendgrid/sendgrid-go/issues) or make a pull request for any use cases you would like us to document here. Thank you!

# Table of Contents

* [Transactional Templates](#transactional_templates)
* [CustomArgs](#customargs)
* [Personalizations](#personalizations)
* [Substitutions](#substitutions)
* [Sections](#sections)
* [Attachments](#attachments)
* [Integrations - Slack Events](#slack)

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
#Integrations - Slack Events

This tutorial will go over how to setup a simple integration with Slack that allows an email to be triggered based on some defined event. Pre-requisites for this tutorial include:

* SendGrid API Key
* SendGrid Transactional Template
* Slack account & Valid Slack Workspace
* Slack API key
* ngrok (if you do not have an publically accessible web endpoint)

### Credential Storage

It's not a good practice to directly insert API keys (or any authorization information) into your script. As an alternative, please save your credentials in their respective `.env` files, with the following command:
```
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "export SLACK_API_KEY='YOUR_API_KEY'" > slack.env
echo "sendgrid.env" >> .gitignore
echo "slack.env" >> .gitignore
source ./sendgrid.env
source ./slack.env
```

### Example Overview

In this example we will explore how to setup an endpoint for the [Slack Events API](https://api.slack.com/events-api) that will listen for a specific event type, and then react by using SendGrid's mail send to send a customized message.

We will be listening for a [message event](https://api.slack.com/events/message) with the text `send starred test@email.com`, which will trigger an email sent to the specified address with all starred items as attachments.

### Slack Setup

Once you have a functional Slack workspace running, you'll need to:

* Create a Slack [app](https://api.slack.com/apps), and then click on it for further configuration.
* On the sidebar under Features select *Event Subscriptions*.
* Toggle "On" for Enable Events. We will insert a Request URL later.
* Under Subscribe to Workspace Events, click `Add Workspace Event` and add the [message.channels](https://api.slack.com/events/message.channels) event type.


### Setting up a Webhook

Slack's Event API requires a Request URL - an HTTP endpoint that is capable of accepting POST requests. Since we are assuming a development environment, we will be using [ngrok](https://ngrok.com/) to expose a locally available endpoint for testing purposes. Please follow installation instructions, and continue this tutorial once you are able to run `./ngrok http <port number>`.

First, let's setup a basic server and endpoint that will listen for and consume events in real-time:

```go
//starts server on given port at localhost.
func startServer(port string) error {
	portNum := fmt.Sprintf(":%v", port)
	fmt.Printf("Currently listening on http://localhost:%v ", portNum)
	err := http.ListenAndServe(portNum, nil)
	return err
}

//

```

Going back to the Event Subscriptions settings page for your Slack app, insert a valid request URL (if ngrok-generated - use the provided base URL appended with your endpoint slug e.g. http://d770fcbd.ngrok.io/post_events).

You'll be asked to have your endpoint return a response containing the value of `challenge` for verification purposes.

We'll need to model Slack's Event API JSON responses as Go structs in order to properly unmarshal them.

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

Now we can add code to our endpoint that will properly decode incoming request bodies, and then write a response containing the provided `challenger` value.

```go

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
  fmt.Print("RESP", resp)
  w.Header().Set("Content-Type", "application/json")
  w.Write(resp)
}

```

This should verify our request URL. Click "Save Changes".

### Sending Mail
