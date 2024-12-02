package resultapp

import "fmt"

var (
	CODE_SUCCESS             = 0
	CODE_UNAUTHORIZED        = 1
	CODE_DUPLICATED_ENTRY    = 2
	CODE_NOT_FOUND           = 3
	CODE_SERVER_ERROR        = 4
	CODE_CONSTRAINT_VIOLATED = 5
	BAD_REQUEST              = 6
	CODE_UNAVAILABLE         = 7
	CODE_UNMODIFIED          = 8
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
