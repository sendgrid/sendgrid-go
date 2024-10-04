package client_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

func NewRequestHandler(accountSid string, authToken string) *client.RequestHandler {
	c := NewClient(accountSid, authToken)
	return client.NewRequestHandler(c)
}

func TestRequestHandler_BuildUrlSetRegion(t *testing.T) {
	// Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.region.twilio.com"))

	// Region set via requestHandler
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.twilio.com"))
	assert.Equal(t, "https://api.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeDefaultRegion(t *testing.T) {
	// Edge set via client
	requestHandler := NewRequestHandler("user", "pass")
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.us1.twilio.com", assertAndGetURL(t, requestHandler, "https://api.twilio.com"))
}

func TestRequestHandler_BuildUrlSetEdgeRegion(t *testing.T) {
	//Edge and Region set via url
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://api.edge.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.edge.region.twilio.com"))

	// Edge and Region set via client
	requestHandler.Edge = "edge"
	assert.Equal(t, "https://api.edge.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.region.twilio.com"))
	requestHandler.Region = "region"
	assert.Equal(t, "https://api.edge.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.twilio.com"))
	assert.Equal(t, "https://api.edge.region.twilio.com", assertAndGetURL(t, requestHandler, "https://api.urlEdge.urlRegion.twilio.com"))
}

func TestRequestHandler_BuildHostRawHostWithoutPeriods(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	assert.Equal(t, "https://prism_twilio:4010", assertAndGetURL(t, requestHandler, "https://prism_twilio:4010"))
}

func TestRequestHandler_BuildUrlInvalidCTLCharacter(t *testing.T) {
	requestHandler := NewRequestHandler("user", "pass")
	rawURL := "https://api.twilio.com/ServiceId\n"
	parsedURL, err := requestHandler.BuildUrl(rawURL)

	expectedErr := url.Error{Op: "parse", URL: rawURL, Err: errors.New("net/url: invalid control character in URL")}
	assert.Equal(t, &expectedErr, err)
	assert.Equal(t, parsedURL, "")
}

func assertAndGetURL(t *testing.T, requestHandler *client.RequestHandler, rawURL string) string {
	parsedURL, err := requestHandler.BuildUrl(rawURL)
	assert.Nil(t, err)
	return parsedURL
}

func TestRequestHandler_SendGetRequest(t *testing.T) {
	errorResponse := `{
	"status": 400,
	"code":20001,
	"message":"Bad request",
	"more_info":"https://www.twilio.com/docs/errors/20001"
}`
	errorServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write([]byte(errorResponse))
		}))
	defer errorServer.Close()

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Get(errorServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*client.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Nil(t, twilioError.Details)
}

func TestRequestHandler_SendPostRequest(t *testing.T) {
	errorResponse := `{
	"status": 400,
	"code":20001,
	"message":"Bad request",
	"more_info":"https://www.twilio.com/docs/errors/20001"
}`
	errorServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write([]byte(errorResponse))
		}))
	defer errorServer.Close()

	requestHandler := NewRequestHandler("user", "pass")
	resp, err := requestHandler.Post(errorServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*client.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Nil(t, twilioError.Details)
}
