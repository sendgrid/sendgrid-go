package main

import (
	"io/ioutil"
	"mime/multipart"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := loadConfig("./conf.json")
	if conf.Endpoint == "" {
		t.Errorf("conf.json did not load correctly, Endpoint empty")
	}
	if conf.Port == "" {
		t.Errorf("conf.json did not load correctly, Port empty")
	}
}

func TestGetBoundary(t *testing.T) {
	file, _ := ioutil.ReadFile("./sample_data/raw_data.txt")
	boundary, body := getBoundary(string(file), "Content-Type: multipart/alternative; ")
	raw := multipart.NewReader(body, boundary)
	next, _ := raw.NextPart()
	value, _ := ioutil.ReadAll(next)
	if boundary != "001a113ee97c89842f0539be8e7a" {
		t.Errorf("The boundary was not found.")
	}
	if string(value) != "Hello SendGrid!\n" {
		t.Errorf("The email was not parsed properly.")
	}
}
