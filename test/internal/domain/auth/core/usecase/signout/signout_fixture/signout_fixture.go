package signout_fixture

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/signout/application"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/session_spy"

	"github.com/rafaelbatistaroque/validation"
)

func NewSut() (signout.UseCase, *session_spy.SessionServiceSpy) {
	sessionServiceSpy := session_spy.New()

	return sut.New(sessionServiceSpy), sessionServiceSpy
}

func GetInvalidInputWithError() (*signout.Input, *shared_error.Error) {
	return &signout.Input{Token: ""},
		shared_error.New(shared_error.UNAUTHORIZED_CODE, fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Token"))
}

func GetValidInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}
