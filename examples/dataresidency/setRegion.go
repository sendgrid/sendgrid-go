package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sendgrid/sendgrid-go"
)

var SAMPLE_EMAIL = "test@example.com"

// SetDataResidency : Set region for sendgrid.
func SetDataResidencyGlobal() {
	message := buildHelloEmail()
	request, err := buildSendgridObj("global")
	if err != nil {
		log.Println(err)
	} else {
		request.Body = mail.GetRequestBody(message)
		response, err := sendgrid.API(request)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}
}

func SetDataResidencyEu() {
	message := buildHelloEmail()
	request, err := buildSendgridObj("eu")
	if err != nil {
		log.Println(err)
	} else {
		request.Body = mail.GetRequestBody(message)
		response, err := sendgrid.API(request)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}
}

func SetDataResidencyDefault() {
	message := buildHelloEmail()
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(message)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func buildHelloEmail() *mail.SGMailV3 {
	// Note that when you use this constructor an initial personalization object
	// is created for you. It can be accessed via
	// mail.personalization.get(0) as it is a List object

	from := mail.NewEmail("test_user", SAMPLE_EMAIL)
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("test_user", SAMPLE_EMAIL)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	email := mail.NewEmail("test_user", SAMPLE_EMAIL)

	p := mail.NewPersonalization()
	p.AddTos(email)
	message.AddPersonalizations(p)

	return message
}

func buildSendgridObj(region string) (rest.Request, error) {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "")
	request.Method = "POST"
	request, err := sendgrid.SetDataResidency(request, region)
	if err != nil {
		return request, err
	}
	return request, nil
}

func main() {
	// add your function calls here
}
