package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testURL   = "https://mycompany.com/myapp.php?foo=1&bar=2"
	signature = "vOEb5UThFn24KEfnOFLQY2AE5FY=" // of the testURL above with the params below
	bodyHash  = "0a1ff7634d9ab3b95db5c9a2dfe9416e41502b283a80c7cf19632632f96e6620"
)

var (
	validator = NewRequestValidator("12345")
	params    = map[string]string{
		"Digits":                "1234",
		"CallSid":               "CA1234567890ABCDE",
		"To":                    "+18005551212",
		"Caller":                "+14158675309",
		"From":                  "+14158675309",
		"ReasonConferenceEnded": "test",
		"Reason":                "Participant",
	}
	jsonBody = []byte(`{"property": "value", "boolean": true}`)
	formBody = []byte(`property=value&boolean=true`)
)

func TestRequestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("returns true when validation succeeds", func(t *testing.T) {
		assert.True(t, validator.Validate(testURL, params, signature))
	})

	t.Run("returns false when validation fails", func(t *testing.T) {
		assert.False(t, validator.Validate(testURL, params, "WRONG SIGNATURE"))
	})

	t.Run("returns true when https and port is specified but signature is generated without it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com:1234", 1)
		assert.True(t, validator.Validate(theURL, params, signature))
	})

	t.Run("returns true when https and port is specified and signature is generated with it", func(t *testing.T) {
		expectedSignature := "vOEb5UThFn24KEfnOFLQY2AE5FY=" // hash of https uri without port
		assert.True(t, validator.Validate(testURL, params, expectedSignature))
	})

	t.Run("returns true when http and port port is specified but signature is generated without it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com", 1)
		theURL = strings.Replace(theURL, "https", "http", 1)
		expectedSignature := "n2xBNyzSW7rfYStDtOFiFMv7qNo=" // hash of http uri without port
		assert.True(t, validator.Validate(theURL, params, expectedSignature))
	})

	t.Run("returns true when http and port is specified and signature is generated with it", func(t *testing.T) {
		theURL := strings.Replace(testURL, ".com", ".com:1234", 1)
		theURL = strings.Replace(theURL, "https", "http", 1)
		expectedSignature := "n2xBNyzSW7rfYStDtOFiFMv7qNo=" // hash of http uri with port 1234
		assert.True(t, validator.Validate(theURL, params, expectedSignature))
	})

	t.Run("return false when params are sorted incorrectly", func(t *testing.T) {
		incorrectSignature := "95+Bu0JVPi0r/SsESZCVf0dWAjw=" //Params ReasonConferenceEnded is sorted before Reason
		assert.False(t, validator.Validate(testURL, params, incorrectSignature))
	})
}

func TestRequestValidator_ValidateBody(t *testing.T) {
	t.Parallel()

	t.Run("returns true when validation succeeds with json body", func(t *testing.T) {
		theURL := testURL + "&bodySHA256=" + bodyHash
		signatureWithBodyHash := "a9nBmqA0ju/hNViExpshrM61xv4="
		assert.True(t, validator.ValidateBody(theURL, jsonBody, signatureWithBodyHash))
	})

	t.Run("returns true when validation succeeds with form body", func(t *testing.T) {
		expectedSignature := "NBdBDr/T/lgjI+tlgpXjKZQZs/k="
		assert.True(t, validator.ValidateBody(testURL, formBody, expectedSignature))
	})

	t.Run("returns false when validation fails with json body", func(t *testing.T) {
		assert.False(t, validator.ValidateBody(testURL, jsonBody, signature))
	})

	t.Run("returns true when there's no other parameters and the signature is right", func(t *testing.T) {
		theURL := "https://mycompany.com/myapp.php?bodySHA256=" + bodyHash
		signatureForURL := "y77kIzt2vzLz71DgmJGsen2scGs="
		assert.True(t, validator.ValidateBody(theURL, jsonBody, signatureForURL))
	})
}
