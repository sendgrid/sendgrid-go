package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go/helpers/mock"
)

func main() {

	sendMail()       // Send Mail
	simpleSendMail() // Response Error

	// start mocks server
	mock.StartTestServer()

	// add mock value
	mock.Add(&mock.Mock{
		StatusCode: 400,
		Body:       `{ "errors":[{ "message":"Example error.", "field":"example field" }] }`,
	})
	sendMail()       // Response with mock data
	simpleSendMail() // Response with mock data

	// add mock value
	mock.Add(&mock.Mock{
		StatusCode: 200,
	})
	sendMail()       // Response with mock data
	simpleSendMail() // Response with mock data

	// add mock value
	mock.Add(&mock.Mock{
		Err: errors.New("Has been an mock error"),
	})

	// stop mocks server
	mock.StopTestServer()

	sendMail()       // Send Mail
	simpleSendMail() // Response Error
	// start mocks server
	mock.StartTestServer()

	sendMail()       // Response with mock data
	simpleSendMail() // Response with mock data

	// clear mock
	mock.Flush()
	sendMail()       // Send Mail
	simpleSendMail() // Response Error

	// stop mocks server
	mock.StopTestServer()

}

func sendMail() {
	m := mail.NewV3Mail()
	content := mail.NewContent("text/html", "<h1>Hello world</h1>This is an example")

	from := mail.NewEmail("envalidMail", "invalidMail@mail.com")
	m.SetFrom(from)

	m.AddContent(content)

	personalization := mail.NewPersonalization()

	personalization.AddTos(mail.NewEmail("envalidMail", "invalidMail@mail.com"))
	personalization.Subject = "Example"

	m.AddPersonalizations(personalization)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)

	if err != nil {
		fmt.Println("Sengird Error: ")
		fmt.Println(err)
		fmt.Println("________________________________")
		fmt.Println()
		return
	}

	fmt.Println("Sengird Response: ")
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	fmt.Println("________________________________")
	fmt.Println()
}

func simpleSendMail() {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClientMock(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Simple Sengird Error: ")
		fmt.Println(err)
		fmt.Println("________________________________")
		fmt.Println()
	} else {
		fmt.Println("Simple Sengird Response: ")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		fmt.Println("________________________________")
		fmt.Println()
	}
}
