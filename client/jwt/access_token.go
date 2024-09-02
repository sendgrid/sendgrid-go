package jwt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	. "github.com/twilio/twilio-go/client/jwt/util"
)

type AccessToken struct {
	baseJwt *Jwt
	// List of permissions that the token grants
	Grants []BaseGrant `json:"grants,omitempty"`
	// Twilio Account SID
	AccountSid string `json:"account_sid,omitempty"`
	// API key
	SigningKeySid string `json:"signing_key_sid,omitempty"`
	// User's identity
	Identity string `json:"identity,omitempty"`
	// User's region
	Region interface{} `json:"region,omitempty"`
}

type AccessTokenParams struct {
	// Twilio Account sid
	AccountSid string
	// The issuer of the token
	SigningKeySid string
	// The secret used to sign the token
	Secret string
	// Identity of the token issuer
	Identity string
	// User's Region
	Region string
	// Time in secs since epoch before which this JWT is invalid, defaults to now
	Nbf float64
	// Time to live of the JWT in seconds, defaults to 1 hour
	Ttl float64
	// Time in secs since epoch this JWT is valid for. Overrides ttl if provided.
	ValidUntil float64
	// Access permissions granted to this token
	Grants []BaseGrant
}

func CreateAccessToken(params AccessTokenParams) AccessToken {
	return AccessToken{
		baseJwt: &Jwt{
			SecretKey:  params.Secret,
			Issuer:     params.SigningKeySid,
			Subject:    params.AccountSid,
			Algorithm:  HS256,
			Nbf:        params.Nbf,
			Ttl:        Max(params.Ttl, 3600),
			ValidUntil: params.ValidUntil,
		},
		Grants:        params.Grants,
		AccountSid:    params.AccountSid,
		SigningKeySid: params.SigningKeySid,
		Identity:      params.Identity,
		Region:        params.Region,
	}
}

func (token *AccessToken) Payload() map[string]interface{} {
	if token.baseJwt.DecodedPayload == nil {
		token.baseJwt.DecodedPayload = token.GeneratePayload()
	}

	return token.baseJwt.DecodedPayload
}

func (token *AccessToken) AddGrant(grant BaseGrant) {
	if grant == nil {
		panic("Grant to add is nil")
	}
	token.Grants = append(token.Grants, grant)
}

func (token *AccessToken) Headers() map[string]interface{} {
	if token.baseJwt.DecodedHeaders == nil {
		token.baseJwt.DecodedHeaders = token.generateHeaders()
	}

	return token.baseJwt.DecodedHeaders
}

func (token *AccessToken) generateHeaders() map[string]interface{} {
	headers := make(map[string]interface{})
	headers["cty"] = CType

	if token.Region != "" {
		headers["twr"] = token.Region
	}

	headers["alg"] = HS256
	headers["typ"] = JWT

	return headers
}

func (token *AccessToken) GeneratePayload() map[string]interface{} {
	now := float64(time.Now().Unix())

	grants := make(map[string]interface{})
	for _, grant := range token.Grants {
		grants[grant.Key()] = grant.ToPayload()
	}

	payload := map[string]interface{}{
		"jti":    fmt.Sprintf("%s-%s", token.SigningKeySid, strconv.Itoa(int(now))),
		"grants": grants,
	}

	if token.Identity != "" {
		val := payload["grants"].(map[string]interface{})
		val["identity"] = token.Identity
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
	if token.baseJwt.Subject != "" {
		payload["sub"] = token.baseJwt.Subject
	}

	return payload
}

// Encode this JWT struct into a string.
// algorithm - algorithm used to encode the JWT that overrides the default
// ttl - specify ttl to override the default
func (token *AccessToken) ToJwt() (string, error) {
	signedToken, err := token.baseJwt.ToJwt(token.generateHeaders, token.GeneratePayload)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func decodeGrants(grants interface{}) []BaseGrant {
	var decodedGrants []BaseGrant

	for k, v := range grants.(map[string]interface{}) {
		var grant BaseGrant
		if data, err := json.Marshal(v); err == nil {
			switch k {
			case "chat":
				grant = &ChatGrant{}
			case "rtc":
				grant = &ConversationsGrant{}
			case "ip_messaging":
				grant = &IpMessagingGrant{}
			case "data_sync":
				grant = &SyncGrant{}
			case "task_router":
				grant = &TaskRouterGrant{}
			case "video":
				grant = &VideoGrant{}
			case "voice":
				grant = &VoiceGrant{}
			case "player":
				grant = &PlaybackGrant{}
			}

			if errJson := json.Unmarshal(data, &grant); errJson == nil {
				decodedGrants = append(decodedGrants, grant)
			}
		}
	}

	return decodedGrants
}

// Decode a JWT string into a Jwt struct.
// jwt - JWT string
// key - string key used to verify the JWT signature; if not provided, then validation is skipped
func (token *AccessToken) FromJwt(jwtStr string, key string) (*AccessToken, error) {
	baseToken, err := token.baseJwt.FromJwt(jwtStr, key)
	if err != nil {
		return nil, err
	}

	decodedToken := &AccessToken{
		baseJwt:       baseToken,
		Grants:        decodeGrants(baseToken.Payload()["grants"]),
		AccountSid:    baseToken.Payload()["sub"].(string),
		SigningKeySid: baseToken.Payload()["iss"].(string),
	}

	if val, ok := baseToken.Headers()["twr"]; ok {
		decodedToken.Region = val
	}
	if val, ok := baseToken.Payload()["grants"]; ok {
		if iVal, iOk := val.(map[string]interface{})["identity"]; iOk {
			decodedToken.Identity = iVal.(string)
		}
	}

	return decodedToken, nil
}
