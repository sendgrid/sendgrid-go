package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// Retrieveallmailsettings : Retrieve all mail settings
// GET /mail_settings
func Retrieveallmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateaddresswhitelistmailsettings : Update address whitelist mail settings
// PATCH /mail_settings/address_whitelist
func Updateaddresswhitelistmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/address_whitelist", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true, 
  "list": [
    "email1@example.com", 
    "example.com"
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

// Retrieveaddresswhitelistmailsettings : Retrieve address whitelist mail settings
// GET /mail_settings/address_whitelist
func Retrieveaddresswhitelistmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/address_whitelist", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// UpdateBCCmailsettings : Update BCC mail settings
// PATCH /mail_settings/bcc
func UpdateBCCmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/bcc", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "email@example.com", 
  "enabled": false
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

// RetrieveallBCCmailsettings : Retrieve all BCC mail settings
// GET /mail_settings/bcc
func RetrieveallBCCmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/bcc", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updatebouncepurgemailsettings : Update bounce purge mail settings
// PATCH /mail_settings/bounce_purge
func Updatebouncepurgemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/bounce_purge", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true, 
  "hard_bounces": 5, 
  "soft_bounces": 5
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

// Retrievebouncepurgemailsettings : Retrieve bounce purge mail settings
// GET /mail_settings/bounce_purge
func Retrievebouncepurgemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/bounce_purge", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updatefootermailsettings : Update footer mail settings
// PATCH /mail_settings/footer
func Updatefootermailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/footer", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true, 
  "html_content": "...", 
  "plain_content": "..."
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

// Retrievefootermailsettings : Retrieve footer mail settings
// GET /mail_settings/footer
func Retrievefootermailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/footer", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateforwardbouncemailsettings : Update forward bounce mail settings
// PATCH /mail_settings/forward_bounce
func Updateforwardbouncemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/forward_bounce", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "example@example.com", 
  "enabled": true
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

// Retrieveforwardbouncemailsettings : Retrieve forward bounce mail settings
// GET /mail_settings/forward_bounce
func Retrieveforwardbouncemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/forward_bounce", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateforwardspammailsettings : Update forward spam mail settings
// PATCH /mail_settings/forward_spam
func Updateforwardspammailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/forward_spam", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "", 
  "enabled": false
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

// Retrieveforwardspammailsettings : Retrieve forward spam mail settings
// GET /mail_settings/forward_spam
func Retrieveforwardspammailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/forward_spam", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateplaincontentmailsettings : Update plain content mail settings
// PATCH /mail_settings/plain_content
func Updateplaincontentmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/plain_content", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": false
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

// Retrieveplaincontentmailsettings : Retrieve plain content mail settings
// GET /mail_settings/plain_content
func Retrieveplaincontentmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/plain_content", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updatespamcheckmailsettings : Update spam check mail settings
// PATCH /mail_settings/spam_check
func Updatespamcheckmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/spam_check", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true, 
  "max_score": 5, 
  "url": "url"
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

// Retrievespamcheckmailsettings : Retrieve spam check mail settings
// GET /mail_settings/spam_check
func Retrievespamcheckmailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/spam_check", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updatetemplatemailsettings : Update template mail settings
// PATCH /mail_settings/template
func Updatetemplatemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/template", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true, 
  "html_content": "<% body %>"
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

// Retrievelegacytemplatemailsettings : Retrieve legacy template mail settings
// GET /mail_settings/template
func Retrievelegacytemplatemailsettings() {
	apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/mail_settings/template", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// add your function calls here
}
