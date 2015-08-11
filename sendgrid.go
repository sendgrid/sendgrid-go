// Package sendgrid provides a simple interface to interact with the SendGrid API
package sendgrid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	Version  = "3.0.0"
	endpoint = "https://api.sendgrid.com/api/mail.send.json?"
)

// SGClient will contain the credentials and default values
type Client struct {
	apiUser  string
	apiPwd   string
	endpoint string
	Client   *http.Client
}

// NewClient will return a new Client. Used for username and password
func NewClient(apiUser, apiKey string) *Client {

	Client := &Client{
		apiUser:  apiUser,
		apiPwd:   apiKey,
		endpoint: endpoint,
	}

	return Client
}

// NewSendGridClient will return a new Client. Used for api key
func NewClientWithKey(apiKey string) *Client {
	Client := &Client{
		apiPwd:   apiKey,
		endpoint: endpoint,
	}

	return Client
}

// Send will send mail using SG web API
func (c *Client) Send(m *Mail) error {
	if c.Client == nil {
		c.Client = &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Second,
		}
	}
	var e error
	values, e := m.values()
	if e != nil {
		return e
	}
	if c.apiUser != "" {
		values.Set("api_user", c.apiUser)
		values.Set("api_key", c.apiPwd)
	}
	req, e := http.NewRequest("POST", c.endpoint, strings.NewReader(values.Encode()))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "sendgrid/"+Version+";go")

	// Using API key
	if c.apiUser == "" {
		req.Header.Set("Authorization", "Bearer "+c.apiPwd)
	}

	res, e := c.Client.Do(req)
	if e != nil {
		return fmt.Errorf("sendgrid.go: error:%v; response:%v", e, res)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	}

	body, _ := ioutil.ReadAll(res.Body)

	return fmt.Errorf("sendgrid.go: code:%d error:%v body:%s", res.StatusCode, e, body)
}
