package domain

import (
	"net/http"
)

// ResponseError implements Error structure to return in response payloads.
type ResponseError struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Details    []string    `json:"details"`
	StackTrace interface{} `json:"stack_trace"`
}

// Error returns error message.
func (r *ResponseError) Error() string {
	return r.Message
}

// NewBadRequestError implements bad request error structure for our application.
func NewBadRequestError(err error, details ...string) *ResponseError {
	return &ResponseError{
		Code:       http.StatusBadRequest,
		Message:    "Bad Request",
		Details:    details,
		StackTrace: err,
	}
}
