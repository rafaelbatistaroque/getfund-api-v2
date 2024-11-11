package inputvalidation

import "errors"

var (
	Err_Msg_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO = errors.New("parameter %s should be greather than zero")
	Err_Msg_PARAMETER_NOT_EMPTY                    = errors.New("parameter %s can`t be null or empty")
)
