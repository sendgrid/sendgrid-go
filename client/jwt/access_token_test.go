package jwt

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var AccountSid string
var SigningKeySid string
var Params AccessTokenParams

func TestMain(m *testing.M) {
	AccountSid = "AC123"
	SigningKeySid = "SK123"
	Params = AccessTokenParams{
		AccountSid:    AccountSid,
		SigningKeySid: SigningKeySid,
		Secret:        "secret",
	}
	ret := m.Run()
	os.Exit(ret)
}

func validateClaims(t *testing.T, payload map[string]interface{}) {
	assert.Equal(t, SigningKeySid, payload["iss"])
	assert.Equal(t, AccountSid, payload["sub"])
	assert.NotZero(t, payload["jti"])
	assert.NotZero(t, payload["exp"])
	assert.NotNil(t, payload["grants"])
	now := float64(time.Now().Unix())
	assert.Greater(t, payload["exp"].(float64), now)
	assert.True(t, strings.Contains(payload["jti"].(string), payload["iss"].(string)))
}

func testInit(t *testing.T) *AccessToken {
	accessToken := CreateAccessToken(Params)
	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	return decodedToken
}

func assertPanic(t *testing.T, f func(grant BaseGrant)) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f(nil)
}

func TestAccessTokenSimple(t *testing.T) {
	params := Params
	params.Identity = "identity"
	params.Nbf = 3600
	accessToken := CreateAccessToken(params)

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Equal(t, "secret", decodedToken.baseJwt.SecretKey)
	validateClaims(t, decodedToken.Payload())
}

func TestAccessTokenWithNoValidation(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Equal(t, "", decodedToken.baseJwt.SecretKey)
	validateClaims(t, decodedToken.Payload())
}

func TestAccessTokenWithoutSecret(t *testing.T) {
	params := Params
	params.Secret = ""
	accessToken := CreateAccessToken(params)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_, _ = accessToken.ToJwt()
}

func TestAccessTokenHeaders(t *testing.T) {
	params := Params
	params.Identity = "identity"
	params.Region = "US"
	params.Nbf = 0
	params.Ttl = 50
	params.ValidUntil = 0
	accessToken := CreateAccessToken(params)

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Equal(t, "secret", decodedToken.baseJwt.SecretKey)

	headers := decodedToken.Headers()

	assert.Equal(t, "US", headers["twr"])
	assert.Equal(t, "HS256", headers["alg"])
	assert.Equal(t, "twilio-fpa;v=1", headers["cty"])
	assert.Equal(t, "JWT", headers["typ"])
}

func TestParams(t *testing.T) {
	now := float64(time.Now().Unix())
	params := Params
	params.Identity = "identity"
	params.Region = "US"
	params.Nbf = now
	params.Ttl = 7200
	params.ValidUntil = now + 10800
	accessToken := CreateAccessToken(params)
	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	payload := decodedToken.Payload()
	assert.Equal(t, now, payload["nbf"])
	assert.Equal(t, now+10800, payload["exp"])
}

func TestAccessTokenPayload(t *testing.T) {
	now := float64(time.Now().Unix())
	params := Params
	params.Identity = "identity"
	params.Region = "US"
	params.Nbf = 360
	params.Ttl = 50
	params.ValidUntil = 200 + now
	accessToken := CreateAccessToken(params)
	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Equal(t, "secret", decodedToken.baseJwt.SecretKey)
	assert.Equal(t, "US", decodedToken.Region)
	assert.Equal(t, "identity", decodedToken.Identity)
	assert.Equal(t, AccountSid, decodedToken.AccountSid)
	assert.Equal(t, SigningKeySid, decodedToken.SigningKeySid)

	payload := decodedToken.Payload()
	assert.Equal(t, "SK123", payload["iss"])
	assert.Equal(t, float64(360), payload["nbf"])
	assert.Equal(t, "AC123", payload["sub"])
	assert.True(t, strings.Contains(payload["jti"].(string), payload["iss"].(string)))
	assert.Len(t, decodedToken.Grants, 0)
}

func TestHeaders(t *testing.T) {
	decodedToken := testInit(t)
	headers := decodedToken.Headers()
	assert.Equal(t, "HS256", headers["alg"])
	assert.Equal(t, "twilio-fpa;v=1", headers["cty"])
	assert.Equal(t, "JWT", headers["typ"])
}

