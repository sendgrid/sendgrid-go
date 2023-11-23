package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/assert"
)

func TestLicenseYear(t *testing.T) {
	d, err := os.ReadFile("LICENSE")
	assert.Nil(t, err, "Cannot read the LICENSE file")
	l := fmt.Sprintf("Copyright (C) %v, Twilio SendGrid, Inc.", time.Now().Year())
	assert.True(t, strings.Contains(string(d), l), fmt.Sprintf("License date range is not correct, it should be: %v", l))
}

func TestRepoFiles(t *testing.T) {
	fs := []string{
		"Dockerfile",
		".env_sample",
		".gitignore",
		".github/workflows/test-and-deploy.yml",
		"CHANGELOG.md",
		"CODE_OF_CONDUCT.md",
		"CONTRIBUTING.md",
		"LICENSE",
		"PULL_REQUEST_TEMPLATE.md",
		"README.md",
		"TROUBLESHOOTING.md",
		"USAGE.md",
	}
	for _, f := range fs {
		_, err := os.Stat(f)
		assert.False(t, os.IsNotExist(err), fmt.Sprintf("Repo file does not exist: %[1]v", f))
	}
}

func TestGetRequest(t *testing.T) {
	request := GetRequest("", "", "")
	assert.Equal(t, "https://api.sendgrid.com", request.BaseURL, "Host default not set")
	assert.Equal(t, "Bearer ", request.Headers["Authorization"], "Wrong default Authorization")
	assert.Equal(t, "sendgrid/"+Version+";go", request.Headers["User-Agent"], "Wrong default User Agent")
	request = GetRequest("API_KEY", "/v3/endpoint", "https://test.api.com")
	assert.Equal(t, "Bearer API_KEY", request.Headers["Authorization"], "Wrong Authorization")
	assert.Equal(t, "sendgrid/"+Version+";go", request.Headers["User-Agent"], "Wrong User Agent")
	assert.Equal(t, "application/json", request.Headers["Accept"], "Wrong Accept Agent")
}

func ShouldHaveHeaders(request *rest.Request, t *testing.T) {
	if request.Headers["Authorization"] != "Bearer API_KEY" {
		t.Error("Wrong Authorization")
	}
	if request.Headers["User-Agent"] != "sendgrid/"+Version+";go" {
		t.Error("Wrong User Agent")
	}
	if request.Headers["Accept"] != "application/json" {
		t.Error("Wrong Accept header")
	}
	if request.Headers["On-Behalf-Of"] != "subuserUsername" {
		t.Error("Wrong On-Behalf-Of")
	}
}

func TestGetRequestSubuser(t *testing.T) {
	request := GetRequestSubuser("API_KEY", "/v3/endpoint", "https://test.api.com", "subuserUsername")

	if request.BaseURL != "https://test.api.com/v3/endpoint" {
		t.Error("Host not set correctly")
	}

	ShouldHaveHeaders(&request, t)
}

func TestSetDataResidencyEU(t *testing.T) {
	request := GetRequest("API_KEY", "", "")
	request, err := SetDataResidency(request, "eu")
	assert.Nil(t, err)
	assert.Equal(t, "https://api.eu.sendgrid.com", request.BaseURL, "Host not correct as per the region")
}

func TestSetDataResidencyGlobal(t *testing.T) {
	request := GetRequest("API_KEY", "", "https://api.sendgrid.com")
	request, err := SetDataResidency(request, "global")
	assert.Nil(t, err)
	assert.Equal(t, "https://api.sendgrid.com", request.BaseURL, "Host not correct as per the region")
}

func TestSetDataResidencyOverrideHost(t *testing.T) {
	request := GetRequest("API_KEY", "", "https://test.api.com")
	request, err := SetDataResidency(request, "eu")
	assert.Nil(t, err)
	assert.Equal(t, "https://api.eu.sendgrid.com", request.BaseURL, "Host not correct as per the region")
}

func TestSetDataResidencyIncorrectRegion(t *testing.T) {
	request := GetRequest("API_KEY", "", "")
	_, err := SetDataResidency(request, "foo")
	assert.NotNil(t, err, "error: region can only be \"eu\" or \"global\"")
}

func TestSetDataResidencyNullRegion(t *testing.T) {
	request := GetRequest("API_KEY", "", "")
	_, err := SetDataResidency(request, "")
	assert.NotNil(t, err, "error: region can only be \"eu\" or \"global\"")
}

func TestSetDataResidencyDefaultRegion(t *testing.T) {
	request := GetRequest("API_KEY", "", "")
	assert.Equal(t, "https://api.sendgrid.com", request.BaseURL, "Host not correct as per the region")
}

func getRequest(endpoint string) rest.Request {
	return GetRequest("SENDGRID_APIKEY", endpoint, "")
}

func TestCustomHTTPClient(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 20)
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	_, err := custom.Send(request)
	assert.NotNil(t, err, "A timeout did not trigger as expected")
	assert.True(t, strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers"), "We did not receive the Timeout error")
}

func TestRequestRetry_rateLimit(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", strconv.Itoa(int(time.Now().Add(1*time.Second).Unix())))
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	DefaultClient = &custom
	_, err := MakeRequestRetry(request)
	assert.NotNil(t, err, "An error did not trigger")
	assert.True(t, strings.Contains(err.Error(), "rate limit retry exceeded"), "We did not receive the rate limit error")
	DefaultClient = rest.DefaultClient
}

func TestRequestRetry_rateLimit_noHeader(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	DefaultClient = &custom
	_, err := MakeRequestRetry(request)
	assert.NotNil(t, err, "An error did not trigger")
	assert.True(t, strings.Contains(err.Error(), "rate limit retry exceeded"), "We did not receive the rate limit error")
	DefaultClient = rest.DefaultClient
}

