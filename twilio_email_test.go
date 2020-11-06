package sendgrid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTwilioEmailSendClient(t *testing.T) {
	mailClient := NewTwilioEmailSendClient("username", "password")
	assert.Equal(t, "https://email.twilio.com/v3/mail/send", mailClient.BaseURL)
	assert.Equal(t, "Basic dXNlcm5hbWU6cGFzc3dvcmQ=", mailClient.Headers["Authorization"])
}

func TestGetTwilioEmailRequest(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer fakeServer.Close()

	request := GetTwilioEmailRequest(TwilioEmailOptions{
		Username: "username",
		Password: "password",
		Host:     fakeServer.URL,
	})
	response, err := MakeRequest(request)
	require.NoError(t, err)
	assert.True(t, strings.Contains(response.Body, "success"))
}
