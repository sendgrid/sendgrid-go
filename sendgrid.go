// Package sendgrid provides a simple interface to interact with the Twilio SendGrid API
package sendgrid

import (
	"errors"
	"net/http"
	"strconv"
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

// Client is the Twilio SendGrid Go client
type Client struct {
	// rest.Request
	rest.Request
}

// options for requestNew
type options struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

func (o *options) baseURL() string {
	return o.Host + o.Endpoint
}

// GetRequest
// @return [Request] a default request object
func GetRequest(key, endpoint, host string) rest.Request {
	return requestNew(options{key, endpoint, host, ""})
}

// GetRequestSubuser like GetRequest but with On-Behalf of Subuser
// @return [Request] a default request object
func GetRequestSubuser(key, endpoint, host, subuser string) rest.Request {
	return requestNew(options{key, endpoint, host, subuser})
}

// requestNew create Request
// @return [Request] a default request object
func requestNew(options options) rest.Request {
	if options.Host == "" {
		options.Host = "https://api.sendgrid.com"
	}

	requestHeaders := map[string]string{
		"Authorization": "Bearer " + options.Key,
		"User-Agent":    "sendgrid/" + Version + ";go",
		"Accept":        "application/json",
	}

	if len(options.Subuser) != 0 {
		requestHeaders["On-Behalf-Of"] = options.Subuser
	}

	return rest.Request{
		BaseURL: options.baseURL(),
		Headers: requestHeaders,
	}
}

// Send sends an email through Twilio SendGrid
func (cl *Client) Send(email *mail.SGMailV3) (*rest.Response, error) {
	cl.Body = mail.GetRequestBody(email)
	return MakeRequest(cl.Request)
}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(key string) *Client {
	request := GetRequest(key, "/v3/mail/send", "")
	request.Method = "POST"
	return &Client{request}
}

// GetRequestSubuser like NewSendClient but with On-Behalf of Subuser
// @return [Client]
func NewSendClientSubuser(key, subuser string) *Client {
	request := GetRequestSubuser(key, "/v3/mail/send", "", subuser)
	request.Method = "POST"
	return &Client{request}
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = rest.DefaultClient

// API sets up the request to the Twilio SendGrid API, this is main interface.
// Please use the MakeRequest or MakeRequestAsync functions instead.
// (deprecated)
func API(request rest.Request) (*rest.Response, error) {
	return MakeRequest(request)
}

// MakeRequest attempts a Twilio SendGrid request synchronously.
func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.Send(request)
}

// MakeRequestRetry a synchronous request, but retry in the event of a rate
// limited response.
func MakeRequestRetry(request rest.Request) (*rest.Response, error) {
	retry := 0
	var response *rest.Response
	var err error

	for {
		response, err = MakeRequest(request)
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusTooManyRequests {
			return response, nil
		}

		if retry > rateLimitRetry {
			return nil, errors.New("Rate limit retry exceeded")
		}
		retry++

		resetTime := time.Now().Add(rateLimitSleep * time.Millisecond)

		reset, ok := response.Headers["X-RateLimit-Reset"]
		if ok && len(reset) > 0 {
			t, err := strconv.Atoi(reset[0])
			if err == nil {
				resetTime = time.Unix(int64(t), 0)
			}
		}
		time.Sleep(resetTime.Sub(time.Now()))
	}
}

// MakeRequestAsync attempts a request asynchronously in a new go
// routine. This function returns two channels: responses
// and errors. This function will retry in the case of a
// rate limit.
func MakeRequestAsync(request rest.Request) (chan *rest.Response, chan error) {
	r := make(chan *rest.Response)
	e := make(chan error)

	go func() {
		response, err := MakeRequestRetry(request)
		if err != nil {
			e <- err
		}
		if response != nil {
			r <- response
		}
	}()

	return r, e
}