func TestRequestRetryWithContext_cancel(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	DefaultClient = &custom
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()
	_, err := MakeRequestRetryWithContext(ctx, request)

	assert.NotNil(t, err, "An error did not trigger")
	assert.Equal(t, context.Canceled, err)
	DefaultClient = rest.DefaultClient
}

func TestRequestAsync(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	DefaultClient = &custom
	r, e := MakeRequestAsync(request)

	select {
	case <-r:
	case err := <-e:
		t.Errorf("Received an error,:%v", err)
	case <-time.After(10 * time.Second):
		t.Error("Timed out waiting for a response")
	}
	DefaultClient = rest.DefaultClient
}

func TestRequestAsync_rateLimit(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", strconv.Itoa(int(time.Now().Add(1*time.Second).Unix())))
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer fakeServer.Close()
	apiKey := "SENDGRID_APIKEY"
	host := fakeServer.URL
	request := GetRequest(apiKey, "/v3/test_endpoint", host)
	request.Method = "GET"
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	DefaultClient = &custom
	r, e := MakeRequestAsync(request)

	select {
	case <-r:
		t.Error("Received a valid response")
		return
	case err := <-e:
		assert.True(t, strings.Contains(err.Error(), "rate limit retry exceeded"), "We did not receive the rate limit error")
	case <-time.After(10 * time.Second):
		t.Error("Timed out waiting for an error")
	}
	DefaultClient = rest.DefaultClient
}

