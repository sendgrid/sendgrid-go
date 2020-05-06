package securewebhook

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecureWebhookNewSettings(t *testing.T) {
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")
}

func TestSecureWebhookSetEnable(t *testing.T) {
	s := NewSettings()
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")

	s.SetEnable(true)
	assert.Equal(t, true, *s.Enable, fmt.Sprintf("SecureWebhook.Enable should be 'true', got %v", *s.Enable))

	s.SetEnable(false)
	assert.Equal(t, false, *s.Enable, fmt.Sprintf("SecureWebhook.Enable should be 'false', got %v", *s.Enable))
}

func TestSecureWebhookGetRequestBody(t *testing.T) {
	expectedJSONEnabled := []byte(`{"enabled":true}`)
	expectedJSONDisabled := []byte(`{"enabled":false}`)

	s := NewSettings()
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")

	s.SetEnable(false)
	actualJSON, err := GetRequestBody(s)
	require.NoError(t, err)
	assert.Equal(t, expectedJSONDisabled, actualJSON, fmt.Sprintf("SecureWebhook.Enable should be '%b', got %b", expectedJSONDisabled, actualJSON))

	s.SetEnable(true)
	actualJSON, err = GetRequestBody(s)
	require.NoError(t, err)
	assert.Equal(t, expectedJSONEnabled, actualJSON, fmt.Sprintf("SecureWebhook.Enable should be '%b', got %b", expectedJSONEnabled, actualJSON))
}

func TestVerifySignature(t *testing.T) {
	priKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	payload, err := json.Marshal(map[string]interface{}{
		"event":      "test_event",
		"category":   "example_payload",
		"message_id": "message_id",
	})
	require.NoError(t, err)

	h := sha256.New()
	h.Write([]byte(timestamp))
	h.Write(payload)

	r, s, err := ecdsa.Sign(rand.Reader, priKey, h.Sum(nil))
	require.NoError(t, err)
	sigBytes, err := asn1.Marshal(RS{
		R: r,
		S: s,
	})
	require.NoError(t, err)
	signature := base64.StdEncoding.EncodeToString(sigBytes)

	// verifications
	verified, err := VerifySignature(&priKey.PublicKey, payload, signature, timestamp)
	require.NoError(t, err)
	assert.True(t, verified)

	// not valid payload
	verified, err = VerifySignature(&priKey.PublicKey, []byte("this is not valid payload for the given sign"), signature, timestamp)
	require.NoError(t, err)
	assert.False(t, verified)

	// not valid signature
	verified, err = VerifySignature(&priKey.PublicKey, payload, signature+"causing failure", timestamp)
	require.Error(t, err)
	assert.False(t, verified)

	// not valid timestamp
	verified, err = VerifySignature(&priKey.PublicKey, payload, signature, "invalid timestamp")
	require.NoError(t, err)
	assert.False(t, verified)

	// empty timestamp
	verified, err = VerifySignature(&priKey.PublicKey, payload, signature, "")
	require.NoError(t, err)
	assert.False(t, verified)
}
