package inbound

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRequest(filename string) *http.Request {
	file, err := os.ReadFile(filename)
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
		name          string
		file          string
		expectedError error
	}{
		{
			name: "NoAttachment",
			file: "./sample_data/raw_data.txt",
		},
		{
			name: "Attachment",
			file: "./sample_data/raw_data_with_attachments.txt",
		},
		{
			name: "DefaultData",
			file: "./sample_data/default_data.txt",
		},
		{
			name:          "BadData",
			file:          "./sample_data/bad_data.txt",
			expectedError: fmt.Errorf("multipart: NextPart: EOF"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(subTest *testing.T) {
			//Load POST body
			req := createRequest(test.file)

			// Invoke callback handler
			email, err := Parse(req)
			if test.expectedError != nil {
				assert.Error(subTest, err, "expected an error to occur")
				return
			}

			assert.NoError(subTest, err, "did NOT expect an error to occur")

			from := "Example User <test@example.com>"
			assert.Equalf(subTest, email.Headers["From"], from, "Expected From: %s, Got: %s", from, email.Headers["From"])
		})
	}
}

func ExampleParsedEmail_parseHeaders() {
	headers := `
Foo: foo
Bar: baz
`
	email := ParsedEmail{
		Headers:     make(map[string]string),
		Body:        make(map[string]string),
		Attachments: make(map[string][]byte),
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
		Attachments: make(map[string][]byte),
		rawRequest:  nil,
	}

	if err := email.parseRawEmail(rawEmail); err != nil {
		log.Fatal(err)
	}

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

func TestValidate(t *testing.T) {
	tests := []struct {
		name          string
		values        map[string][]string
		expectedError error
	}{
		{
			name:          "MissingHeaders",
			values:        map[string][]string{},
			expectedError: fmt.Errorf("missing DKIM and SPF score"),
		},
		{
			name:          "FailedDkim",
			values:        map[string][]string{"dkim": {"pass", "fail", "pass"}, "SPF": {"pass"}},
			expectedError: fmt.Errorf("DKIM validation failed"),
		},
		{
			name:          "FailedSpf",
			values:        map[string][]string{"dkim": {"pass", "pass", "pass"}, "SPF": {"pass", "fail", "pass"}},
			expectedError: fmt.Errorf("SPF validation failed"),
		},
		{
			name:          "FailedSpfandDkim",
			values:        map[string][]string{"dkim": {"pass", "pass", "fail"}, "SPF": {"pass", "fail", "pass"}},
			expectedError: fmt.Errorf("DKIM validation failed"),
		},
		{
			name:   "success",
			values: map[string][]string{"dkim": {"pass", "pass", "pass"}, "SPF": {"pass", "pass", "pass"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(subTest *testing.T) {
			//Load POST body
			email := ParsedEmail{rawValues: test.values}
			err := email.Validate()

			if test.expectedError != nil {
				assert.EqualError(subTest, test.expectedError, err.Error())
				return
			}

			assert.NoError(subTest, err, "did NOT expect an error to occur")
		})
	}
}
