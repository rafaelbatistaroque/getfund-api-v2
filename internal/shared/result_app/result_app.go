package result_app

import "fmt"

var (
	SUCCESS_CODE             = 0
	UNAUTHORIZED_CODE        = 1
	DUPLICATED_ENTRY_CODE    = 2
	NOT_FOUND_CODE           = 3
	SERVER_ERROR_CODE        = 4
	CONSTRAINT_VIOLATED_CODE = 5
	BAD_REQUEST_CODE         = 6
	UNAVAILABLE_CODE         = 7
	UNMODIFIED_CODE          = 8
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
