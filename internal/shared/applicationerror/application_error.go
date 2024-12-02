package applicationerror

import (
	"fmt"
)

type ApplicationError struct {
	Code    int
	Message error
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func New(code int, message error) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Message: message,
	}
}
