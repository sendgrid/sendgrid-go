#SendGrid-Go

SendGrid Helper Library to send emails very easily using Go.

##Installation

```bash
go get github.com/sendgrid/sendgrid-go
```

##Example

```Go
package main

import (
	"fmt"
	"net/mail"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
	message := sendgrid.NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("SendGrid Testing")
	message.AddText("WIN")
	message.AddFrom("yamil@sendgrid.com")
    if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
}

```

### Adding Receipients

```Go
message := sendgrid.NewMail()
message.AddTo("example@sendgrid.com") // Returns error if email string is not valid RFC 5322
 
// or

address, _ := mail.ParseAddress("Example <example@sendgrid.com>")
message.AddReceipient(address) // Receives a vaild mail.Address
```

### Adding BCC Receipients

Same concept as regular receipients excepts the methods are:

*	AddBCC
* 	AddReceipientBCC

### Setting the Subject

```Go
message := sendgrid.NewMail()

message.AddSubject("New email")
```

### Set Text or HTML

```Go
message := sendgrid.NewMail()

message.AddText("Body")
```
### Set From

```Go
message := sendgrid.NewMail()

message.AddHTML("<html><body>Stuff, you know?</body></html>")
```
### Set Custom Headers

```Go
message := sendgrid.NewMail()
message.AddHeader("X-Mailer", "Header")
```
### Set File Attachments

```Go
message := sendgrid.NewMail()
message.AddAttachment("./stuff.txt")
```
### Sending Methods

There are 3 ways to send emails using the library. 

* 	SendAPI
* 	SendSMTP
* 	Send 

Send will try to use **SendAPI** first. If this method fails, it will fail silently and will invoke **SendSMTP** has a fallback. If this also fails, an []error will be returned with the feedback.

##AppEngine Example

```Go
package main

import (
	"fmt"
	"appengine/urlfetch"
	"github.com/elbuo8/sendgrid-go/sendgrid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
	c := appengine.NewContext(r)
	// set http.Client to use the appengine client
	sg.Client = urlfetch.Client(c)
	message := sendgrid.NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("SendGrid is Baller")
	message.AddHTML("Simple Text")
	message.AddFrom("kunal@sendgrid.com")
	// use sendAPI instead of SMTP
	if r := sg.SendAPI(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		c.Errorf("Unable to send mail %v",r)
	}
}

```

###Tests

Please run it before sending pull requests

```bash
go test
```

###TODO

* Finish implementing SMTP
* Implement additional API endpoints

##MIT License

Enjoy. Feel free to make pull requests :)
