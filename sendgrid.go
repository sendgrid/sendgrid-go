// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import (
	"errors"
        "regexp"
	"github.com/sendgrid/rest" // depends on version 2.2.0
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Version is this client library's current version
const Version = "3.1.0"

type Client struct {
	rest.Request
}

// GetRequest returns a default request object.
func GetRequest(key string, endpoint string, host string) rest.Request {
	if host == "" {
		host = "https://api.sendgrid.com"
	}
	baseURL := host + endpoint
	requestHeaders := make(map[string]string)
	requestHeaders["Authorization"] = "Bearer " + key
	requestHeaders["User-Agent"] = "sendgrid/" + Version + ";go"
	requestHeaders["Accept"] = "application/json"
	request := rest.Request{
		BaseURL: baseURL,
		Headers: requestHeaders,
	}
	return request
}

func (cl *Client) Send(email *mail.SGMailV3) (*rest.Response, error) {
	cl.Body = mail.GetRequestBody(email)
	return API(cl.Request)
}

func NewSendClient(key string) *Client {
	request := GetRequest(key, "/v3/mail/send", "")
	request.Method = "POST"
	return &Client{request}
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = rest.DefaultClient

// API sets up the request to the SendGrid API, this is main interface.
func API(request rest.Request) (*rest.Response, error) {
	return DefaultClient.API(request)
}

// Checks email content for presence of an API key before sending.
func CheckSecrets(mailContent string) (int, error) {
	var secret = regexp.MustCompile(`SG.[a-zA-Z0-9_-]+.[a-zA-Z0-9_-]+`)
	if secret.MatchString(mailContent) {
		return -1, errors.New("Error: API key found in email content. Do not send secrets! ")
	}
	return 0, nil
}
