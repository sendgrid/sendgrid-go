package securewebhook

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"math/big"
)

// Settings ...
type Settings struct {
	Enable *bool `json:"enabled,omitempty"`
}

// RS ...
type RS struct {
	R *big.Int
	S *big.Int
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

// VerifySignature ...
func VerifySignature(publicKey *ecdsa.PublicKey, payload []byte, signature, timestamp string) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	ecdsaSig := &RS{}
	_, err = asn1.Unmarshal(signatureBytes, ecdsaSig)
	if err != nil {
		return false, err
	}

	h := sha256.New()
	if timestamp != "" {
		h.Write([]byte(timestamp))
	}
	h.Write(payload)

	return ecdsa.Verify(publicKey, h.Sum(nil), ecdsaSig.R, ecdsaSig.S), nil
}
