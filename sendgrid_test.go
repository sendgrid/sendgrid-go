package sendgrid

import "testing"

func TestSendGridVersion(t *testing.T) {
	if Version != "3.0.0" {
		t.Error("SendGrid version does not match")
	}
}

func TestGetRequest(t *testing.T) {
	request := GetRequest("", "", "", "")
	if request.BaseURL != "https://api.sendgrid.com/v3" {
		t.Error("Host default not set")
	}
	if request.RequestHeaders["Content-Type"] != "application/json" {
		t.Error("Wrong default content type")
	}
	if request.RequestHeaders["Authorization"] != "Bearer " {
		t.Error("Wrong default Authorization")
	}
	if request.RequestHeaders["User-Agent"] != "sendgrid/"+Version+";go" {
		t.Error("Wrong default User Agent")
	}

	request = GetRequest("API_KEY", "/endpoint", "https://test.api.com", "v4")
	if request.BaseURL != "https://test.api.com/v4/endpoint" {
		t.Error("Host not set correctly")
	}
	if request.RequestHeaders["Content-Type"] != "application/json" {
		t.Error("Wrong content type")
	}
	if request.RequestHeaders["Authorization"] != "Bearer API_KEY" {
		t.Error("Wrong Authorization")
	}
	if request.RequestHeaders["User-Agent"] != "sendgrid/"+Version+";go" {
		t.Error("Wrong User Agent")
	}
}
