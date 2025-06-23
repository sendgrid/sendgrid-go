package eventwebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	publicKey = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE83T4O/n84iotIvIW4mdBgQ/7dAfSmpqIM8kF9mN1flpVKS3GRqe62gw+2fNNRaINXvVpiglSI8eNEc6wEA3F+g=="
	signature = "MEUCIGHQVtGj+Y3LkG9fLcxf3qfI10QysgDWmMOVmxG0u6ZUAiEAyBiXDWzM+uOe5W0JuG+luQAbPIqHh89M15TluLtEZtM="
	timestamp = "1600112502"
)

func generateTestPayload() []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode([]map[string]interface{}{
		{
			"email":         "hello@world.com",
			"event":         "dropped",
			"reason":        "Bounced Address",
			"sg_event_id":   "ZHJvcC0xMDk5NDkxOS1MUnpYbF9OSFN0T0doUTRrb2ZTbV9BLTA",
			"sg_message_id": "LRzXl_NHStOGhQ4kofSm_A.filterdrecv-p3mdw1-756b745b58-kmzbl-18-5F5FC76C-9.0",
			"smtp-id":       "<LRzXl_NHStOGhQ4kofSm_A@ismtpd0039p1iad1.sendgrid.net>",
			"timestamp":     1600112492,
		},
	})
	payload := buffer.Bytes()
	payload = payload[:len(payload)-1]           // Drop the trailing newline the encoder adds.
	payload = append(payload, []byte("\r\n")...) // Append the expected trailing carriage return and newline!
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
	ecdsaKey, err := ConvertPublicKeyBase64ToECDSA(publicKey)
	require.NoError(t, err)
	assert.NotNil(t, ecdsaKey, "publicKey shouldn't be nil")

	ecdsaKey, err = ConvertPublicKeyBase64ToECDSA(publicKey + "corrupting the public key")
	require.Error(t, err)
	assert.Nil(t, ecdsaKey, "publicKey should be nil")
}

func TestVerifySignature(t *testing.T) {
	ecdsaKey, err := ConvertPublicKeyBase64ToECDSA(publicKey)
	require.NoError(t, err)

	payload := generateTestPayload()

	// verifications
	verified, err := VerifySignature(ecdsaKey, payload, signature, timestamp)
	require.NoError(t, err)
	assert.True(t, verified)

	// not valid payload
	verified, err = VerifySignature(ecdsaKey, []byte("this is not valid payload for the given signature"), signature, timestamp)
	require.NoError(t, err)
	assert.False(t, verified)

	// not valid signature
	verified, err = VerifySignature(ecdsaKey, payload, signature+"causing failure", timestamp)
	require.Error(t, err)
	assert.False(t, verified)

	// not valid timestamp
	verified, err = VerifySignature(ecdsaKey, payload, signature, "invalid timestamp")
	require.NoError(t, err)
	assert.False(t, verified)

	// empty timestamp
	verified, err = VerifySignature(ecdsaKey, payload, signature, "")
	require.NoError(t, err)
	assert.False(t, verified)
}
