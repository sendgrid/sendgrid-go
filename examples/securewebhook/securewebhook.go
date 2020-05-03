package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/securewebhook"
)

// EnableSecureWebhooksettings : Enables Secure Webhook.
// PATCH /user/webhooks/event/settings/signed
func EnableSecureWebhooksettings() {
	var err error
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/event/settings/signed", host)
	request.Method = rest.Patch
	s := securewebhook.NewSettings()
	s.SetEnable(true)
	request.Body, err = securewebhook.GetRequestBody(s)
	if err != nil {
		log.Println(err)
		return
	}
	response, err := sendgrid.MakeRequest(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// GetPublicKeyForWebhook : Get Public Key if SecureWebhook feature is enabled.
// Get /user/webhooks/event/settings/signed
func GetPublicKeyForWebhook() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/user/webhooks/event/settings/signed", host)
	request.Method = rest.Get
	response, err := sendgrid.MakeRequest(request)
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
