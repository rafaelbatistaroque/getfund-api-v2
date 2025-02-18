package signin_fixture

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/signin/application"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/test/helper/auth_service_spy"
	"getfund-api-v2/test/helper/mapper_spy/signin_mapper_spy"
	"getfund-api-v2/test/helper/session_spy"

	validation "github.com/rafaelbatistaroque/validation"
)

type SigninFixture struct {
	AuthServiceSpy *auth_service_spy.AuthServiceSpy
	SessionSpy     *session_spy.SessionServiceSpy
	MapperSpy      *signin_mapper_spy.SigninMapperSpy
}

func NewSut() (signin.UseCase, *SigninFixture) {
	authServiceSpy := auth_service_spy.New()
	sessionServiceSpy := session_spy.New()
	mapperSpy := signin_mapper_spy.New()

	return sut.New(authServiceSpy, sessionServiceSpy, mapperSpy),
		&SigninFixture{
			AuthServiceSpy: authServiceSpy,
			SessionSpy:     sessionServiceSpy,
			MapperSpy:      mapperSpy,
		}
}

func GetValidInput() *signin.Input {
	return &signin.Input{Password: "fake-password", UserName: "fake-username"}
}

func GetInputWithUserNameInvalid() (*signin.Input, *result_app.ApplicationError) {
	return &signin.Input{UserName: "", Password: "fake-password"},
		result_app.New(result_app.BAD_REQUEST_CODE, fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "UserName"))
}

func GetInputWithPasswordInvalid() (*signin.Input, *result_app.ApplicationError) {
	return &signin.Input{Password: "", UserName: "fake-username"},
		result_app.New(result_app.BAD_REQUEST_CODE, fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}
