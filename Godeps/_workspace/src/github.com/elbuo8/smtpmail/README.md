# SMTPMail

Simple SMTP interface that supports file attachments and different text types.

## Install

```bash
go get github.com/elbuo8/smtpmail
```

## Example

```Go

import (
  "fmt"
  "github.com/elbuo8/smtpmail"
  "os"
)

func main() {
  c := smtpmail.NewSMTPClient(os.Getenv("SG_USER"), os.Getenv("SG_PWD"), "smtp.sendgrid.net", "587")
  m := smtpmail.NewMail()
  m.AddTo("hello@yamilasusta.com")
  m.AddFrom("yamil@sendgrid.com")
  m.AddSubject("Testing")
  m.AddText("Some text :)")
  filepath, _ := os.Getwd()
  m.AddAttachment(filepath + "/mail.go")
  if e := c.Send(m); e != nil {
    fmt.Println(e)
  } else {
    fmt.Println("All good brah")
  }
}

```

Pretty simple eh?

## Available Fields

```Go
To       []string
ToName   []string
Subject  string
HTML     string
Text     string
From     string
Bcc      []string
FromName string
ReplyTo  string
Date     string
Files    map[string]string
Headers  string
```

## MIT License
