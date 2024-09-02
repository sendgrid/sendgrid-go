package taskrouter

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var AccountSid string
var AuthToken string
var WorkspaceSid string
var TaskqueueSid string
var Params CapabilityTokenParams

func TestMain(m *testing.M) {
	AccountSid = "AC123"
	AuthToken = "foobar"
	WorkspaceSid = "WS123"
	TaskqueueSid = "WQ123"
	Params = CapabilityTokenParams{
		AccountSid:   AccountSid,
		AuthToken:    AuthToken,
		WorkspaceSid: WorkspaceSid,
		ChannelID:    TaskqueueSid,
	}
	ret := m.Run()
	os.Exit(ret)
}

func TestCapabilityTokenGenerate(t *testing.T) {
	capabilityToken := CreateCapabilityToken(Params)

	token, _ := capabilityToken.ToJwt()
	assert.NotNil(t, token)

	decodedToken, err := capabilityToken.FromJwt(token, AuthToken)
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	payload := decodedToken.Payload()

	assert.Equal(t, AccountSid, payload["iss"])
	assert.Len(t, payload["policies"], 2)
	assert.Equal(t, AccountSid, payload["iss"])
	assert.Equal(t, TaskqueueSid, payload["channel"])
	assert.Equal(t, WorkspaceSid, payload["workspace_sid"])
	assert.Equal(t, "v1", payload["version"])
	assert.Equal(t, TaskqueueSid, payload["friendly_name"])
}

func TestHeaders(t *testing.T) {
	capabilityToken := CreateCapabilityToken(Params)
	token, _ := capabilityToken.ToJwt()
	decodedToken, err := capabilityToken.FromJwt(token, AuthToken)
	assert.Nil(t, err)
	headers := decodedToken.Headers()
	assert.Equal(t, "HS256", headers["alg"])
	assert.Equal(t, "JWT", headers["typ"])
}

func TestToString(t *testing.T) {
	capabilityToken := CreateCapabilityToken(Params)
	tokenStr := capabilityToken.ToString()
	assert.True(t, strings.HasPrefix(tokenStr, "<TaskRouterCapabilityToken"))
}

func checkPolicy(t *testing.T, policy interface{}, index int, method string, url string, postFilters map[string]interface{}, queryFilters map[string]interface{}) {
	switch reflect.TypeOf(policy).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(policy)
		policyToTest := s.Index(index).Interface().(map[string]interface{})
		assert.Equal(t, policyToTest["url"].(string), url)
		assert.Equal(t, policyToTest["method"].(string), method)
		assert.Len(t, policyToTest["post_filter"], 0)
		assert.Len(t, policyToTest["query_filter"], 0)
	}
}

func TestAddingPolicyToToken(t *testing.T) {
	capabilityToken := CreateCapabilityToken(Params)

	policy := GeneratePolicy(Workspace(WorkspaceSid), "GET", true, nil, nil)
	capabilityToken.AddPolicy(policy)
	token, _ := capabilityToken.ToJwt()
	assert.NotNil(t, token)

	decodedToken, err := capabilityToken.FromJwt(token, AuthToken)
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)

	assert.Len(t, decodedToken.Policies, 3)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123", decodedToken.Policies[0].Url)
	assert.Equal(t, "https://event-bridge.twilio.com/v1/wschannels/AC123/WQ123", decodedToken.Policies[1].Url)
	assert.Equal(t, "https://event-bridge.twilio.com/v1/wschannels/AC123/WQ123", decodedToken.Policies[2].Url)
	assert.Equal(t, "GET", decodedToken.Policies[0].Method)
	assert.Equal(t, "GET", decodedToken.Policies[1].Method)
	assert.Equal(t, "POST", decodedToken.Policies[2].Method)

	payload := decodedToken.Payload()
	assert.GreaterOrEqual(t, len(payload), 1)
	assert.NotNil(t, payload["policies"])
	assert.Len(t, payload["policies"], 3)

	checkPolicy(t, payload["policies"], 0, "GET", "https://taskrouter.twilio.com/v1/Workspaces/WS123", nil, nil)
	checkPolicy(t, payload["policies"], 1, "GET", "https://event-bridge.twilio.com/v1/wschannels/AC123/WQ123", nil, nil)
	checkPolicy(t, payload["policies"], 2, "POST", "https://event-bridge.twilio.com/v1/wschannels/AC123/WQ123", nil, nil)
}

