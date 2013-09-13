#SendGrid-Go

SendGrid Helper Library to send emails very easily, brah. **Still ALPHA.**

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
}

```

###TODO

* Add support for SMTP
* Write Tests

##MIT License
===
Enjoy. Feel free to make pull requests and stuff since I'm learning Go while I write this.
