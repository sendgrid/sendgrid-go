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

###TODO

* Write Tests
* More robust documentation
* Implement missing Mail API parameters
* Implement additional API endpoints

##MIT License

Enjoy. Feel free to make pull requests and stuff since I'm learning Go while I write this.