func TestChatGrant(t *testing.T) {
	params := Params
	params.Identity = "identity"
	params.Region = "US"
	params.Nbf = 3600
	accessToken := CreateAccessToken(params)

	accessToken.AddGrant(&ChatGrant{
		ServiceSid:        "IS123",
		EndpointID:        "Endpoint123",
		DeploymentRoleSid: "Role123",
		PushCredentialSid: "CR123",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	// identity should exist in the grants map as well
	assert.Len(t, payload["grants"], 2)

	chatGrantDecoded := payload["grants"].(map[string]interface{})["chat"].(map[string]interface{})
	assert.NotNil(t, chatGrantDecoded)
	assert.Equal(t, "Role123", chatGrantDecoded["deployment_role_sid"])
	assert.Equal(t, "Endpoint123", chatGrantDecoded["endpoint_id"])
	assert.Equal(t, "CR123", chatGrantDecoded["push_credential_sid"])
	assert.Equal(t, "IS123", chatGrantDecoded["service_sid"])
}

func TestConversationsGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&ConversationsGrant{ConfigurationProfileSid: "CP123"})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	payload := decodedToken.Payload()
	assert.Len(t, decodedToken.Grants, 1)
	assert.Len(t, payload["grants"], 1)

	conversationsGrantDecoded := payload["grants"].(map[string]interface{})["rtc"].(map[string]interface{})
	assert.NotNil(t, conversationsGrantDecoded)
	assert.Equal(t, "CP123", conversationsGrantDecoded["configuration_profile_sid"])
}

func TestIpMessagingGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&IpMessagingGrant{
		ServiceSid:        "IS123",
		EndpointID:        "Endpoint123",
		DeploymentRoleSid: "Role123",
		PushCredentialSid: "CR123",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	ipMessagingGrantDecoded := payload["grants"].(map[string]interface{})["ip_messaging"].(map[string]interface{})
	assert.NotNil(t, ipMessagingGrantDecoded)
	assert.Equal(t, "IS123", ipMessagingGrantDecoded["service_sid"])
	assert.Equal(t, "Endpoint123", ipMessagingGrantDecoded["endpoint_id"])
	assert.Equal(t, "Role123", ipMessagingGrantDecoded["deployment_role_sid"])
	assert.Equal(t, "CR123", ipMessagingGrantDecoded["push_credential_sid"])
}

func TestSyncGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&SyncGrant{
		ServiceSid: "IS123",
		EndpointID: "Endpoint123",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	syncGrantDecoded := payload["grants"].(map[string]interface{})["data_sync"].(map[string]interface{})
	assert.NotNil(t, syncGrantDecoded)
	assert.Equal(t, "IS123", syncGrantDecoded["service_sid"])
	assert.Equal(t, "Endpoint123", syncGrantDecoded["endpoint_id"])
}

func TestVideoGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&VideoGrant{Room: "RM123"})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	videoGrantDecoded := payload["grants"].(map[string]interface{})["video"].(map[string]interface{})
	assert.NotNil(t, videoGrantDecoded)
	assert.Equal(t, "RM123", videoGrantDecoded["room"])
}

func TestVoiceGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&VoiceGrant{
		Incoming: Incoming{Allow: true},
		Outgoing: Outgoing{
			ApplicationSid: "SID123",
			ApplicationParams: map[string]interface{}{
				"foo": "bar",
			},
		},
		PushCredentialSid: "Push123",
		EndpointID:        "Endpoint123",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	voiceGrantDecoded := payload["grants"].(map[string]interface{})["voice"].(map[string]interface{})
	assert.NotNil(t, voiceGrantDecoded)
	assert.Equal(t, "Endpoint123", voiceGrantDecoded["endpoint_id"])
	assert.Equal(t, "Push123", voiceGrantDecoded["push_credential_sid"])

	incoming := voiceGrantDecoded["incoming"].(map[string]interface{})
	assert.Equal(t, true, incoming["allow"])
	outgoing := voiceGrantDecoded["outgoing"].(map[string]interface{})
	assert.Equal(t, "SID123", outgoing["application_sid"])
	assert.Equal(t, "bar", outgoing["params"].(map[string]interface{})["foo"])
}

func TestTaskRouterGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&TaskRouterGrant{
		WorkspaceSid: "WS123",
		WorkerSid:    "WK123",
		Role:         "worker",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	taskRouterGrantDecoded := payload["grants"].(map[string]interface{})["task_router"].(map[string]interface{})
	assert.NotNil(t, taskRouterGrantDecoded)
	assert.Equal(t, "WS123", taskRouterGrantDecoded["workspace_sid"])
	assert.Equal(t, "WK123", taskRouterGrantDecoded["worker_sid"])
	assert.Equal(t, "worker", taskRouterGrantDecoded["role"])
}

func TestPlaybackGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&PlaybackGrant{
		"requestCredentials": nil,
		"playbackUrl":        "https://000.us-east-1.playback.live-video.net/api/video/v1/us-east-000.channel.000?token=xxxxx",
		"playerStreamerSid":  "VJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	assert.Len(t, decodedToken.Grants, 1)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 1)

	playbackGrantDecoded := payload["grants"].(map[string]interface{})["player"].(map[string]interface{})
	assert.NotNil(t, playbackGrantDecoded)
	assert.Equal(t, nil, playbackGrantDecoded["requestCredentials"])
	assert.Equal(t, "https://000.us-east-1.playback.live-video.net/api/video/v1/us-east-000.channel.000?token=xxxxx", playbackGrantDecoded["playbackUrl"])
	assert.Equal(t, "VJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", playbackGrantDecoded["playerStreamerSid"])
}

func TestMultipleGrants(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	accessToken.AddGrant(&TaskRouterGrant{
		WorkspaceSid: "WS123",
		WorkerSid:    "WK123",
		Role:         "worker",
	})

	accessToken.AddGrant(&SyncGrant{})

	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)
	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 2)

	taskRouterGrantDecoded := payload["grants"].(map[string]interface{})["task_router"].(map[string]interface{})
	assert.NotNil(t, taskRouterGrantDecoded)
	assert.Equal(t, "WS123", taskRouterGrantDecoded["workspace_sid"])
	assert.Equal(t, "WK123", taskRouterGrantDecoded["worker_sid"])
	assert.Equal(t, "worker", taskRouterGrantDecoded["role"])

	syncGrantDecoded := payload["grants"].(map[string]interface{})["data_sync"].(map[string]interface{})
	assert.NotNil(t, syncGrantDecoded)
}

func TestGrantsDuringInit(t *testing.T) {
	grants := []BaseGrant{
		&VideoGrant{Room: "room"},
		&VoiceGrant{},
	}
	params := Params
	params.Grants = grants
	accessToken := CreateAccessToken(params)
	token, err := accessToken.ToJwt()
	assert.Nil(t, err)
	assert.NotNil(t, token)

	decodedToken, err := accessToken.FromJwt(token, "secret")
	assert.Nil(t, err)
	assert.NotNil(t, decodedToken)

	payload := decodedToken.Payload()
	assert.Len(t, payload["grants"], 2)

	videoGrantDecoded := payload["grants"].(map[string]interface{})["video"].(map[string]interface{})
	assert.NotNil(t, videoGrantDecoded)
	assert.Equal(t, "room", videoGrantDecoded["room"])

	voiceGrantDecoded := payload["grants"].(map[string]interface{})["voice"].(map[string]interface{})
	assert.NotNil(t, voiceGrantDecoded)
}

func TestValidateGrant(t *testing.T) {
	accessToken := CreateAccessToken(Params)
	assertPanic(t, accessToken.AddGrant)
}

func TestGrantsToString(t *testing.T) {
	chatGrant := &ChatGrant{}
	assert.True(t, strings.HasPrefix(chatGrant.ToString(), "<ChatGrant"))

	conversationsGrant := &ConversationsGrant{}
	assert.True(t, strings.HasPrefix(conversationsGrant.ToString(), "<ConversationsGrant"))

	ipMessagingGrant := &IpMessagingGrant{}
	assert.True(t, strings.HasPrefix(ipMessagingGrant.ToString(), "<IpMessagingGrant"))

	syncGrant := &SyncGrant{}
	assert.True(t, strings.HasPrefix(syncGrant.ToString(), "<SyncGrant"))

	voiceGrant := &VoiceGrant{}
	assert.True(t, strings.HasPrefix(voiceGrant.ToString(), "<VoiceGrant"))

	videoGrant := &VideoGrant{}
	assert.True(t, strings.HasPrefix(videoGrant.ToString(), "<VideoGrant"))

	playbackGrant := &PlaybackGrant{}
	assert.True(t, strings.HasPrefix(playbackGrant.ToString(), "<PlaybackGrant"))
}
