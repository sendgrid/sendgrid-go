package sendgrid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	APIUser     = "API_USER"
	APIPassword = "API_PASSWORD"
)

func TestSendGridVersion(t *testing.T) {
	if Version != "2.0.0" {
		t.Error("SendGrid version does not match")
	}
}

func TestNewSendGridClient(t *testing.T) {
	client := NewSendGridClient(APIUser, APIPassword)
	if client == nil {
		t.Error("NewSendGridClient should never return nil")
	}
}

func TestNewSendGridClientUsernamePassword(t *testing.T) {
	client := NewSendGridClient(APIUser, APIPassword)
	if client.apiUser == "" || client.apiPwd == "" {
		t.Error("NewSendGridClient should have username and password set")
	}
}

func TestNewSendGridClientApiKey(t *testing.T) {
	client := NewSendGridClientWithApiKey(APIPassword)
	if client.apiUser != "" && client.apiPwd == APIPassword {
		t.Error("NewSendGridClient should have api ket set")
	}
}

func TestSend(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer fakeServer.Close()
	m := NewMail()
	client := NewSendGridClient(APIUser, APIPassword)
	client.APIMail = fakeServer.URL
	m.AddTo("Test! <test@email.com>")
	m.SetSubject("Test")
	m.SetText("Text")

	if e := client.Send(m); e != nil {
		t.Errorf("Send failed to send email. Returned error: %v", e)
	}
}

func TestSendForAuthorizationHeader(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+APIPassword {
			t.Error("Send failed to have authorization header")
		}
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer fakeServer.Close()
	m := NewMail()
	client := NewSendGridClientWithApiKey(APIPassword)
	client.APIMail = fakeServer.URL
	m.AddTo("Test! <test@email.com>")
	m.SetSubject("Test")
	m.SetText("Text")

	if e := client.Send(m); e != nil {
		t.Errorf("Send failed to send email. Returned error: %v", e)
	}
}
