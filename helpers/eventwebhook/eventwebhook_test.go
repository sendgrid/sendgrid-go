package eventwebhook

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	timestamp      = "1588788367"
	testPublicKey  = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEDr2LjtURuePQzplybdC+u4CwrqDqBaWjcMMsTbhdbcwHBcepxo7yAQGhHPTnlvFYPAZFceEu/1FwCM/QmGUhA=="
	testPrivateKey = "MHcCAQEEIEQmZgBEh9DcU9zMl34czK3xov4AYvm9P3r7rNB2dNXtoAoGCCqGSM49AwEHoUQDQgAEEDr2LjtURuePQzplybdC+u4CwrqDqBaWjcMMsTbhdbcwHBcepxo7yAQGhHPTnlvFYPAZFceEu/1FwCM/QmGUhA=="
	signature      = "MEUCIQCtIHJeH93Y+qpYeWrySphQgpNGNr/U+UyUlBkU6n7RAwIgJTz2C+8a8xonZGi6BpSzoQsbVRamr2nlxFDWYNH2j/0="
)

func generateTestPayload() []byte {
	payload, _ := json.Marshal(map[string]interface{}{
		"event":      "test_event",
		"category":   "example_payload",
		"message_id": "message_id",
	})
	return payload
}

func TestEventWebhookNewSettings(t *testing.T) {
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")
}

func TestSetEnableSignedWebhook(t *testing.T) {
	s := NewSettings()
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")

	s.SetEnableSignedWebhook(true)
	assert.Equal(t, true, *s.EnableSignedWebhook, fmt.Sprintf("EnableSignedWebhook should be 'true', got %v", *s.EnableSignedWebhook))

	s.SetEnableSignedWebhook(false)
	assert.Equal(t, false, *s.EnableSignedWebhook, fmt.Sprintf("EnableSignedWebhook should be 'false', got %v", *s.EnableSignedWebhook))
}

func TestSignedWebhookGetRequestBody(t *testing.T) {
	expectedJSONEnabled := []byte(`{"enabled":true}`)
	expectedJSONDisabled := []byte(`{"enabled":false}`)

	s := NewSettings()
	assert.NotNil(t, NewSettings(), "NewSettings() shouldn't return nil")

	s.SetEnableSignedWebhook(false)
	actualJSON, err := GetRequestBody(s)
	require.NoError(t, err)
	assert.Equal(t, expectedJSONDisabled, actualJSON, fmt.Sprintf("EnableSignedWebhook should be '%b', got %b", expectedJSONDisabled, actualJSON))

	s.SetEnableSignedWebhook(true)
	actualJSON, err = GetRequestBody(s)
	require.NoError(t, err)
	assert.Equal(t, expectedJSONEnabled, actualJSON, fmt.Sprintf("EnableSignedWebhook should be '%b', got %b", expectedJSONEnabled, actualJSON))
}

func TestConvertPublicKeyBase64ToECDSA(t *testing.T) {
	publicKey, err := ConvertPublicKeyBase64ToECDSA(testPublicKey)
	require.NoError(t, err)
	assert.NotNil(t, publicKey, "publicKey shouldn't be nil")

	publicKey, err = ConvertPublicKeyBase64ToECDSA(testPublicKey + "corrupting the public key")
	require.Error(t, err)
	assert.Nil(t, publicKey, "publicKey should be nil")
}

func TestVerifySignature(t *testing.T) {
	publicKey, err := ConvertPublicKeyBase64ToECDSA(testPublicKey)
	require.NoError(t, err)

	payload := generateTestPayload()

	// verifications
	verified, err := VerifySignature(publicKey, payload, signature, timestamp)
	require.NoError(t, err)
	assert.True(t, verified)

	// not valid payload
	verified, err = VerifySignature(publicKey, []byte("this is not valid payload for the given signature"), signature, timestamp)
	require.NoError(t, err)
	assert.False(t, verified)

	// not valid signature
	verified, err = VerifySignature(publicKey, payload, signature+"causing failure", timestamp)
	require.Error(t, err)
	assert.False(t, verified)

	// not valid timestamp
	verified, err = VerifySignature(publicKey, payload, signature, "invalid timestamp")
	require.NoError(t, err)
	assert.False(t, verified)

	// empty timestamp
	verified, err = VerifySignature(publicKey, payload, signature, "")
	require.NoError(t, err)
	assert.False(t, verified)
}
