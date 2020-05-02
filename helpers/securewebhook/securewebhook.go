package securewebhook

import (
	"encoding/json"
	"fmt"
)

// Settings ...
type Settings struct {
	Enable bool `json:"enable,omitempty"`
}

// NewSettings ...
func NewSettings() *Settings {
	return &Settings{}
}

// SetEnable ...
func (s *Settings) SetEnable(enable bool) *Settings {
	return &Settings{Enable: enable}
}

// GetRequestBody ...
func GetRequestBody(s *Settings) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
