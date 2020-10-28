package mock

import "github.com/sendgrid/rest"

// Mock - Mock struct, has Response Code, Response Body and errors
type Mock struct {
	StatusCode int
	Body       string
	Err        error
}

var mock *Mock

// Add - add mock method
func Add(m *Mock) {
	mock = m
}

// Flush - Flush mock method
func Flush() {
	mock = nil
}

// Get - Get mock method
func Get() *Mock {
	return mock
}

// Request - return mock sengrid response and error
func Request() (*rest.Response, error) {

	if mock == nil {
		panic("There isn't a mock")
	}

	if mock.Err != nil {
		return nil, mock.Err
	}

	r := &rest.Response{
		StatusCode: mock.StatusCode,
		Body:       mock.Body,
	}

	return r, nil
}
