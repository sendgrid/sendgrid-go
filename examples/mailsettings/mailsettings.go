package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

// GetAllMailSettings : Retrieves all mail settings
// GET /mail_settings
func GetAllMailSettings() {
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

// UpdateAddressWhitelistMailSettings : Update address whitelist mail settings
// PATCH /mail_settings/address_whitelist
func UpdateAddressWhitelistMailSettings() {
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

// GetAddressWhitelistMailSettings : Retrieve address whitelist mail settings
// GET /mail_settings/address_whitelist
func GetAddressWhitelistMailSettings() {
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

// UpdateBCCMailSettings : Update BCC mail settings
// PATCH /mail_settings/bcc
func UpdateBCCMailSettings() {
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

// GetBCCMailSettings : Retrieve all BCC mail settings
// GET /mail_settings/bcc
func GetBCCMailSettings() {
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

// UpdateBouncePurgeMailSettings : Update bounce purge mail settings
// PATCH /mail_settings/bounce_purge
func UpdateBouncePurgeMailSettings() {
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

// GetBouncePurgeMailSettings : Retrieve bounce purge mail settings
// GET /mail_settings/bounce_purge
func GetBouncePurgeMailSettings() {
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

// UpdateFooterMailSettings : Updates footer mail settings
// PATCH /mail_settings/footer
func UpdateFooterMailSettings() {
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

// GetFooterMailSettings : Retrieve footer mail settings
// GET /mail_settings/footer
func GetFooterMailSettings() {
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

// UpdateForwardBounceMailSettings : Update forward bounce mail settings
// PATCH /mail_settings/forward_bounce
func UpdateForwardBounceMailSettings() {
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

// GetForwardBounceMailSettings : Retrieve forward bounce mail settings
// GET /mail_settings/forward_bounce
func GetForwardBounceMailSettings() {
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

// UpdateForwardSpamMailSettings : Update forward spam mail settings
// PATCH /mail_settings/forward_spam
func UpdateForwardSpamMailSettings() {
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

// GetForwardSpamMailSettings : Retrieve forward spam mail settings
// GET /mail_settings/forward_spam
func GetForwardSpamMailSettings() {
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

// UpdatePlainContentMailSettings : Update plain content mail settings
// PATCH /mail_settings/plain_content
func UpdatePlainContentMailSettings() {
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

// GetPlainContentMailSettings : Retrieve plain content mail settings
// GET /mail_settings/plain_content
func GetPlainContentMailSettings() {
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

// UpdateSpamCheckMailSettings : Update spam check mail settings
// PATCH /mail_settings/spam_check
func UpdateSpamCheckMailSettings() {
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

// GetSpamCheckMailSettings : Retrieve spam check mail settings
// GET /mail_settings/spam_check
func GetSpamCheckMailSettings() {
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

// UpdateTemplateMailSettings : Update template mail settings
// PATCH /mail_settings/template
func UpdateTemplateMailSettings() {
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

// GetLegacyTemplateMailSettings:  Retrieves legacy template mail settings
// GET /mail_settings/template
func GetLegacyTemplateMailSettings() {
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
