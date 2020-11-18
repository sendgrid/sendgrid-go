package sendgrid

import (
	"github.com/sendgrid/rest"
)

// sendGridOptions for CreateRequest
type sendGridOptions struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

// GetRequest
// @return [Request] a default request object
func GetRequest(key, endpoint, host string) rest.Request {
	return createSendGridRequest(sendGridOptions{key, endpoint, host, ""})
}

// GetRequestSubuser like GetRequest but with On-Behalf of Subuser
// @return [Request] a default request object
func GetRequestSubuser(key, endpoint, host, subuser string) rest.Request {
	return createSendGridRequest(sendGridOptions{key, endpoint, host, subuser})
}

// createSendGridRequest create Request
// @return [Request] a default request object
func createSendGridRequest(sgOptions sendGridOptions) rest.Request {
	options := options{
		"Bearer " + sgOptions.Key,
		sgOptions.Endpoint,
		sgOptions.Host,
		sgOptions.Subuser,
	}

	if options.Host == "" {
		options.Host = "https://api.sendgrid.com"
	}

	return requestNew(options)
}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(key string) *Client {
	request := GetRequest(key, "/v3/mail/send", "")
	request.Method = "POST"
	return &Client{request}
}

// NewClientForEndpoint returns a client that can send requests to a specific endpoint.
func NewClientForEndpoint(key, endpoint string) *Client {
	request := GetRequest(key, endpoint, "")
	return &Client{request}
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = rest.DefaultClient

// API sets up the request to the SendGrid API, this is main interface.
// This function is deprecated. Please use the MakeRequest or
// MakeRequestAsync functions.
func API(request rest.Request) (*rest.Response, error) {
	return DefaultClient.API(request)
}

// MakeRequest attempts a SendGrid request synchronously.
func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.API(request)
}

// MakeRequestRetry a synchronous request, but retry in the event of a rate
// limited response.
func MakeRequestRetry(request rest.Request) (*rest.Response, error) {
	retry := 0
	var response *rest.Response
	var err error

	for {
		response, err = DefaultClient.API(request)
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

// GetRequestSubuser like NewSendClient but with On-Behalf of Subuser
// @return [Client]
func NewSendClientSubuser(key, subuser string) *Client {
	request := GetRequestSubuser(key, "/v3/mail/send", "", subuser)
	request.Method = "POST"
	return &Client{request}
}
