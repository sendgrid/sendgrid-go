# Special Announcement

We have released a [v3 beta branch](https://github.com/sendgrid/sendgrid-go/tree/v3beta) for this library that supports our new v3 Mail Send endpoint which is in open beta. The v3/mail/send/beta endpoint is not a production endpoint, so you should not integrate with it for your production email sending. However, when we make this an officially released feature it will be available at v3/mail/send.

Please try it out and let us know what you think about the endpoint and the library in the [issues area of this repo](https://github.com/sendgrid/sendgrid-go/issues), all of your feedback will be taken into account to influence the endpoint and this library.

Beginning with v3/mail/send/beta, the new version of our library will only support v3 endpoints.. Once this endpoint is out of beta, we will update the endpoint, removing the “/beta” from the URI. At this point, the v3 beta branch will be merged to master and will be our official library going forward. This means that we will no longer formally support the v2 mail.send.json endpoint in any of our libraries.

So long as you are not automatically pulling new versions of the library into your production code base, your integration will not break regardless of which endpoint you’re using. By the way, don't pull new versions into your production code base, because breaking changes break things.

The /api/mail.send.json endpoint, known as v2 mail send, is NOT going away. It will continue to work as it always has, happily sending your emails along as if nothing happened.

# SendGrid-Go
[![GoDoc](https://godoc.org/github.com/sendgrid/sendgrid-go?status.png)](http://godoc.org/github.com/sendgrid/sendgrid-go) 
Visit the GoDoc.

[![Build Status](https://travis-ci.org/sendgrid/sendgrid-go.svg?branch=master)](https://travis-ci.org/sendgrid/sendgrid-go)
SendGrid Helper Library to send emails very easily using Go.

### Warning

Version ``2.x.x`` drops support for Go versions < 1.3.

Version ``1.2.x`` behaves differently in the ``AddTo`` method. In the past this method defaulted to using the ``SMTPAPI`` header. Now you must explicitly call the ``SMTPAPIHeader.AddTo`` method. More on the ``SMTPAPI`` section.

## Installation

```bash
go get github.com/sendgrid/sendgrid-go

# Or pin the version with gopkg
go get gopkg.in/sendgrid/sendgrid-go.v1

echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore

```

## Example

```go
package main

import (
        "fmt"
        "github.com/sendgrid/sendgrid-go"
		"os"
)

func main() {
	sendgridKey := os.Getenv("SENDGRID_API_KEY")
	if sendgridKey == "" {
	  			fmt.Println("Environment variable SENDGRID_API_KEY is undefined. Did you forget to source sendgrid.env?")
	  			os.Exit(1);
	}
    sg := sendgrid.NewSendGridClientWithApiKey(sendgridKey)
    message := sendgrid.NewMail()
    message.AddTo("community@sendgrid.com")
    message.AddToName("SendGrid Community Dev Team")
    message.SetSubject("SendGrid Testing")
    message.SetText("WIN")
    message.SetFrom("you@yourdomain.com")
    if r := sg.Send(message); r == nil {
                fmt.Println("Email sent!")
        } else {
                fmt.Println(r)
        }
}
```

## Usage

To begin using this library, call `NewSendGridClientWithApiKey` with a SendGrid API Key.  

### Creating a Client

```go
sg := sendgrid.NewSendGridClientWithApiKey("sendgrid_api_key")
```

### Creating a Mail
```go
message := sendgrid.NewMail()
```

### Adding Recipients

```go
message.AddTo("example@sendgrid.com") // Returns error if email string is not valid RFC 5322
// or
address, _ := mail.ParseAddress("Example <example@sendgrid.com>")
message.AddRecipient(address) // Receives a vaild mail.Address
```

### Adding BCC Recipients

Same concept as regular recipient excepts the methods are:

*   AddBcc
*   AddBccRecipient

### Setting the Subject

```go
message.SetSubject("New email")
```

### Set Text or HTML

```go
message.SetText("Add Text Here..")
//or
message.SetHTML("<html><body>Stuff, you know?</body></html>")
```
### Set From

```go
message.SetFrom("example@lol.com")
```
### Set File Attachments

```go
message.AddAttachment("text.txt", file) // file needs to implement the io.Reader interface
//or
message.AddAttachmentFromStream("filename", "some file content")
```
### Adding ContentIDs

```go
message.AddContentID("id", "content")
```

## SendGrid's  [X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/)

If you wish to use the X-SMTPAPI on your own app, you can use the [SMTPAPI Go library](https://github.com/sendgrid/smtpapi-go).


### Recipients

```go
message.SMTPAPIHeader.AddTo("addTo@mailinator.com")
// or
message.SMTPAPIHeader.AddTo("test@test.com", "test@email.com")
// or
message.SMTPAPIHeader.SetTos(tos)
```

### [Substitutions](http://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html)

```go
message.AddSubstitution("key", "value")
// or
values := []string{"value1", "value2"}
message.AddSubstitutions("key", values)
//or
sub := make(map[string][]string)
sub["key"] = values
message.SetSubstitutions(sub)
```

### [Section](http://sendgrid.com/docs/API_Reference/SMTP_API/section_tags.html)

```go
message.AddSection("section", "value")
// or
sections := make(map[string]string)
sections["section"] = "value"
message.SetSections(sections)
```

### [Category](http://sendgrid.com/docs/Delivery_Metrics/categories.html)

```go
message.AddCategory("category")
// or
categories := []string{"setCategories"}
message.AddCategories(categories)
// or
message.SetCategories(categories)
```

### [Unique Arguments](http://sendgrid.com/docs/API_Reference/SMTP_API/unique_arguments.html)

```go
message.AddUniqueArg("key", "value")
// or
args := make(map[string]string)
args["key"] = "value"
message.SetUniqueArgs(args)
```

### [Filters](http://sendgrid.com/docs/API_Reference/SMTP_API/apps.html)

```go
message.AddFilter("filter", "setting", "value")
// or
filter := &Filter{
  Settings: make(map[string]string),
}
filter.Settings["enable"] = "1"
filter.Settings["text/plain"] = "You can haz footers!"
message.SetFilter("footer", filter)
```

### JSONString

```go
message.JSONString() //returns a JSON string representation of the headers
```

Kudos to [Matthew Zimmerman](https://github.com/mzimmerman) for this example.

###Tests

Please run the test suite in before sending a pull request.

```bash
go test -v
```

### TODO:
* Add Versioning
* Add proper support for BCC

##MIT License

Enjoy. Feel free to make pull requests :)
