#SendGrid-Go

SendGrid Helper Library to send emails very easily, brah. 
SMTP has been added has a fallback. Basically invoking **Send** will call **SendAPI**. If that fails, **SendSMTP** will be invoked as a fallback.
 
**Still ALPHA.**

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
	message.AddTo("ya@adcade.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("Go Lib is alive")
	message.AddHTML("Still needs work and SMTP")
	message.AddFrom("ya@adcade.com")
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
	message.AddTo("ya@adcade.com")
	message.AddToName("Yamil Asusta")
	message.AddSubject("Go Lib is alive")
	message.AddHTML("Still needs work and SMTP")
	message.AddFrom("ya@adcade.com")
	// use sendAPI instead of SMTP
	if r := sg.SendAPI(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		c.Errorf("Unable to send mail %v",r)
	}
}

```

###TODO

* Write Tests
* More robust documentation
* Implement missing Mail API parameters
* Implement additional API endpoints

##MIT License

Enjoy. Feel free to make pull requests and stuff since I'm learning Go while I write this.
