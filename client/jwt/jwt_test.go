package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/twilio/twilio-go/client/jwt/util"
)

func TestGetHeaders(t *testing.T) {
	jwtT := Jwt{
		SecretKey: "secret",
		Issuer:    "twilio",
		Subject:   "twilio jwt",
		Algorithm: "HS256",
	}
	headers := jwtT.Headers()
	assert.Equal(t, HS256, headers["alg"])
	assert.Equal(t, JWT, headers["typ"])
}

func TestGetPayload(t *testing.T) {
	jwtT := Jwt{
		SecretKey:      "secret",
		Issuer:         "twilio",
		Subject:        "twilio jwt",
		Algorithm:      "HS256",
		Nbf:            0,
		Ttl:            0,
		ValidUntil:     0,
		DecodedHeaders: nil,
		DecodedPayload: nil,
	}

	payload := map[string]interface{}{}
	payload = jwtT.generatePayload(payload)
	assert.Equal(t, "twilio", payload["iss"])
	assert.NotZero(t, payload["exp"])
	assert.NotZero(t, payload["nbf"])
	assert.Equal(t, "twilio jwt", payload["sub"])
}
