package securewebhook

import (
	"encoding/json"
)

// Settings ...
type Settings struct {
	Enable *bool `json:"enabled,omitempty"`
}

// NewSettings ...
func NewSettings() *Settings {
	return &Settings{}
}

// SetEnable ...
func (s *Settings) SetEnable(enable bool) {
	s.Enable = &enable
}

// GetRequestBody ...
func GetRequestBody(s *Settings) ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}
