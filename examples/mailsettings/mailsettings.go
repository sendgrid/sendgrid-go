package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve all mail settings
// GET /mail_settings

func Retrieveallmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["limit"] = "1"
queryParams["offset"] = "1"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update address whitelist mail settings
// PATCH /mail_settings/address_whitelist

func Updateaddresswhitelistmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/address_whitelist", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": true, 
  "list": [
    "email1@example.com", 
    "example.com"
  ]
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve address whitelist mail settings
// GET /mail_settings/address_whitelist

func Retrieveaddresswhitelistmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/address_whitelist", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update BCC mail settings
// PATCH /mail_settings/bcc

func UpdateBCCmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/bcc", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "email": "email@example.com", 
  "enabled": false
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all BCC mail settings
// GET /mail_settings/bcc

func RetrieveallBCCmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/bcc", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update bounce purge mail settings
// PATCH /mail_settings/bounce_purge

func Updatebouncepurgemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/bounce_purge", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": true, 
  "hard_bounces": 5, 
  "soft_bounces": 5
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve bounce purge mail settings
// GET /mail_settings/bounce_purge

func Retrievebouncepurgemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/bounce_purge", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update footer mail settings
// PATCH /mail_settings/footer

func Updatefootermailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/footer", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": true, 
  "html_content": "...", 
  "plain_content": "..."
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve footer mail settings
// GET /mail_settings/footer

func Retrievefootermailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/footer", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update forward bounce mail settings
// PATCH /mail_settings/forward_bounce

func Updateforwardbouncemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/forward_bounce", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "email": "example@example.com", 
  "enabled": true
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve forward bounce mail settings
// GET /mail_settings/forward_bounce

func Retrieveforwardbouncemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/forward_bounce", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update forward spam mail settings
// PATCH /mail_settings/forward_spam

func Updateforwardspammailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/forward_spam", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "email": "", 
  "enabled": false
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve forward spam mail settings
// GET /mail_settings/forward_spam

func Retrieveforwardspammailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/forward_spam", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update plain content mail settings
// PATCH /mail_settings/plain_content

func Updateplaincontentmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/plain_content", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": false
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve plain content mail settings
// GET /mail_settings/plain_content

func Retrieveplaincontentmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/plain_content", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update spam check mail settings
// PATCH /mail_settings/spam_check

func Updatespamcheckmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/spam_check", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": true, 
  "max_score": 5, 
  "url": "url"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve spam check mail settings
// GET /mail_settings/spam_check

func Retrievespamcheckmailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/spam_check", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update template mail settings
// PATCH /mail_settings/template

func Updatetemplatemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/template", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "enabled": true, 
  "html_content": "<% body %>"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve legacy template mail settings
// GET /mail_settings/template

func Retrievelegacytemplatemailsettings() {
  apiKey := os.Getenv("YOUR_SENDGRID_APIKEY")
  host := "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/mail_settings/template", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

func main() {
    // add your function calls here
}
