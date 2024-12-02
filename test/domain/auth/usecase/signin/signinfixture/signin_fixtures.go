package signinapplicationfixtures

import (
	"errors"
	"fmt"
	authmodel "getfund-api-v2/internal/domain/auth/model"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	sut "getfund-api-v2/internal/domain/auth/usecase/signin/application"
	validation "getfund-api-v2/internal/pkg/inputvalidation"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/test/spyshared/mapperspy/signinmapperspy"
)

func NewSut() (signin.UseCase, *authServiceSpy, *sessionServiceSpy, *signinmapperspy.SigninMapperSpy) {
	mapperSpy := signinmapperspy.New()
	authServiceSpy := &authServiceSpy{Params: make(map[string]string), CallsCount: 0}
	sessionServiceSpy := &sessionServiceSpy{CallsCount: make(map[string]int), Params: make(map[string]string), SuccessResult: make(map[string]interface{}), ErrorResult: make(map[string]error)}

	return sut.New(authServiceSpy, sessionServiceSpy, mapperSpy), authServiceSpy, sessionServiceSpy, mapperSpy
}

func GetValidInput() *signin.Input {
	return &signin.Input{Password: "fake-password", UserName: "fake-username"}
}

func GetInputWithUserNameInvalid() (*signin.Input, *resultapp.ApplicationError) {
	return &signin.Input{UserName: "", Password: "fake-password"},
		resultapp.New(resultapp.BAD_REQUEST, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "UserName"))
}

func GetInputWithPasswordInvalid() (*signin.Input, *resultapp.ApplicationError) {
	return &signin.Input{Password: "", UserName: "fake-username"},
		resultapp.New(resultapp.BAD_REQUEST, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Password"))
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

type sessionServiceSpy struct {
	Params     map[string]string
	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func (s *sessionServiceSpy) SaveSession(session string) (string, error) {
	s.Params["SaveSession:session"] = session

	s.CallsCount["SaveSession"]++

	success := s.SuccessResult["SaveSession"]
	if success != nil {
		return s.SuccessResult["SaveSession"].(string), s.ErrorResult["SaveSession"]
	}

	return "", s.ErrorResult["SaveSession"]
}

func (s *sessionServiceSpy) DeleteSession(session string) error {
	s.Params["DeleteSession:session"] = session

	s.CallsCount["DeleteSession"]++

	return s.ErrorResult["DeleteSession"]
}

func (s *sessionServiceSpy) GetSession(session string) (string, error) {
	return "", nil
}

func (s *sessionServiceSpy) DefineError() {
	s.ErrorResult["SaveSession"] = errors.New("any-error")
}

func (s *sessionServiceSpy) DefineSuccess() {
	s.SuccessResult["SaveSession"] = "fake-success"
}
