# SendGrid-Go

SendGrid Helper Library to send emails very easily using Go.

This library is built on top of [smtpmail](https://github.com/elbuo8/smtpmail) which provides a simple way to send emails using generic SMTP transport.

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

//or

message.AddHTML("<html><body>Stuff, you know?</body></html>")
```
### Set From

```Go
message := sendgrid.NewMail()

message.addFrom("example@lol.com")
```
### Set File Attachments

```Go
message := sendgrid.NewMail()

message.AddAttachment("./stuff.txt")
```

## SendGrid's  [X-SMTPAPI](http://sendgrid.com/docs/API_Reference/SMTP_API/)

### [Substitution](http://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html)

```Go
message := sendgrid.NewMail()

message.AddSubstitution("key", "value")
```

### [Section](http://sendgrid.com/docs/API_Reference/SMTP_API/section_tags.html)

```Go
message := sendgrid.NewMail()

message.AddSection("section", "value")
```

### [Category](http://sendgrid.com/docs/Delivery_Metrics/categories.html)

```Go
message := sendgrid.NewMail()

message.AddSubstitution("category")
```

### [Unique Arguments](http://sendgrid.com/docs/API_Reference/SMTP_API/unique_arguments.html)

```Go
message := sendgrid.NewMail()

message.AddUniqueArg("key", "value")
```

### [Filter](http://sendgrid.com/docs/API_Reference/SMTP_API/apps.html)

```Go
message := sendgrid.NewMail()

message.AddFilter("filter", "setting", "value")
```


### Sending Methods

There are 3 ways to send emails using the library. 

* 	SendAPI
* 	SendSMTP
* 	Send 


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

* Implement additional API endpoints

##MIT License

Enjoy. Feel free to make pull requests :)
