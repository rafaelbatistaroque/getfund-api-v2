package signinapplicationfixtures

import (
	"fmt"
	authmodel "getfund-api-v2/internal/domain/auth/model"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	sut "getfund-api-v2/internal/domain/auth/usecase/signin/application"
	validation "getfund-api-v2/internal/pkg/inputvalidation"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/test/spyshared/mapperspy/signinmapperspy"
	"getfund-api-v2/test/spyshared/sessionspy"
)

func NewSut() (signin.UseCase, *authServiceSpy, *sessionspy.SessionServiceSpy, *signinmapperspy.SigninMapperSpy) {
	mapperSpy := signinmapperspy.New()
	authServiceSpy := &authServiceSpy{Params: make(map[string]string), CallsCount: 0}
	sessionServiceSpy := sessionspy.New()

	return sut.New(authServiceSpy, sessionServiceSpy, mapperSpy), authServiceSpy, sessionServiceSpy, mapperSpy
}

func GetValidInput() *signin.Input {
	return &signin.Input{Password: "fake-password", UserName: "fake-username"}
}

func GetInputWithUserNameInvalid() (*signin.Input, *resultapp.ApplicationError) {
	return &signin.Input{UserName: "", Password: "fake-password"},
		resultapp.New(resultapp.BAD_REQUEST_CODE, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "UserName"))
}

func GetInputWithPasswordInvalid() (*signin.Input, *resultapp.ApplicationError) {
	return &signin.Input{Password: "", UserName: "fake-username"},
		resultapp.New(resultapp.BAD_REQUEST_CODE, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

type authServiceSpy struct {
	Params map[string]string

	CallsCount int

	SuccessResult *authmodel.SessionModel
	errorResult   *resultapp.ApplicationError
}

func (a *authServiceSpy) Authenticate(username string, password string) (*authmodel.SessionModel, *resultapp.ApplicationError) {
	a.Params["username"] = username
	a.Params["password"] = password

	a.CallsCount++

	return a.SuccessResult, a.errorResult
}

func (a *authServiceSpy) DefineNotAuthenticate(code int, message error) {
	a.errorResult = resultapp.New(code, message)
}

func (a *authServiceSpy) DefineAuthenticate() {
	a.SuccessResult = &authmodel.SessionModel{ID: "fake-id", FirstName: "fake-first-name", IsAdmin: 0}
}
