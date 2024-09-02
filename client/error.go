// Package error provides the interface for Twilio specific errors.
package client

import (
	"fmt"
)

// SendgridRestError provides information about an unsuccessful request.
type ErrorField struct {
	Field   interface{} `json:"field"` // null can be captured as interface{}
	Message string      `json:"message"`
}

type SendgridRestError struct {
	Errors []ErrorField `json:"errors"`
}

func (err *SendgridRestError) Error() string {
	return fmt.Sprintf("Error Field: %v - Error Message: %s",
		err.Errors[0].Field, err.Errors[0].Message)
}
