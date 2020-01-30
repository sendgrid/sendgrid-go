package main

import (
	"io/ioutil"
	"mime/multipart"
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
	boundary, body := getBoundary(string(file), "Content-Type: multipart/alternative; ")
	raw := multipart.NewReader(body, boundary)
	next, _ := raw.NextPart()
	value, _ := ioutil.ReadAll(next)
	assert.Equal(t, "001a113ee97c89842f0539be8e7a", boundary, "The boundary was not found.")
	assert.Equal(t, "Hello Twilio SendGrid!\n", string(value), "The email was not parsed properly.")
}
