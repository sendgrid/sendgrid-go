package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	conf := loadConfig("./conf.json")
	assert.NotNil(t, conf.Endpoint, "conf.json did not load correctly, Endpoint empty")
	assert.NotNil(t, conf.Port, "conf.json did not load correctly, Port empty")
}

func TestGetBoundary(t *testing.T) {
	file, _ := ioutil.ReadFile("./sample_data/raw_data.txt")
	boundary, body := getBoundary(string(file), "multipart/alternative; boundary=xYzZY")
	assert.Equal(t, "xYzZY", boundary, "The boundary was not found.")
	raw := multipart.NewReader(body, boundary)
	next, _ := raw.NextPart()
	value := readBody(next)
	assert.Equal(t, "{@sendgrid.com : pass}", string(value), "The email was not parsed properly.")
}

func TestInboundHandler(t *testing.T) {
       // Build a table of tests to run with each one having a name, the sample data file to post,
      // and the expected HTTP response from the handler
	tests := []struct {
		name             string
		file             string
		expectedResponse int
	}{
		{"NoAttachment", "./sample_data/raw_data.txt", http.StatusOK},
		{"Attachment", "./sample_data/raw_data_with_attachments.txt", http.StatusOK},
	}

	for _, test := range tests {
		t.Run(test.name, func(subTest *testing.T) {
			//Load POST body
			file, err := ioutil.ReadFile(test.file)
			if err != nil {
				subTest.Fatal("could not load test file")
			}

			resp := httptest.NewRecorder()
			// Build POST request
			req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(file))
			req.Header.Set("Content-Type", "multipart/form-data; boundary=xYzZY")
			req.Header.Set("User-Agent", "Twilio-SendGrid-Test")

			// Invoke callback handler
			inboundHandler(resp, req)
			if resp.Code != test.expectedResponse {
				subTest.Errorf("Wrong HTTP response: got: %d, expected: %d", resp.Code, test.expectedResponse)
			}
		})
	}
}