func Test_test_access_settings_activity_get(t *testing.T) {
	request := getRequest("/v3/access_settings/activity")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_access_settings_whitelist_post(t *testing.T) {
	request := getRequest("/v3/access_settings/whitelist")
	request.Method = "POST"
	request.Body = []byte(` {
  "ips": [
    {
      "ip": "192.168.1.1"
    },
    {
      "ip": "192.*.*.*"
    },
    {
      "ip": "192.168.1.3/32"
    }
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_access_settings_whitelist_get(t *testing.T) {
	request := getRequest("/v3/access_settings/whitelist")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_access_settings_whitelist_delete(t *testing.T) {
	request := getRequest("/v3/access_settings/whitelist")
	request.Method = "DELETE"
	request.Body = []byte(` {
  "ids": [
    1,
    2,
    3
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_access_settings_whitelist__rule_id__get(t *testing.T) {
	request := getRequest("/v3/access_settings/whitelist/{rule_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_access_settings_whitelist__rule_id__delete(t *testing.T) {
	request := getRequest("/v3/access_settings/whitelist/{rule_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_alerts_post(t *testing.T) {
	request := getRequest("/v3/alerts")
	request.Method = "POST"
	request.Body = []byte(` {
  "email_to": "example@example.com",
  "frequency": "daily",
  "type": "stats_notification"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_alerts_get(t *testing.T) {
	request := getRequest("/v3/alerts")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_alerts__alert_id__patch(t *testing.T) {
	request := getRequest("/v3/alerts/{alert_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email_to": "example@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_alerts__alert_id__get(t *testing.T) {
	request := getRequest("/v3/alerts/{alert_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_alerts__alert_id__delete(t *testing.T) {
	request := getRequest("/v3/alerts/{alert_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys_post(t *testing.T) {
	request := getRequest("/v3/api_keys")
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "My API Key",
  "sample": "data",
  "scopes": [
    "mail.send",
    "alerts.create",
    "alerts.read"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys_get(t *testing.T) {
	request := getRequest("/v3/api_keys")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys__api_key_id__put(t *testing.T) {
	request := getRequest("/v3/api_keys/{api_key_id}")
	request.Method = "PUT"
	request.Body = []byte(` {
  "name": "A New Hope",
  "scopes": [
    "user.profile.read",
    "user.profile.update"
  ]
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys__api_key_id__patch(t *testing.T) {
	request := getRequest("/v3/api_keys/{api_key_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "A New Hope"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys__api_key_id__get(t *testing.T) {
	request := getRequest("/v3/api_keys/{api_key_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_api_keys__api_key_id__delete(t *testing.T) {
	request := getRequest("/v3/api_keys/{api_key_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups_post(t *testing.T) {
	request := getRequest("/v3/asm/groups")
	request.Method = "POST"
	request.Body = []byte(` {
  "description": "Suggestions for products our users might like.",
  "is_default": true,
  "name": "Product Suggestions"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups_get(t *testing.T) {
	request := getRequest("/v3/asm/groups")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__patch(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "description": "Suggestions for items our users might like.",
  "id": 103,
  "name": "Item Suggestions"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__get(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__delete(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__suppressions_post(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}/suppressions")
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__suppressions_get(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}/suppressions")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__suppressions_search_post(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}/suppressions/search")
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "exists1@example.com",
    "exists2@example.com",
    "doesnotexists@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_groups__group_id__suppressions__email__delete(t *testing.T) {
	request := getRequest("/v3/asm/groups/{group_id}/suppressions/{email}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_suppressions_get(t *testing.T) {
	request := getRequest("/v3/asm/suppressions")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_suppressions_global_post(t *testing.T) {
	request := getRequest("/v3/asm/suppressions/global")
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com",
    "test2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_suppressions_global__email__get(t *testing.T) {
	request := getRequest("/v3/asm/suppressions/global/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_suppressions_global__email__delete(t *testing.T) {
	request := getRequest("/v3/asm/suppressions/global/{email}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_asm_suppressions__email__get(t *testing.T) {
	request := getRequest("/v3/asm/suppressions/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_browsers_stats_get(t *testing.T) {
	request := getRequest("/v3/browsers/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["browsers"] = "test_string"
	queryParams["limit"] = "test_string"
	queryParams["offset"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns_post(t *testing.T) {
	request := getRequest("/v3/campaigns")
	request.Method = "POST"
	request.Body = []byte(` {
  "categories": [
    "spring line"
  ],
  "custom_unsubscribe_url": "",
  "html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
  "ip_pool": "marketing",
  "list_ids": [
    110,
    124
  ],
  "plain_content": "Check out our spring line!",
  "segment_ids": [
    110
  ],
  "sender_id": 124451,
  "subject": "New Products for Spring!",
  "suppression_group_id": 42,
  "title": "March Newsletter"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns_get(t *testing.T) {
	request := getRequest("/v3/campaigns")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__patch(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "categories": [
    "summer line"
  ],
  "html_content": "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>",
  "plain_content": "Check out our summer line!",
  "subject": "New Products for Summer!",
  "title": "May Newsletter"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__get(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__delete(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_patch(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "send_at": 1489451436
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_post(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules")
	request.Method = "POST"
	request.Body = []byte(` {
  "send_at": 1489771528
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_get(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_delete(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_now_post(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules/now")
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_campaigns__campaign_id__schedules_test_post(t *testing.T) {
	request := getRequest("/v3/campaigns/{campaign_id}/schedules/test")
	request.Method = "POST"
	request.Body = []byte(` {
  "to": "your.email@example.com"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_categories_get(t *testing.T) {
	request := getRequest("/v3/categories")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["category"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_categories_stats_get(t *testing.T) {
	request := getRequest("/v3/categories/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["categories"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_categories_stats_sums_get(t *testing.T) {
	request := getRequest("/v3/categories/stats/sums")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["sort_by_direction"] = "asc"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_clients_stats_get(t *testing.T) {
	request := getRequest("/v3/clients/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_clients__client_type__stats_get(t *testing.T) {
	request := getRequest("/v3/clients/{client_type}/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_custom_fields_post(t *testing.T) {
	request := getRequest("/v3/contactdb/custom_fields")
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "pet",
  "type": "text"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_custom_fields_get(t *testing.T) {
	request := getRequest("/v3/contactdb/custom_fields")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_custom_fields__custom_field_id__get(t *testing.T) {
	request := getRequest("/v3/contactdb/custom_fields/{custom_field_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_custom_fields__custom_field_id__delete(t *testing.T) {
	request := getRequest("/v3/contactdb/custom_fields/{custom_field_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "202"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 202, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists_post(t *testing.T) {
	request := getRequest("/v3/contactdb/lists")
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "your list name"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists_get(t *testing.T) {
	request := getRequest("/v3/contactdb/lists")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists_delete(t *testing.T) {
	request := getRequest("/v3/contactdb/lists")
	request.Method = "DELETE"
	request.Body = []byte(` [
  1,
  2,
  3,
  4
]`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__patch(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "newlistname"
}`)
	queryParams := make(map[string]string)
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__get(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__delete(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}")
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "202"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 202, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__recipients_post(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}/recipients")
	request.Method = "POST"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__recipients_get(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}/recipients")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__recipients__recipient_id__post(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}/recipients/{recipient_id}")
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_lists__list_id__recipients__recipient_id__delete(t *testing.T) {
	request := getRequest("/v3/contactdb/lists/{list_id}/recipients/{recipient_id}")
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["recipient_id"] = "1"
	queryParams["list_id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_patch(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients")
	request.Method = "PATCH"
	request.Body = []byte(` [
  {
    "email": "jones@example.com",
    "first_name": "Guy",
    "last_name": "Jones"
  }
]`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_post(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients")
	request.Method = "POST"
	request.Body = []byte(` [
  {
    "age": 25,
    "email": "example@example.com",
    "first_name": "",
    "last_name": "User"
  },
  {
    "age": 25,
    "email": "example2@example.com",
    "first_name": "Example",
    "last_name": "User"
  }
]`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_delete(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients")
	request.Method = "DELETE"
	request.Body = []byte(` [
  "recipient_id1",
  "recipient_id2"
]`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_billable_count_get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/billable_count")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_count_get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/count")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients_search_get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/search")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["{field_name}"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients__recipient_id__get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/{recipient_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients__recipient_id__delete(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/{recipient_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_recipients__recipient_id__lists_get(t *testing.T) {
	request := getRequest("/v3/contactdb/recipients/{recipient_id}/lists")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_reserved_fields_get(t *testing.T) {
	request := getRequest("/v3/contactdb/reserved_fields")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments_post(t *testing.T) {
	request := getRequest("/v3/contactdb/segments")
	request.Method = "POST"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    },
    {
      "and_or": "and",
      "field": "last_clicked",
      "operator": "gt",
      "value": "01/02/2015"
    },
    {
      "and_or": "or",
      "field": "clicks.campaign_identifier",
      "operator": "eq",
      "value": "513"
    }
  ],
  "list_id": 4,
  "name": "Last Name Miller"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments_get(t *testing.T) {
	request := getRequest("/v3/contactdb/segments")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments__segment_id__patch(t *testing.T) {
	request := getRequest("/v3/contactdb/segments/{segment_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "conditions": [
    {
      "and_or": "",
      "field": "last_name",
      "operator": "eq",
      "value": "Miller"
    }
  ],
  "list_id": 5,
  "name": "The Millers"
}`)
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments__segment_id__get(t *testing.T) {
	request := getRequest("/v3/contactdb/segments/{segment_id}")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["segment_id"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments__segment_id__delete(t *testing.T) {
	request := getRequest("/v3/contactdb/segments/{segment_id}")
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["delete_contacts"] = "true"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_contactdb_segments__segment_id__recipients_get(t *testing.T) {
	request := getRequest("/v3/contactdb/segments/{segment_id}/recipients")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["page"] = "1"
	queryParams["page_size"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_devices_stats_get(t *testing.T) {
	request := getRequest("/v3/devices/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_geo_stats_get(t *testing.T) {
	request := getRequest("/v3/geo/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["country"] = "US"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_get(t *testing.T) {
	request := getRequest("/v3/ips")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["subuser"] = "test_string"
	queryParams["ip"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["exclude_whitelabels"] = "true"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_assigned_get(t *testing.T) {
	request := getRequest("/v3/ips/assigned")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools_post(t *testing.T) {
	request := getRequest("/v3/ips/pools")
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "marketing"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools_get(t *testing.T) {
	request := getRequest("/v3/ips/pools")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools__pool_name__put(t *testing.T) {
	request := getRequest("/v3/ips/pools/{pool_name}")
	request.Method = "PUT"
	request.Body = []byte(` {
  "name": "new_pool_name"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools__pool_name__get(t *testing.T) {
	request := getRequest("/v3/ips/pools/{pool_name}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools__pool_name__delete(t *testing.T) {
	request := getRequest("/v3/ips/pools/{pool_name}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools__pool_name__ips_post(t *testing.T) {
	request := getRequest("/v3/ips/pools/{pool_name}/ips")
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "0.0.0.0"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_pools__pool_name__ips__ip__delete(t *testing.T) {
	request := getRequest("/v3/ips/pools/{pool_name}/ips/{ip}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_warmup_post(t *testing.T) {
	request := getRequest("/v3/ips/warmup")
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "0.0.0.0"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_warmup_get(t *testing.T) {
	request := getRequest("/v3/ips/warmup")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_warmup__ip_address__get(t *testing.T) {
	request := getRequest("/v3/ips/warmup/{ip_address}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips_warmup__ip_address__delete(t *testing.T) {
	request := getRequest("/v3/ips/warmup/{ip_address}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_ips__ip_address__get(t *testing.T) {
	request := getRequest("/v3/ips/{ip_address}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_batch_post(t *testing.T) {
	request := getRequest("/v3/mail/batch")
	request.Method = "POST"
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_batch__batch_id__get(t *testing.T) {
	request := getRequest("/v3/mail/batch/{batch_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_send_client_with_mail_body_compression_enabled(t *testing.T) {
	apiKey := "SENDGRID_API_KEY"
	client := NewSendClient(apiKey)
	client.Headers["Content-Encoding"] = "gzip"

	emailBytes := []byte(` {
		"asm": {
			"group_id": 1,
			"groups_to_display": [
			1,
			2,
			3
			]
		},
		"attachments": [
			{
			"content": "[BASE64 encoded content block here]",
			"content_id": "ii_139db99fdb5c3704",
			"disposition": "inline",
			"filename": "file1.jpg",
			"name": "file1",
			"type": "jpg"
			}
		],
		"batch_id": "[YOUR BATCH ID GOES HERE]",
		"categories": [
			"category1",
			"category2"
		],
		"content": [
			{
			"type": "text/html",
			"value": "<html><p>Hello, world!</p><img src=[CID GOES HERE]></img></html>"
			}
		],
		"custom_args": {
			"New Argument 1": "New Value 1",
			"activationAttempt": "1",
			"customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
		},
		"from": {
			"email": "sam.smith@example.com",
			"name": "Sam Smith"
		},
		"headers": {},
		"ip_pool_name": "[YOUR POOL NAME GOES HERE]",
		"mail_settings": {
			"bcc": {
			"email": "ben.doe@example.com",
			"enable": true
			},
			"bypass_list_management": {
			"enable": true
			},
			"footer": {
			"enable": true,
			"html": "<p>Thanks</br>The Twilio SendGrid Team</p>",
			"text": "Thanks,/n The Twilio SendGrid Team"
			},
			"sandbox_mode": {
			"enable": false
			},
			"spam_check": {
			"enable": true,
			"post_to_url": "http://example.com/compliance",
			"threshold": 3
			}
		},
		"personalizations": [
			{
			"bcc": [
				{
				"email": "sam.doe@example.com",
				"name": "Sam Doe"
				}
			],
			"cc": [
				{
				"email": "jane.doe@example.com",
				"name": "Jane Doe"
				}
			],
			"custom_args": {
				"New Argument 1": "New Value 1",
				"activationAttempt": "1",
				"customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
			},
			"headers": {
				"X-Accept-Language": "en",
				"X-Mailer": "MyApp"
			},
			"send_at": 1409348513,
			"subject": "Hello, World!",
			"substitutions": {
				"id": "substitutions",
				"type": "object"
			},
			"to": [
				{
				"email": "john.doe@example.com",
				"name": "John Doe"
				}
			]
			}
		],
		"reply_to": {
			"email": "sam.smith@example.com",
			"name": "Sam Smith"
		},
		"send_at": 1409348513,
		"subject": "Hello, World!",
		"template_id": "[YOUR TEMPLATE ID GOES HERE]",
		"tracking_settings": {
			"click_tracking": {
			"enable": true,
			"enable_text": true
			},
			"ganalytics": {
			"enable": true,
			"utm_campaign": "[NAME OF YOUR REFERRER SOURCE]",
			"utm_content": "[USE THIS SPACE TO DIFFERENTIATE YOUR EMAIL FROM ADS]",
			"utm_medium": "[NAME OF YOUR MARKETING MEDIUM e.g. email]",
			"utm_name": "[NAME OF YOUR CAMPAIGN]",
			"utm_term": "[IDENTIFY PAID KEYWORDS HERE]"
			},
			"open_tracking": {
			"enable": true,
			"substitution_tag": "%opentrack"
			},
			"subscription_tracking": {
			"enable": true,
			"html": "If you would like to unsubscribe and stop receiving these emails <% clickhere %>.",
			"substitution_tag": "<%click here%>",
			"text": "If you would like to unsubscribe and stop receiving these emails <% click here %>."
			}
		}
	}`)
	email := &mail.SGMailV3{}
	err := json.Unmarshal(emailBytes, email)
	assert.Nil(t, err, fmt.Sprintf("Unmarshal error: %v", err))
	client.Request.Headers["X-Mock"] = "202"
	response, err := client.Send(email)
	if err != nil {
		t.Log(err)
	}
	t.Log(response)
	assert.Equal(t, 202, response.StatusCode, "Wrong status code returned")

}

func Test_test_send_client(t *testing.T) {
	apiKey := "SENDGRID_APIKEY"
	client := NewSendClient(apiKey)

	emailBytes := []byte(` {
		"asm": {
			"group_id": 1,
			"groups_to_display": [
			1,
			2,
			3
			]
		},
		"attachments": [
			{
			"content": "[BASE64 encoded content block here]",
			"content_id": "ii_139db99fdb5c3704",
			"disposition": "inline",
			"filename": "file1.jpg",
			"name": "file1",
			"type": "jpg"
			}
		],
		"batch_id": "[YOUR BATCH ID GOES HERE]",
		"categories": [
			"category1",
			"category2"
		],
		"content": [
			{
			"type": "text/html",
			"value": "<html><p>Hello, world!</p><img src=[CID GOES HERE]></img></html>"
			}
		],
		"custom_args": {
			"New Argument 1": "New Value 1",
			"activationAttempt": "1",
			"customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
		},
		"from": {
			"email": "sam.smith@example.com",
			"name": "Sam Smith"
		},
		"headers": {},
		"ip_pool_name": "[YOUR POOL NAME GOES HERE]",
		"mail_settings": {
			"bcc": {
			"email": "ben.doe@example.com",
			"enable": true
			},
			"bypass_list_management": {
			"enable": true
			},
			"footer": {
			"enable": true,
			"html": "<p>Thanks</br>The Twilio SendGrid Team</p>",
			"text": "Thanks,/n The Twilio SendGrid Team"
			},
			"sandbox_mode": {
			"enable": false
			},
			"spam_check": {
			"enable": true,
			"post_to_url": "http://example.com/compliance",
			"threshold": 3
			}
		},
		"personalizations": [
			{
			"bcc": [
				{
				"email": "sam.doe@example.com",
				"name": "Sam Doe"
				}
			],
			"cc": [
				{
				"email": "jane.doe@example.com",
				"name": "Jane Doe"
				}
			],
			"custom_args": {
				"New Argument 1": "New Value 1",
				"activationAttempt": "1",
				"customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
			},
			"headers": {
				"X-Accept-Language": "en",
				"X-Mailer": "MyApp"
			},
			"send_at": 1409348513,
			"subject": "Hello, World!",
			"substitutions": {
				"id": "substitutions",
				"type": "object"
			},
			"to": [
				{
				"email": "john.doe@example.com",
				"name": "John Doe"
				}
			]
			}
		],
		"reply_to": {
			"email": "sam.smith@example.com",
			"name": "Sam Smith"
		},
		"send_at": 1409348513,
		"subject": "Hello, World!",
		"template_id": "[YOUR TEMPLATE ID GOES HERE]",
		"tracking_settings": {
			"click_tracking": {
			"enable": true,
			"enable_text": true
			},
			"ganalytics": {
			"enable": true,
			"utm_campaign": "[NAME OF YOUR REFERRER SOURCE]",
			"utm_content": "[USE THIS SPACE TO DIFFERENTIATE YOUR EMAIL FROM ADS]",
			"utm_medium": "[NAME OF YOUR MARKETING MEDIUM e.g. email]",
			"utm_name": "[NAME OF YOUR CAMPAIGN]",
			"utm_term": "[IDENTIFY PAID KEYWORDS HERE]"
			},
			"open_tracking": {
			"enable": true,
			"substitution_tag": "%opentrack"
			},
			"subscription_tracking": {
			"enable": true,
			"html": "If you would like to unsubscribe and stop receiving these emails <% clickhere %>.",
			"substitution_tag": "<%click here%>",
			"text": "If you would like to unsubscribe and stop receiving these emails <% click here %>."
			}
		}
	}`)
	email := &mail.SGMailV3{}
	err := json.Unmarshal(emailBytes, email)
	assert.Nil(t, err, fmt.Sprintf("Unmarshal error: %v", err))
	client.Request.Headers["X-Mock"] = "202"
	response, err := client.Send(email)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 202, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_send_post(t *testing.T) {
	request := getRequest("/v3/mail/send")
	request.Method = "POST"
	request.Body = []byte(` {
  "asm": {
    "group_id": 1,
    "groups_to_display": [
      1,
      2,
      3
    ]
  },
  "attachments": [
    {
      "content": "[BASE64 encoded content block here]",
      "content_id": "ii_139db99fdb5c3704",
      "disposition": "inline",
      "filename": "file1.jpg",
      "name": "file1",
      "type": "jpg"
    }
  ],
  "batch_id": "[YOUR BATCH ID GOES HERE]",
  "categories": [
    "category1",
    "category2"
  ],
  "content": [
    {
      "type": "text/html",
      "value": "<html><p>Hello, world!</p><img src=[CID GOES HERE]></img></html>"
    }
  ],
  "custom_args": {
    "New Argument 1": "New Value 1",
    "activationAttempt": "1",
    "customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
  },
  "from": {
    "email": "sam.smith@example.com",
    "name": "Sam Smith"
  },
  "headers": {},
  "ip_pool_name": "[YOUR POOL NAME GOES HERE]",
  "mail_settings": {
    "bcc": {
      "email": "ben.doe@example.com",
      "enable": true
    },
    "bypass_list_management": {
      "enable": true
    },
    "footer": {
      "enable": true,
      "html": "<p>Thanks</br>The Twilio SendGrid Team</p>",
      "text": "Thanks,/n The Twilio SendGrid Team"
    },
    "sandbox_mode": {
      "enable": false
    },
    "spam_check": {
      "enable": true,
      "post_to_url": "http://example.com/compliance",
      "threshold": 3
    }
  },
  "personalizations": [
    {
      "bcc": [
        {
          "email": "sam.doe@example.com",
          "name": "Sam Doe"
        }
      ],
      "cc": [
        {
          "email": "jane.doe@example.com",
          "name": "Jane Doe"
        }
      ],
      "custom_args": {
        "New Argument 1": "New Value 1",
        "activationAttempt": "1",
        "customerAccountNumber": "[CUSTOMER ACCOUNT NUMBER GOES HERE]"
      },
      "headers": {
        "X-Accept-Language": "en",
        "X-Mailer": "MyApp"
      },
      "send_at": 1409348513,
      "subject": "Hello, World!",
      "substitutions": {
        "id": "substitutions",
        "type": "object"
      },
      "to": [
        {
          "email": "john.doe@example.com",
          "name": "John Doe"
        }
      ]
    }
  ],
  "reply_to": {
    "email": "sam.smith@example.com",
    "name": "Sam Smith"
  },
  "sections": {
    "section": {
      ":sectionName1": "section 1 text",
      ":sectionName2": "section 2 text"
    }
  },
  "send_at": 1409348513,
  "subject": "Hello, World!",
  "template_id": "[YOUR TEMPLATE ID GOES HERE]",
  "tracking_settings": {
    "click_tracking": {
      "enable": true,
      "enable_text": true
    },
    "ganalytics": {
      "enable": true,
      "utm_campaign": "[NAME OF YOUR REFERRER SOURCE]",
      "utm_content": "[USE THIS SPACE TO DIFFERENTIATE YOUR EMAIL FROM ADS]",
      "utm_medium": "[NAME OF YOUR MARKETING MEDIUM e.g. email]",
      "utm_name": "[NAME OF YOUR CAMPAIGN]",
      "utm_term": "[IDENTIFY PAID KEYWORDS HERE]"
    },
    "open_tracking": {
      "enable": true,
      "substitution_tag": "%opentrack"
    },
    "subscription_tracking": {
      "enable": true,
      "html": "If you would like to unsubscribe and stop receiving these emails <% clickhere %>.",
      "substitution_tag": "<%click here%>",
      "text": "If you would like to unsubscribe and stop receiving these emails <% click here %>."
    }
  }
}`)
	request.Headers["X-Mock"] = "202"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 202, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_get(t *testing.T) {
	request := getRequest("/v3/mail_settings")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_address_whitelist_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/address_whitelist")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "list": [
    "email1@example.com",
    "example.com"
  ]
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_address_whitelist_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/address_whitelist")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_bcc_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/bcc")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "email@example.com",
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_bcc_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/bcc")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_bounce_purge_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/bounce_purge")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "hard_bounces": 5,
  "soft_bounces": 5
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_bounce_purge_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/bounce_purge")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_footer_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/footer")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "...",
  "plain_content": "..."
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_footer_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/footer")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_forward_bounce_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/forward_bounce")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "example@example.com",
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_forward_bounce_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/forward_bounce")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_forward_spam_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/forward_spam")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "email": "",
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_forward_spam_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/forward_spam")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_plain_content_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/plain_content")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_plain_content_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/plain_content")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_spam_check_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/spam_check")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "max_score": 5,
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_spam_check_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/spam_check")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_template_patch(t *testing.T) {
	request := getRequest("/v3/mail_settings/template")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "<% body %>"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mail_settings_template_get(t *testing.T) {
	request := getRequest("/v3/mail_settings/template")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_mailbox_providers_stats_get(t *testing.T) {
	request := getRequest("/v3/mailbox_providers/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["mailbox_providers"] = "test_string"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_partner_settings_get(t *testing.T) {
	request := getRequest("/v3/partner_settings")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_partner_settings_new_relic_patch(t *testing.T) {
	request := getRequest("/v3/partner_settings/new_relic")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enable_subuser_statistics": true,
  "enabled": true,
  "license_key": ""
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_partner_settings_new_relic_get(t *testing.T) {
	request := getRequest("/v3/partner_settings/new_relic")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_scopes_get(t *testing.T) {
	request := getRequest("/v3/scopes")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders_post(t *testing.T) {
	request := getRequest("/v3/senders")
	request.Method = "POST"
	request.Body = []byte(` {
  "address": "123 Elm St.",
  "address_2": "Apt. 456",
  "city": "Denver",
  "country": "United States",
  "from": {
    "email": "from@example.com",
    "name": "Example INC"
  },
  "nickname": "My Sender ID",
  "reply_to": {
    "email": "replyto@example.com",
    "name": "Example INC"
  },
  "state": "Colorado",
  "zip": "80202"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders_get(t *testing.T) {
	request := getRequest("/v3/senders")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders__sender_id__patch(t *testing.T) {
	request := getRequest("/v3/senders/{sender_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "address": "123 Elm St.",
  "address_2": "Apt. 456",
  "city": "Denver",
  "country": "United States",
  "from": {
    "email": "from@example.com",
    "name": "Example INC"
  },
  "nickname": "My Sender ID",
  "reply_to": {
    "email": "replyto@example.com",
    "name": "Example INC"
  },
  "state": "Colorado",
  "zip": "80202"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders__sender_id__get(t *testing.T) {
	request := getRequest("/v3/senders/{sender_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders__sender_id__delete(t *testing.T) {
	request := getRequest("/v3/senders/{sender_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_senders__sender_id__resend_verification_post(t *testing.T) {
	request := getRequest("/v3/senders/{sender_id}/resend_verification")
	request.Method = "POST"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_stats_get(t *testing.T) {
	request := getRequest("/v3/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_post(t *testing.T) {
	request := getRequest("/v3/subusers")
	request.Method = "POST"
	request.Body = []byte(` {
  "email": "John@example.com",
  "ips": [
    "1.1.1.1",
    "2.2.2.2"
  ],
  "password": "johns_password",
  "username": "John@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_get(t *testing.T) {
	request := getRequest("/v3/subusers")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_reputations_get(t *testing.T) {
	request := getRequest("/v3/subusers/reputations")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["usernames"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_stats_get(t *testing.T) {
	request := getRequest("/v3/subusers/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["subusers"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_stats_monthly_get(t *testing.T) {
	request := getRequest("/v3/subusers/stats/monthly")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["subuser"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers_stats_sums_get(t *testing.T) {
	request := getRequest("/v3/subusers/stats/sums")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["end_date"] = "2016-04-01"
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["sort_by_direction"] = "asc"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__patch(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "disabled": false
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__delete(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__ips_put(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/ips")
	request.Method = "PUT"
	request.Body = []byte(` [
  "127.0.0.1"
]`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__monitor_put(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/monitor")
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 500
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__monitor_post(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/monitor")
	request.Method = "POST"
	request.Body = []byte(` {
  "email": "example@example.com",
  "frequency": 50000
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__monitor_get(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/monitor")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__monitor_delete(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/monitor")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_subusers__subuser_name__stats_monthly_get(t *testing.T) {
	request := getRequest("/v3/subusers/{subuser_name}/stats/monthly")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["date"] = "test_string"
	queryParams["sort_by_direction"] = "asc"
	queryParams["limit"] = "1"
	queryParams["sort_by_metric"] = "test_string"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_blocks_get(t *testing.T) {
	request := getRequest("/v3/suppression/blocks")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_blocks_delete(t *testing.T) {
	request := getRequest("/v3/suppression/blocks")
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_blocks__email__get(t *testing.T) {
	request := getRequest("/v3/suppression/blocks/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_blocks__email__delete(t *testing.T) {
	request := getRequest("/v3/suppression/blocks/{email}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_bounces_get(t *testing.T) {
	request := getRequest("/v3/suppression/bounces")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["end_time"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_bounces_delete(t *testing.T) {
	request := getRequest("/v3/suppression/bounces")
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": true,
  "emails": [
    "example@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_bounces__email__get(t *testing.T) {
	request := getRequest("/v3/suppression/bounces/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_bounces__email__delete(t *testing.T) {
	request := getRequest("/v3/suppression/bounces/{email}")
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["email_address"] = "example@example.com"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_invalid_emails_get(t *testing.T) {
	request := getRequest("/v3/suppression/invalid_emails")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_invalid_emails_delete(t *testing.T) {
	request := getRequest("/v3/suppression/invalid_emails")
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_invalid_emails__email__get(t *testing.T) {
	request := getRequest("/v3/suppression/invalid_emails/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_invalid_emails__email__delete(t *testing.T) {
	request := getRequest("/v3/suppression/invalid_emails/{email}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_spam_report__email__get(t *testing.T) {
	request := getRequest("/v3/suppression/spam_reports/{email}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_spam_report__email__delete(t *testing.T) {
	request := getRequest("/v3/suppression/spam_reports/{email}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_spam_reports_get(t *testing.T) {
	request := getRequest("/v3/suppression/spam_reports")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_spam_reports_delete(t *testing.T) {
	request := getRequest("/v3/suppression/spam_reports")
	request.Method = "DELETE"
	request.Body = []byte(` {
  "delete_all": false,
  "emails": [
    "example1@example.com",
    "example2@example.com"
  ]
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_suppression_unsubscribes_get(t *testing.T) {
	request := getRequest("/v3/suppression/unsubscribes")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["start_time"] = "1"
	queryParams["limit"] = "1"
	queryParams["end_time"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates_post(t *testing.T) {
	request := getRequest("/v3/templates")
	request.Method = "POST"
	request.Body = []byte(` {
  "name": "example_name"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates_get(t *testing.T) {
	request := getRequest("/v3/templates")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__patch(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "name": "new_example_name"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__get(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__delete(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__versions_post(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}/versions")
	request.Method = "POST"
	request.Body = []byte(` {
  "active": 1,
  "html_content": "<%body%>",
  "name": "example_version_name",
  "plain_content": "<%body%>",
  "subject": "<%subject%>",
  "template_id": "ddb96bbc-9b92-425e-8979-99464621b543"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__versions__version_id__patch(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}/versions/{version_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "active": 1,
  "html_content": "<%body%>",
  "name": "updated_example_name",
  "plain_content": "<%body%>",
  "subject": "<%subject%>"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__versions__version_id__get(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}/versions/{version_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__versions__version_id__delete(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}/versions/{version_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_templates__template_id__versions__version_id__activate_post(t *testing.T) {
	request := getRequest("/v3/templates/{template_id}/versions/{version_id}/activate")
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_get(t *testing.T) {
	request := getRequest("/v3/tracking_settings")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_click_patch(t *testing.T) {
	request := getRequest("/v3/tracking_settings/click")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_click_get(t *testing.T) {
	request := getRequest("/v3/tracking_settings/click")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_google_analytics_patch(t *testing.T) {
	request := getRequest("/v3/tracking_settings/google_analytics")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "utm_campaign": "website",
  "utm_content": "",
  "utm_medium": "email",
  "utm_source": "sendgrid.com",
  "utm_term": ""
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_google_analytics_get(t *testing.T) {
	request := getRequest("/v3/tracking_settings/google_analytics")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_open_patch(t *testing.T) {
	request := getRequest("/v3/tracking_settings/open")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_open_get(t *testing.T) {
	request := getRequest("/v3/tracking_settings/open")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_subscription_patch(t *testing.T) {
	request := getRequest("/v3/tracking_settings/subscription")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "enabled": true,
  "html_content": "html content",
  "landing": "landing page html",
  "plain_content": "text content",
  "replace": "replacement tag",
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_tracking_settings_subscription_get(t *testing.T) {
	request := getRequest("/v3/tracking_settings/subscription")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_account_get(t *testing.T) {
	request := getRequest("/v3/user/account")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_credits_get(t *testing.T) {
	request := getRequest("/v3/user/credits")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_email_put(t *testing.T) {
	request := getRequest("/v3/user/email")
	request.Method = "PUT"
	request.Body = []byte(` {
  "email": "example@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_email_get(t *testing.T) {
	request := getRequest("/v3/user/email")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_password_put(t *testing.T) {
	request := getRequest("/v3/user/password")
	request.Method = "PUT"
	request.Body = []byte(` {
  "new_password": "new_password",
  "old_password": "old_password"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_profile_patch(t *testing.T) {
	request := getRequest("/v3/user/profile")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "city": "Orange",
  "first_name": "Example",
  "last_name": "User"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_profile_get(t *testing.T) {
	request := getRequest("/v3/user/profile")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_scheduled_sends_post(t *testing.T) {
	request := getRequest("/v3/user/scheduled_sends")
	request.Method = "POST"
	request.Body = []byte(` {
  "batch_id": "YOUR_BATCH_ID",
  "status": "pause"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_scheduled_sends_get(t *testing.T) {
	request := getRequest("/v3/user/scheduled_sends")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_scheduled_sends__batch_id__patch(t *testing.T) {
	request := getRequest("/v3/user/scheduled_sends/{batch_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "status": "pause"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_scheduled_sends__batch_id__get(t *testing.T) {
	request := getRequest("/v3/user/scheduled_sends/{batch_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_scheduled_sends__batch_id__delete(t *testing.T) {
	request := getRequest("/v3/user/scheduled_sends/{batch_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_settings_enforced_tls_patch(t *testing.T) {
	request := getRequest("/v3/user/settings/enforced_tls")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "require_tls": true,
  "require_valid_cert": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_settings_enforced_tls_get(t *testing.T) {
	request := getRequest("/v3/user/settings/enforced_tls")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_username_put(t *testing.T) {
	request := getRequest("/v3/user/username")
	request.Method = "PUT"
	request.Body = []byte(` {
  "username": "test_username"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_username_get(t *testing.T) {
	request := getRequest("/v3/user/username")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_event_settings_patch(t *testing.T) {
	request := getRequest("/v3/user/webhooks/event/settings")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "bounce": true,
  "click": true,
  "deferred": true,
  "delivered": true,
  "dropped": true,
  "enabled": true,
  "group_resubscribe": true,
  "group_unsubscribe": true,
  "open": true,
  "processed": true,
  "spam_report": true,
  "unsubscribe": true,
  "url": "url"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_event_settings_get(t *testing.T) {
	request := getRequest("/v3/user/webhooks/event/settings")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_event_test_post(t *testing.T) {
	request := getRequest("/v3/user/webhooks/event/test")
	request.Method = "POST"
	request.Body = []byte(` {
  "url": "url"
}`)
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_settings_post(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/settings")
	request.Method = "POST"
	request.Body = []byte(` {
  "hostname": "myhostname.com",
  "send_raw": false,
  "spam_check": true,
  "url": "http://email.myhosthame.com"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_settings_get(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/settings")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_settings__hostname__patch(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/settings/{hostname}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "send_raw": true,
  "spam_check": false,
  "url": "http://newdomain.com/parse"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_settings__hostname__get(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/settings/{hostname}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_settings__hostname__delete(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/settings/{hostname}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_user_webhooks_parse_stats_get(t *testing.T) {
	request := getRequest("/v3/user/webhooks/parse/stats")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "test_string"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains")
	request.Method = "POST"
	request.Body = []byte(` {
  "automatic_security": false,
  "custom_spf": true,
  "default": true,
  "domain": "example.com",
  "ips": [
    "192.168.1.1",
    "192.168.1.2"
  ],
  "subdomain": "news",
  "username": "john@example.com"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	queryParams["domain"] = "test_string"
	queryParams["exclude_subusers"] = "true"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains_default_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/default")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains_subuser_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/subuser")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains_subuser_delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/subuser")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__domain_id__patch(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{domain_id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "custom_spf": true,
  "default": false
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__domain_id__get(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{domain_id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__domain_id__delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{domain_id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__domain_id__subuser_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{domain_id}/subuser")
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__id__ips_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{id}/ips")
	request.Method = "POST"
	request.Body = []byte(` {
  "ip": "192.168.0.1"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__id__ips__ip__delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{id}/ips/{ip}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_domains__id__validate_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/domains/{id}/validate")
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_ips_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/ips")
	request.Method = "POST"
	request.Body = []byte(` {
  "domain": "example.com",
  "ip": "192.168.1.1",
  "subdomain": "email"
}`)
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_ips_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/ips")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["ip"] = "test_string"
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_ips__id__get(t *testing.T) {
	request := getRequest("/v3/whitelabel/ips/{id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_ips__id__delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/ips/{id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_ips__id__validate_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/ips/{id}/validate")
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/links")
	request.Method = "POST"
	request.Body = []byte(` {
  "default": true,
  "domain": "example.com",
  "subdomain": "mail"
}`)
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "201"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 201, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/links")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["limit"] = "1"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links_default_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/default")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["domain"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links_subuser_get(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/subuser")
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links_subuser_delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/subuser")
	request.Method = "DELETE"
	queryParams := make(map[string]string)
	queryParams["username"] = "test_string"
	request.QueryParams = queryParams
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links__id__patch(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/{id}")
	request.Method = "PATCH"
	request.Body = []byte(` {
  "default": true
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links__id__get(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/{id}")
	request.Method = "GET"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links__id__delete(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/{id}")
	request.Method = "DELETE"
	request.Headers["X-Mock"] = "204"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 204, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links__id__validate_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/{id}/validate")
	request.Method = "POST"
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}

func Test_test_whitelabel_links__link_id__subuser_post(t *testing.T) {
	request := getRequest("/v3/whitelabel/links/{link_id}/subuser")
	request.Method = "POST"
	request.Body = []byte(` {
  "username": "jane@example.com"
}`)
	request.Headers["X-Mock"] = "200"
	response, err := MakeRequest(request)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, 200, response.StatusCode, "Wrong status code returned")
}
