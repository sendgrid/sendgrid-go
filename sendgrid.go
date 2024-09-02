package sendgrid

import (
	"errors"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/client"
	Alert "github.com/sendgrid/sendgrid-go/rest/api/v3/alerts"
	Student "github.com/sendgrid/sendgrid-go/rest/api/v3/student"
	"net/url"
	"os"
)

type RestClient struct {
	*client.RequestHandler
	Student *Student.ApiService
	Alert   *Alert.ApiService
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

// sendGridOptions for CreateRequest
type sendGridOptions struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

// sendgrid host map for different regions
var allowedRegionsHostMap = map[string]string{
	"eu":     "https://api.eu.sendgrid.com",
	"global": "https://api.sendgrid.com",
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

type ClientParams struct {
	ApiKey string
	Client client.BaseClient
}

// NewRestClientWithParams provides an initialized Twilio RestClient with params.
func NewRestClientWithParams(params ClientParams) *RestClient {
	requestHandler := client.NewRequestHandler(params.Client)

	if params.Client == nil {
		apiKey := params.ApiKey
		if apiKey == "" {
			apiKey = os.Getenv("SENDGRID_API_KEY")
		}

		defaultClient := &client.Client{
			Credentials: client.NewCredentials(apiKey),
		}

		requestHandler = client.NewRequestHandler(defaultClient)
	}

	c := &RestClient{
		RequestHandler: requestHandler,
	}

	c.Student = Student.NewApiService(c.RequestHandler)
	c.Alert = Alert.NewApiService(c.RequestHandler)
	return c
}

// NewRestClient provides an initialized Twilio RestClient.
func NewRestClient() *RestClient {
	return NewRestClientWithParams(ClientParams{})
}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(key string) *Client {
	request := GetRequest(key, "/v3/mail/send", "")
	request.Method = "POST"
	return &Client{request}
}

// extractEndpoint extracts the endpoint from a baseURL
func extractEndpoint(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	return parsedURL.Path, nil
}

// SetDataResidency modifies the host as per the region
/*
 * This allows support for global and eu regions only. This set will likely expand in the future.
 * Global should be the default
 * Global region means the message should be sent through:
 * HTTP: api.sendgrid.com
 * EU region means the message should be sent through:
 * HTTP: api.eu.sendgrid.com
 */
// @return [Request] the modified request object
func SetDataResidency(request rest.Request, region string) (rest.Request, error) {
	regionalHost, present := allowedRegionsHostMap[region]
	if !present {
		return request, errors.New("error: region can only be \"eu\" or \"global\"")
	}
	endpoint, err := extractEndpoint(request.BaseURL)
	if err != nil {
		return request, err
	}
	request.BaseURL = regionalHost + endpoint
	return request, nil
}
