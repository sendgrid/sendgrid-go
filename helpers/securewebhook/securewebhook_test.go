package securewebhook

import (
	"fmt"
	"testing"

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
