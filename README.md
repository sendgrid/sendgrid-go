#SendGrid-Go

SendGrid Helper Library to send emails very easily using Go.

##Installation

```bash
go get github.com/elbuo8/sendgrid-go/sendgrid
```

##Example

```Go
package main

import (
	"fmt"
	"github.com/elbuo8/sendgrid-go/sendgrid"
)

func main() {
	sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
	message := sendgrid.NewMail()
	message.AddTo("yamil@sendgrid.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("SendGrid is Baller")
	message.AddHTML("Simple Text")
	message.AddFrom("kunal@sendgrid.com")
	if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
	/*
	Additional ways to interface with the client
	sg.SendAPI(message)
	sg.SendSMTP(message)
	*/
}

```

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


###Notes

SMTP has been added has a fallback. Basically invoking **Send** will call **SendAPI**. If that fails, **SendSMTP** will be invoked as a fallback. SMTP is still missing a few features; therefore it is recommended that **SendAPI** & **Send** are the primary methods.


###TODO

* Write Tests
* Finish implementing SMTP
* Implement additional API endpoints

##MIT License

Enjoy. Feel free to make pull requests :)
