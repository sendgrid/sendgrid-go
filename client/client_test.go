package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	twilio "github.com/twilio/twilio-go/client"
)

var mockServer *httptest.Server
var testClient *twilio.Client

func NewClient(accountSid string, authToken string) *twilio.Client {
	c := &twilio.Client{
		Credentials: twilio.NewCredentials(accountSid, authToken),
		HTTPClient:  http.DefaultClient,
	}

	return c
}

func TestMain(m *testing.M) {
	mockServer = httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			encoder := json.NewEncoder(writer)
			_ = encoder.Encode(&d)
		}))
	defer mockServer.Close()

	testClient = NewClient("user", "pass")
	os.Exit(m.Run())
}

func TestClient_SendRequestError(t *testing.T) {
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

	resp, err := testClient.SendRequest("GET", errorServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Nil(t, twilioError.Details)
}

func TestClient_SendRequestDecodeError(t *testing.T) {
	errorResponse := `{
	"status": 400,
	"code":20001,
	"message":"Bad request",
	"more_info":"https://www.twilio.com/docs/errors/20001",
}`
	errorServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write([]byte(errorResponse))
		}))
	defer errorServer.Close()

	resp, err := testClient.SendRequest("GET", errorServer.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error decoding the response for an HTTP error code: 400")
	assert.Nil(t, resp)
}

func TestClient_SendRequestErrorWithDetails(t *testing.T) {
	errorResponse := []byte(`{
	"status": 400,
	"message": "Bad request",
	"code": 20001,
	"more_info": "https://www.twilio.com/docs/errors/20001",
	"details": {
		"foo": "bar"
	}
}`)
	errorServer := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			resp.WriteHeader(400)
			_, _ = resp.Write(errorResponse)
		}))
	defer errorServer.Close()

	resp, err := testClient.SendRequest("GET", errorServer.URL, nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
	details := make(map[string]interface{})
	details["foo"] = "bar"
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 20001, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/20001", twilioError.MoreInfo)
	assert.Equal(t, "Bad request", twilioError.Message)
	assert.Equal(t, details, twilioError.Details)
}

func TestClient_SendRequestUsernameError(t *testing.T) {
	newTestClient := NewClient("user1\nuser2", "pass")
	resp, err := newTestClient.SendRequest("GET", "http://example.org", nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 21222, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/21222", twilioError.MoreInfo)
	assert.Equal(t, "Invalid Username. Illegal chars", twilioError.Message)
}

func TestClient_SendRequestPasswordError(t *testing.T) {
	newTestClient := NewClient("user1", "pass1\npass2")
	resp, err := newTestClient.SendRequest("GET", "http://example.org", nil, nil) //nolint:bodyclose
	twilioError := err.(*twilio.TwilioRestError)
	assert.Nil(t, resp)
	assert.Equal(t, 400, twilioError.Status)
	assert.Equal(t, 21224, twilioError.Code)
	assert.Equal(t, "https://www.twilio.com/docs/errors/21224", twilioError.MoreInfo)
	assert.Equal(t, "Invalid Password. Illegal chars", twilioError.Message)
}

func TestClient_SendRequestWithRedirect(t *testing.T) {
	redirectServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(307)
			_, _ = writer.Write([]byte(`{"redirect_to": "some_place"}`))
		}))
	defer redirectServer.Close()

	resp, _ := testClient.SendRequest("GET", redirectServer.URL, nil, nil) //nolint:bodyclose
	assert.Equal(t, 307, resp.StatusCode)
}

func TestClient_SendRequestCreatesClient(t *testing.T) {
	c := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}
	resp, err := c.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_SendRequestWithData(t *testing.T) {
	dataServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			_ = request.ParseForm()
			assert.Equal(t, "bar", request.FormValue("foo"))
			d := map[string]interface{}{
				"response": "ok",
			}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer dataServer.Close()

	tests := []string{http.MethodGet, http.MethodPost}
	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			data := url.Values{}
			data.Set("foo", "bar")
			headers := map[string]interface{}{
				"Content-Type": "application/x-www-form-urlencoded",
			}
			resp, err := testClient.SendRequest(tc, dataServer.URL, data, headers) //nolint:bodyclose
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)
		})
	}
}

func TestClient_SendRequestWithHeaders(t *testing.T) {
	headerServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "bar", request.Header.Get("foo"))
			d := map[string]interface{}{
				"response": "ok",
			}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer headerServer.Close()

	headers := map[string]interface{}{
		"foo": "bar",
	}
	resp, err := testClient.SendRequest("GET", headerServer.URL, nil, headers) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_SetTimeoutTimesOut(t *testing.T) {
	timeoutServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			time.Sleep(100 * time.Microsecond)
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
			writer.WriteHeader(http.StatusOK)
		}))
	defer timeoutServer.Close()

	c := NewClient("user", "pass")
	c.SetTimeout(10 * time.Microsecond)
	_, err := c.SendRequest("GET", timeoutServer.URL, nil, nil) //nolint:bodyclose
	assert.Error(t, err)
}

func TestClient_SetTimeoutSucceeds(t *testing.T) {
	timeoutServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"response": "ok",
			}
			time.Sleep(100 * time.Microsecond)
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer timeoutServer.Close()

	c := NewClient("user", "pass")
	c.SetTimeout(10 * time.Second)
	resp, err := c.SendRequest("GET", timeoutServer.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_SetTimeoutCreatesClient(t *testing.T) {
	c := &twilio.Client{
		Credentials: twilio.NewCredentials("user", "pass"),
	}
	c.SetTimeout(20 * time.Second)
	resp, err := c.SendRequest("GET", mockServer.URL, nil, nil) //nolint:bodyclose
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_UnicodeResponse(t *testing.T) {
	unicodeServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			d := map[string]interface{}{
				"testing-unicode": "Ω≈ç√, 💩",
			}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(&d)
			if err != nil {
				t.Error(err)
			}
		}))
	defer unicodeServer.Close()

	c := NewClient("user", "pass")
	resp, _ := c.SendRequest("GET", unicodeServer.URL, nil, nil) //nolint:bodyclose
	assert.Equal(t, 200, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "{\"testing-unicode\":\"Ω≈ç√, 💩\"}\n", string(body))
}

func TestClient_SetAccountSid(t *testing.T) {
	client := NewClient("user", "pass")
	client.SetAccountSid("account_sid")
	assert.Equal(t, "account_sid", client.AccountSid())
}

func TestClient_DefaultUserAgentHeaders(t *testing.T) {
	headerServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			assert.Regexp(t, regexp.MustCompile(`^twilio-go/[0-9.]+(-rc.[0-9])*\s\(\w+\s\w+\)\sgo/\S+$`), request.Header.Get("User-Agent"))
		}))

	resp, _ := testClient.SendRequest("GET", headerServer.URL, nil, nil)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestClient_UserAgentExtensionsHeaders(t *testing.T) {
	var expectedExtensions = []string{"twilio-run/2.0.0-test", "flex-plugin/3.4.0"}
	testClient.UserAgentExtensions = expectedExtensions
	headerServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			var headersList = strings.Split(request.Header.Get("User-Agent"), " ")
			assert.Equal(t, headersList[len(headersList)-len(expectedExtensions):], expectedExtensions)
		}))
	resp, _ := testClient.SendRequest("GET", headerServer.URL, nil, nil)
	assert.Equal(t, 200, resp.StatusCode)
}
