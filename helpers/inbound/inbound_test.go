package inbound

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRequest(filename string) *http.Request {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	// Build POST request
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(file))
	req.Header.Set("Content-Type", "multipart/form-data; boundary=xYzZY")
	req.Header.Set("User-Agent", "Twilio-SendGrid-Test")
	return req
}

func TestParse(t *testing.T) {
	// Build a table of tests to run with each one having a name, the sample data file to post,
	// and the expected HTTP response from the handler
	tests := []struct {
		name             string
		file             string
		expectedResponse int
	}{
		{"NoAttachment", "./sample_data/raw_data.txt", http.StatusOK},
		{"Attachment", "./sample_data/raw_data_with_attachments.txt", http.StatusOK},
		{"DefaultData", "./sample_data/default_data.txt", http.StatusOK},
	}

	for _, test := range tests {
		t.Run(test.name, func(subTest *testing.T) {
			//Load POST body
			req := createRequest(test.file)

			// Invoke callback handler
			email := Parse(req)
			from := "Example User <test@example.com>"
			assert.Equalf(subTest, email.Headers["From"], from, "Expected From: %s, Got: %s", from, email.Headers["From"])
		})
	}
}

func TestAttachments(t *testing.T) {
	req := createRequest("./sample_data/raw_data_with_attachments.txt")
	email := Parse(req)
	contentType := "image/jpeg"

	assert.Equalf(t, contentType, email.Attachments[0].ContentType,"Expected From: %s, Got: %s", contentType, email.Attachments[0].ContentType)
}

func ExampleParsedEmail_parseHeaders() {
	headers := `
Foo: foo
Bar: baz
`
	email := ParsedEmail{
		Headers:     make(map[string]string),
		Body:        make(map[string]string),
		Attachments: []ParsedAttachment{},
		rawRequest:  nil,
	}
	email.parseHeaders(headers)
	fmt.Println(email.Headers["Foo"])
	fmt.Println(email.Headers["Bar"])
	// Output:
	// foo
	// baz
}

func ExampleParsedEmail_parseRawEmail() {
	rawEmail := `
To: test@example.com
From: example@example.com
Subject: Test Email
Content-Type: multipart/mixed; boundary=TwiLIo

--TwiLIo
Content-Type: text/plain; charset=UTF-8

Hello Twilio SendGrid!
--TwiLIo
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

<html><body><strong>Hello Twilio SendGrid!</body></html>
--TwiLIo--
`
	email := ParsedEmail{
		Headers:     make(map[string]string),
		Body:        make(map[string]string),
		Attachments: []ParsedAttachment{},
		rawRequest:  nil,
	}
	email.parseRawEmail(rawEmail)
	for key, value := range email.Headers {
		fmt.Println(key, value)
	}
	fmt.Println(email.Body["text/plain; charset=UTF-8"])
	// Unordered Output:
	// To test@example.com
	// From example@example.com
	// Subject Test Email
	// Content-Type multipart/mixed; boundary=TwiLIo
	// Hello Twilio SendGrid!
}