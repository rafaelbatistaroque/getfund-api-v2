package shared_error

import "fmt"

const (
	SUCCESS_CODE               = 0
	UNAUTHORIZED_CODE          = 1
	DUPLICATED_ENTRY_CODE      = 2
	NOT_FOUND_CODE             = 3
	SERVER_ERROR_CODE          = 4
	CONSTRAINT_VIOLATED_CODE   = 5
	BAD_REQUEST_CODE           = 6
	UNAVAILABLE_CODE           = 7
	UNMODIFIED_CODE            = 8
	SUCCESS_CREATED_CODE       = 9
	UNPROCESSABLE_CONTENT_CODE = 10
)

// Error is a custom error type that includes a code and a message.
// This allows for more specific error handling.
type Error struct {
	Code    int
	Message error
}

// Error returns the string representation of the error.
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// New creates a new Error instance.
func New(code int, message error) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
