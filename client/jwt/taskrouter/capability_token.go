package taskrouter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	baseJwt "github.com/twilio/twilio-go/client/jwt"
	. "github.com/twilio/twilio-go/client/jwt/util"
)

type CapabilityToken struct {
	baseJwt      *baseJwt.Jwt
	AccountSid   string
	AuthToken    string
	WorkspaceSid string
	ChannelID    string
	Policies     []Policy
}

type CapabilityTokenParams struct {
	// Twilio Account sid
	AccountSid string
	// Twilio auth token used to sign the JWT
	AuthToken string
	// TaskRouter Workspace SID
	WorkspaceSid string
	// TaskRouter Channel SID
	ChannelID string
	// Time in secs since epoch before which this JWT is invalid, defaults to now
	Nbf float64
	// Time to live of the JWT in seconds, defaults to 1 hour
	Ttl float64
	// Time in secs since epoch this JWT is valid for. Overrides ttl if provided.
	ValidUntil float64
}

// Create Capability Token for TaskRouter
func CreateCapabilityToken(params CapabilityTokenParams) CapabilityToken {
	return CapabilityToken{
		baseJwt: &baseJwt.Jwt{
			SecretKey:  params.AuthToken,
			Issuer:     params.AccountSid,
			Subject:    "",
			Algorithm:  HS256,
			Nbf:        params.Nbf,
			Ttl:        Max(params.Ttl, 3600),
			ValidUntil: params.ValidUntil,
		},
		AccountSid:   params.AccountSid,
		AuthToken:    params.AuthToken,
		WorkspaceSid: params.WorkspaceSid,
		ChannelID:    params.ChannelID,
		Policies:     make([]Policy, 0),
	}
}

func (token *CapabilityToken) AddPolicy(policy Policy) {
	token.Policies = append(token.Policies, policy)
}

func (token *CapabilityToken) generatePayload() map[string]interface{} {
	now := float64(time.Now().Unix())

	// These are required since we want to authenticate and authorize the opening of a websocket in the first place.
	// Subsequent events to GET, POST or DELETE to other APIs will utilize this websocket.
	defaultPolicies := WebSocketPolicies(token.AccountSid, token.ChannelID)
	token.Policies = append(token.Policies, defaultPolicies...)

	payload := map[string]interface{}{
		"version": Version,
	}
	if token.AccountSid != "" {
		payload["account_sid"] = token.AccountSid
	}
	if token.WorkspaceSid != "" {
		payload["workspace_sid"] = token.WorkspaceSid
	}
	if token.ChannelID != "" {
		payload["channel"] = token.ChannelID
		payload["friendly_name"] = token.ChannelID
	}

	var policies []map[string]interface{}
	for _, policy := range token.Policies {
		policyPayload := policy.Payload()
		policies = append(policies, policyPayload)
	}

	if len(policies) > 0 {
		payload["policies"] = policies
	}

	payload["iss"] = token.baseJwt.Issuer
	payload["exp"] = now + token.baseJwt.Ttl
	if token.baseJwt.Nbf != 0 {
		payload["nbf"] = token.baseJwt.Nbf
	} else {
		payload["nbf"] = now
	}
	if token.baseJwt.ValidUntil != 0 {
		payload["exp"] = token.baseJwt.ValidUntil
	}
	if strings.HasPrefix(token.ChannelID, "WK") {
		payload["worker_sid"] = token.ChannelID
	} else if strings.HasPrefix(token.ChannelID, "WQ") {
		payload["taskqueue_sid"] = token.ChannelID
	}

	return payload
}

func (token *CapabilityToken) ToString() string {
	signedStr, err := token.ToJwt()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("<TaskRouterCapabilityToken %s>", signedStr)
}

// Encode the JWT struct into a string.
func (token *CapabilityToken) ToJwt() (string, error) {
	signedToken, err := token.baseJwt.ToJwt(token.baseJwt.Headers, token.generatePayload)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Get the decoded token back from the jwt String
func (token *CapabilityToken) FromJwt(jwtStr string, key string) (*CapabilityToken, error) {
	baseToken, err := token.baseJwt.FromJwt(jwtStr, key)
	if err != nil {
		return nil, err
	}

	return &CapabilityToken{
		baseJwt:      baseToken,
		AccountSid:   baseToken.Issuer,
		AuthToken:    baseToken.SecretKey,
		WorkspaceSid: baseToken.Payload()["workspace_sid"].(string),
		ChannelID:    baseToken.Payload()["channel"].(string),
		Policies:     decodePolicies(baseToken.Payload()["policies"]),
	}, nil
}

func decodePolicies(policies interface{}) []Policy {
	var decodedPolicies []Policy
	switch reflect.TypeOf(policies).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(policies)

		for i := 0; i < s.Len(); i++ {
			var pol Policy
			val := s.Index(i).Interface().(map[string]interface{})
			if data, err := json.Marshal(val); err == nil {
				if errJson := json.Unmarshal(data, &pol); errJson == nil {
					decodedPolicies = append(decodedPolicies, pol)
				}
			}
		}
	}

	return decodedPolicies
}

func (token *CapabilityToken) Headers() map[string]interface{} {
	if token.baseJwt.DecodedHeaders == nil {
		token.baseJwt.DecodedHeaders = token.baseJwt.Headers()
	}

	return token.baseJwt.DecodedHeaders
}

func (token *CapabilityToken) Payload() map[string]interface{} {
	if token.baseJwt.DecodedPayload == nil {
		token.baseJwt.DecodedPayload = token.generatePayload()
	}

	return token.baseJwt.DecodedPayload
}
