// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import "github.com/sendgrid/rest"

const Version = "3.0.0"

// GetRequest returns a default request object.
func GetRequest(key string, endpoint string, host string, version string) rest.Request {
	if host == "" {
		host = "https://api.sendgrid.com"
	}
	if version == "" {
		version = "v3"
	}
	baseURL := host + "/" + version + endpoint
	requestHeaders := make(map[string]string)
	requestHeaders["Content-Type"] = "application/json"
	requestHeaders["Authorization"] = "Bearer " + key
	requestHeaders["User-Agent"] = "sendgrid/" + Version + ";go"
	request := rest.Request{
		BaseURL:        baseURL,
		RequestHeaders: requestHeaders,
	}
	return request
}

// API sets up the request to the SendGrid API, this is main interface.
func API(request rest.Request) (*rest.Response, error) {
	response, err := rest.API(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
