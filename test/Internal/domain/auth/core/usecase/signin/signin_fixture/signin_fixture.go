package signin_fixture

import (
	"fmt"
	authmodel "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/signin/application"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/test/helper/mapper_spy/signin_mapper_spy"
	"getfund-api-v2/test/helper/session_spy"

	validation "github.com/rafaelbatistaroque/validation"
)

func NewSut() (signin.UseCase, *authServiceSpy, *session_spy.SessionServiceSpy, *signin_mapper_spy.SigninMapperSpy) {
	mapperSpy := signin_mapper_spy.New()
	authServiceSpy := &authServiceSpy{Params: make(map[string]string), CallsCount: 0}
	sessionServiceSpy := session_spy.New()

	return sut.New(authServiceSpy, sessionServiceSpy, mapperSpy), authServiceSpy, sessionServiceSpy, mapperSpy
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

type authServiceSpy struct {
	Params map[string]string

	CallsCount int

	SuccessResult *authmodel.SessionModel
	errorResult   *result_app.ApplicationError
}

func (a *authServiceSpy) Authenticate(username string, password string) (*authmodel.SessionModel, *result_app.ApplicationError) {
	a.Params["username"] = username
	a.Params["password"] = password

	a.CallsCount++

	return a.SuccessResult, a.errorResult
}

func (a *authServiceSpy) DefineNotAuthenticate(code int, message error) {
	a.errorResult = result_app.New(code, message)
}

func (a *authServiceSpy) DefineAuthenticate() {
	a.SuccessResult = &authmodel.SessionModel{ID: "fake-id", FirstName: "fake-first-name", IsAdmin: 0}
}
