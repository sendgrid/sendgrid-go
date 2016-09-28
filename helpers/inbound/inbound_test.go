package main

import "testing"

func TestLoadConfig(t *testing.T) {
	conf := LoadConfig("./conf.json")
	if conf.Endpoint == "" {
		t.Errorf("conf.json did not load correctly, Endpoint empty")
	}
	if conf.Port == "" {
		t.Errorf("conf.json did not load correctly, Port empty")
	}
}
