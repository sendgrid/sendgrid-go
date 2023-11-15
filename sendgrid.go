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
	Region   string
}

// GetRequest
// @return [Request] a default request object
func GetRequest(key, endpoint, host string, regionOptional ...string) rest.Request {
	region := ""
	if len(regionOptional) > 0 {
		region = regionOptional[0]
	}
	return createSendGridRequest(sendGridOptions{key, endpoint, host, "", region})
}

// GetRequestSubuser like GetRequest but with On-Behalf of Subuser
// @return [Request] a default request object
func GetRequestSubuser(key, endpoint, host, subuser string, regionOptional ...string) rest.Request {
	region := ""
	if len(regionOptional) > 0 {
		region = regionOptional[0]
	}
	return createSendGridRequest(sendGridOptions{key, endpoint, host, subuser, region})
}

// createSendGridRequest create Request
// @return [Request] a default request object
func createSendGridRequest(sgOptions sendGridOptions) rest.Request {
	options := options{
		"Bearer " + sgOptions.Key,
		sgOptions.Endpoint,
		sgOptions.Host,
		sgOptions.Subuser,
		sgOptions.Region,
	}

	if options.Host == "" {
		options.Host = "https://api.sendgrid.com"
	}

	return requestNew(options)
}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(key string, regionOptional ...string) *Client {
	region := ""
	if len(regionOptional) > 0 {
		region = regionOptional[0]
	}
	request := GetRequest(key, "/v3/mail/send", "", region)
	request.Method = "POST"
	return &Client{request}
}
