// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import (
	"errors"
	"net/http"
	"time"

	"github.com/sendgrid/rest" // depends on version 2.2.0
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Version is this client library's current version
const (
	Version        = "3.1.0"
	rateLimitRetry = 5
	rateLimitSleep = 1100
)

// Client is the SendGrid Go client
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

// Send sends an email through SendGrid
func (cl *Client) Send(email *mail.SGMailV3) (*rest.Response, error) {
	cl.Body = mail.GetRequestBody(email)
	return API(cl.Request)
}

// NewSendClient constructs a new SendGrid client given an API key
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

// Request attempts a request asynchronously in a new go
// routine. This function returns two channels: responses
// and errors. This function will retry in the case of a
// rate limit.
func Request(request rest.Request) (chan *rest.Response, chan error) {
	r := make(chan *rest.Response)
	e := make(chan error)

	go func() {
		retry := 0
		for {
			response, err := DefaultClient.API(request)
			if err != nil {
				e <- err
				return
			}

			if response.StatusCode == http.StatusTooManyRequests {
				if retry > rateLimitRetry {
					e <- errors.New("Rate limit retry exceeded")
					return
				}
				retry++
				time.Sleep(rateLimitSleep * time.Millisecond)
				continue
			}

			r <- response
		}
	}()

	return r, e
}
