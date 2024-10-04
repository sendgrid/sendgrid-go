package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestPageUtil_ReadLimits(t *testing.T) {
	assert.Equal(t, 5, ReadLimits(nil, setLimit(5)))
	assert.Equal(t, 5, ReadLimits(setPageSize(10), setLimit(5)))
	assert.Equal(t, 1000, ReadLimits(nil, setLimit(5000)))
	assert.Equal(t, 10, ReadLimits(setPageSize(10), nil))
	assert.Equal(t, 50, ReadLimits(nil, nil))
}

func setLimit(limit int) *int {
	return &limit
}

func setPageSize(pageSize int) *int {
	return &pageSize
}

func TestPageUtil_GetNextPageUri(t *testing.T) {
	payload := map[string]interface{}{
		"next_page_uri": "/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
		"page_size":     50,
	}
	baseUrl := "https://api.twilio.com/"
	nextPageUrl, err := getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload["next_page_uri"] = "2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1"
	baseUrl = "https://api.twilio.com"
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)

	payload = map[string]interface{}{}
	nextPageUrl, err = getNextPageUrl(baseUrl, payload)
	assert.Nil(t, err)
	assert.Equal(t, "", nextPageUrl)
}

func TestPageUtil_GetNextPageUrl(t *testing.T) {
	payload := map[string]interface{}{
		"meta": map[string]interface{}{
			"next_page_url": "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1",
			"page_size":     50,
		},
	}

	nextPageUrl, err := getNextPageUrl("https://apitest.twilio.com", payload)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/IncomingPhoneNumbers.json?PageSize=50&Page=1", nextPageUrl)
}

func getTestClient(t *testing.T) *MockBaseClient {
	mockCtrl := gomock.NewController(t)
	testClient := NewMockBaseClient(mockCtrl)
	testClient.EXPECT().AccountSid().DoAndReturn(func() string {
		return "AC222222222222222222222222222222"
	}).AnyTimes()

	testClient.EXPECT().SendRequest(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any()).
		DoAndReturn(func(method string, rawURL string, data url.Values,
			headers map[string]interface{}, body ...interface{}) (*http.Response, error) {
			response := map[string]interface{}{
				"end":            4,
				"first_page_uri": "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=0",
				"messages": []map[string]interface{}{
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 0",
						"status":    "delivered",
					},
					{
						"direction": "outbound-api",
						"from":      "4444444444",
						"to":        "9999999999",
						"body":      "Message 1",
						"status":    "delivered",
					},
				},
				"uri":           "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=0&PageToken=dummy",
				"page_size":     5,
				"start":         0,
				"next_page_uri": "/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=1&PageToken=PASMXX",
				"page":          0,
			}

			resp, _ := json.Marshal(response)

			return &http.Response{
				Body: io.NopCloser(bytes.NewReader(resp)),
			}, nil
		},
		)

	return testClient
}

type testResponse struct {
	End             int           `json:"end,omitempty"`
	FirstPageUri    string        `json:"first_page_uri,omitempty"`
	Messages        []testMessage `json:"messages,omitempty"`
	NextPageUri     string        `json:"next_page_uri,omitempty"`
	Page            int           `json:"page,omitempty"`
	PageSize        int           `json:"page_size,omitempty"`
	PreviousPageUri string        `json:"previous_page_uri,omitempty"`
	Start           int           `json:"start,omitempty"`
	Uri             string        `json:"uri,omitempty"`
}

type testMessage struct {
	// The message text
	Body *string `json:"body,omitempty"`
	// The direction of the message
	Direction *string `json:"direction,omitempty"`
	// The phone number that initiated the message
	From *string `json:"from,omitempty"`
	// The status of the message
	Status *string `json:"status,omitempty"`
	// The phone number that received the message
	To *string `json:"to,omitempty"`
}

func getSomething(nextPageUrl string) (interface{}, error) {
	return nextPageUrl, nil
}

func TestPageUtil_GetNext(t *testing.T) {
	testClient := getTestClient(t)
	baseUrl := "https://api.twilio.com"
	response, _ := testClient.SendRequest("get", "", nil, nil) //nolint:bodyclose
	ps := &testResponse{}
	_ = json.NewDecoder(response.Body).Decode(ps)

	nextPageUrl, err := GetNext(baseUrl, ps, getSomething)
	assert.Equal(t, "https://api.twilio.com/2010-04-01/Accounts/ACXX/Messages.json?From=9999999999&PageNumber=&To=4444444444&PageSize=2&Page=1&PageToken=PASMXX", nextPageUrl)
	assert.Nil(t, err)

	nextPageUrl, err = GetNext(baseUrl, nil, getSomething)
	assert.Empty(t, nextPageUrl)
	assert.Nil(t, err)
}

func TestPageUtil_ToMap(t *testing.T) {
	testMap, err := toMap("invalid")
	assert.NotNil(t, err)
	assert.Nil(t, testMap)

	valid := testResponse{
		End:             0,
		FirstPageUri:    "first",
		Messages:        nil,
		NextPageUri:     "next",
		Page:            0,
		PageSize:        0,
		PreviousPageUri: "previous",
		Start:           0,
		Uri:             "uri",
	}
	testMap, err = toMap(valid)
	assert.Nil(t, err)
	assert.NotNil(t, testMap)
}
