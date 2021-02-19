## How to Setup a Domain Authentication

You can find documentation for how to setup a domain authentication via the UI [here](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-domain-authentication/).
Find more information about all of Twilio SendGrid's authentication related documentation [here](https://sendgrid.com/docs/ui/account-and-settings/).

To create a Domain authentication Via the API:
```
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func main() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/whitelabel/domains", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "automatic_security": false, 
  "custom_spf": true, 
  "default": true, 
  "domain": "example.com", 
  "ips": [
    "192.168.1.1", 
    "192.168.1.2"
  ], 
 "subdomain": "SUBDOMAIN", 
 "username": "YOUR_SENDGRID_SUBUSER_NAME"
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
