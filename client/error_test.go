package client_test

import (
	"encoding/json"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/twilio/twilio-go/client"
)

const (
	errorCode     = 20001
	errorMessage  = "Bad request"
	errorMoreInfo = "https://www.twilio.com/docs/errors/20001"
	errorStatus   = 400
)

func TestTwilioRestError_Error(t *testing.T) {
	details := make(map[string]interface{})
	details["foo"] = "bar"
	err := &client.TwilioRestError{
		Code:     errorCode,
		Details:  details,
		Message:  errorMessage,
		MoreInfo: errorMoreInfo,
		Status:   errorStatus,
	}
	expected := "Status: 400 - ApiError 20001: Bad request ({\"foo\":\"bar\"}) More info: https://www.twilio.com/docs/errors/20001"
	assert.Equal(t, expected, err.Error())
}

func TestTwilioRestError_NoDetails(t *testing.T) {
	err := &client.TwilioRestError{}
	response := `{"code":20001,"message":"Bad request","more_info":"https://www.twilio.com/docs/errors/20001","status":400}`
	responseReader := strings.NewReader(response)
	decodeErr := json.NewDecoder(responseReader).Decode(err)
	assert.Nil(t, decodeErr)
	assert.Equal(t, err.Code, errorCode)
	assert.Equal(t, err.Status, errorStatus)
	assert.Equal(t, err.Message, errorMessage)
	assert.Equal(t, err.MoreInfo, errorMoreInfo)
	assert.Nil(t, err.Details)
}