func TestPolicyFilter(t *testing.T) {
	postFilter := map[string]interface{}{
		"test-post-key": "test-post-value",
		"activity-sid": map[string]interface{}{
			"required": true,
		},
	}
	queryFilter := map[string]interface{}{
		"test-query-key": "test-query-value",
	}

	policy := GeneratePolicy(Workspace(WorkspaceSid), "GET", true, postFilter, queryFilter)
	payload := policy.Payload()
	assert.Equal(t, "test-post-value", payload["post_filter"].(map[string]interface{})["test-post-key"])
	assert.Equal(t, "test-query-value", payload["query_filter"].(map[string]interface{})["test-query-key"])
	assert.Equal(t, true, payload["post_filter"].(map[string]interface{})["activity-sid"].(map[string]interface{})["required"])

	capabilityToken := CreateCapabilityToken(Params)
	capabilityToken.AddPolicy(policy)
	token, _ := capabilityToken.ToJwt()
	assert.NotNil(t, token)
	decodedToken, err := capabilityToken.FromJwt(token, AuthToken)
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)

	assert.Len(t, decodedToken.Policies, 3)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123", decodedToken.Policies[0].Url)
	assert.Equal(t, "GET", decodedToken.Policies[0].Method)
	assert.Equal(t, postFilter, decodedToken.Policies[0].PostFilter)
}

func TestWithNoFilters(t *testing.T) {
	policy := GeneratePolicy(Workspace(WorkspaceSid), "GET", true, nil, nil)
	payload := policy.Payload()
	assert.Len(t, payload["post_filter"], 0)
	assert.Len(t, payload["query_filter"], 0)

	params := CapabilityTokenParams{
		AccountSid:   "AC123",
		AuthToken:    "XXX",
		WorkspaceSid: "WS123",
		ChannelID:    "CH123",
	}

	capabilityToken := CreateCapabilityToken(params)
	capabilityToken.AddPolicy(policy)

	jwt, _ := capabilityToken.ToJwt()
	assert.NotNil(t, jwt)
}

func TestPolicyWithPostFilter(t *testing.T) {
	params := CapabilityTokenParams{
		AccountSid:   AccountSid,
		AuthToken:    AuthToken,
		WorkspaceSid: WorkspaceSid,
		ChannelID:    "CH123",
	}
	capabilityToken := CreateCapabilityToken(params)
	postFilter := map[string]interface{}{
		"ActivitySid": map[string]interface{}{
			"required": true,
		},
	}
	workerPolicy := GeneratePolicy(Worker(WorkspaceSid, "WorkerSid"), "POST", true, postFilter, nil)
	capabilityToken.AddPolicy(workerPolicy)

	workerToken, err := capabilityToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, workerToken)

	decodedToken, errDecode := capabilityToken.FromJwt(workerToken, AuthToken)
	assert.Nil(t, errDecode)
	assert.NotNil(t, decodedToken)
	assert.Equal(t, true, decodedToken.Policies[0].PostFilter["ActivitySid"].(map[string]interface{})["required"])
}

func TestUrls(t *testing.T) {
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces", Workspaces())
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123", Workspace("WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/**", AllWorkspaces())
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/Tasks", Tasks("WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/Tasks/TK123", Task("WK123", "TK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/Tasks/**", AllTasks("WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/TaskQueues", TaskQueues("WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/TaskQueues/WQ123", TaskQueue("WK123", "WQ123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WK123/TaskQueues/**", AllTaskQueues("WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Activities", Activities("WS123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Activities/AT123", Activity("WS123", "AT123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Activities/**", AllActivities("WS123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers", Workers("WS123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123", Worker("WS123", "WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/**", AllWorkers("WS123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123/Reservations", Reservations("WS123", "WK123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123/Reservations/RS123", Reservation("WS123", "WK123", "RS123"))
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123/Reservations/**", AllReservations("WS123", "WK123"))
}

func TestDefaultWorkerPolicies(t *testing.T) {
	workerPolicies := WorkerPolicies("WS123", "WK123")
	assert.Equal(t, "GET", workerPolicies[0].Method)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Activities", workerPolicies[0].Url)
	assert.Equal(t, "GET", workerPolicies[1].Method)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Tasks/**", workerPolicies[1].Url)
	assert.Equal(t, "GET", workerPolicies[2].Method)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123/Reservations/**", workerPolicies[2].Url)
	assert.Equal(t, "GET", workerPolicies[3].Method)
	assert.Equal(t, "https://taskrouter.twilio.com/v1/Workspaces/WS123/Workers/WK123", workerPolicies[3].Url)
}

func TestDefaultWebSocketPolicies(t *testing.T) {
	webSocketPolicies := WebSocketPolicies("AC123", "CI123")
	assert.Equal(t, "GET", webSocketPolicies[0].Method)
	assert.Equal(t, "https://event-bridge.twilio.com/v1/wschannels/AC123/CI123", webSocketPolicies[0].Url)
	assert.Equal(t, "POST", webSocketPolicies[1].Method)
	assert.Equal(t, "https://event-bridge.twilio.com/v1/wschannels/AC123/CI123", webSocketPolicies[1].Url)
}
