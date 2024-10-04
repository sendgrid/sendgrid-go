package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/client"
)

func TestUnmarshalFloat32(t *testing.T) {
	unmarshalFloat32 := func(t *testing.T, input interface{}) float32 {
		value, err := client.UnmarshalFloat32(&input)
		assert.Nil(t, err)
		return *value
	}

	assert.Equal(t, float32(0), unmarshalFloat32(t, "0"))
	assert.Equal(t, float32(1), unmarshalFloat32(t, "1"))
	assert.Equal(t, float32(123), unmarshalFloat32(t, 123))
	assert.Equal(t, float32(123.456), unmarshalFloat32(t, "123.456"))
	assert.Equal(t, float32(123.456), unmarshalFloat32(t, 123.456))
	assert.Equal(t, float32(123.456), unmarshalFloat32(t, float32(123.456)))
	assert.Equal(t, float32(7), unmarshalFloat32(t, float64(7)))
}

func TestUnmarshalFloat32_Nil(t *testing.T) {
	value, err := client.UnmarshalFloat32(nil)
	assert.Nil(t, value)
	assert.Nil(t, err)
}

func TestUnmarshalFloat32_ErrorInvalid(t *testing.T) {
	var input interface{} = map[string]interface{}{}
	value, err := client.UnmarshalFloat32(&input)
	assert.Nil(t, value)
	assert.NotNil(t, err)
}

func TestUnmarshalFloat32_ErrorParse(t *testing.T) {
	var input interface{} = "ugh"
	value, err := client.UnmarshalFloat32(&input)
	assert.Nil(t, value)
	assert.NotNil(t, err)
}
