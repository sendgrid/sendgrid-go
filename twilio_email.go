package sendgrid

import (
	"encoding/base64"

	"github.com/sendgrid/rest"
)

// TwilioEmailOptions for GetTwilioEmailRequest
type TwilioEmailOptions struct {
	Username string
	Password string
	Endpoint string
	Host     string
}

// NewTwilioEmailSendClient constructs a new Twilio Email client given a username and password
func NewTwilioEmailSendClient(username, password string) *Client {
	return &Client{emailOptions: TwilioEmailOptions{Username: username, Password: password}, apiKey: ""}
}

// GetTwilioEmailRequest create Request
// @return [Request] a default request object
func GetTwilioEmailRequest(twilioEmailOptions TwilioEmailOptions) rest.Request {
	credentials := twilioEmailOptions.Username + ":" + twilioEmailOptions.Password
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(credentials))

	options := options{
		Auth:     "Basic " + encodedCreds,
		Endpoint: "/v3/mail/send",
		Host:     twilioEmailOptions.Host,
	}

	if options.Host == "" {
		options.Host = "https://email.twilio.com"
	}

	return requestNew(options)
}
