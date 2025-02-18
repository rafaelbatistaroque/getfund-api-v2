package signout_fixture

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/signout/application"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/test/helper/session_spy"

	"github.com/rafaelbatistaroque/validation"
)

func NewSut() (signout.UseCase, *session_spy.SessionServiceSpy) {
	sessionServiceSpy := session_spy.New()

	return sut.New(sessionServiceSpy), sessionServiceSpy
}

func GetInvalidInputWithError() (*signout.Input, *result_app.ApplicationError) {
	return &signout.Input{Token: ""},
		result_app.New(result_app.UNAUTHORIZED_CODE, fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Token"))
}

func GetValidInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}
