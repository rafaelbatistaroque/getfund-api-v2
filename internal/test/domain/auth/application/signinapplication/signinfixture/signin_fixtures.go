package signinapplicationfixtures

import (
	"errors"
	"fmt"
	sut "getfund-api-v2/internal/domain/auth/application/signinapplication"
	entity "getfund-api-v2/internal/domain/auth/entity/sessionentity"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	appErr "getfund-api-v2/internal/pkg/applicationerror"
	appCode "getfund-api-v2/internal/pkg/helpers/applicationcode"
	validation "getfund-api-v2/internal/pkg/inputvalidation"
)

func NewSut() (signin.UseCase, *authServiceSpy, *sessionServiceSpy, *signinMapperSpy) {
	mapperSpy := &signinMapperSpy{Params: make(map[string]interface{}), ForceReturn: true}
	authServiceSpy := &authServiceSpy{Params: make(map[string]string), CallsCount: 0}
	sessionServiceSpy := &sessionServiceSpy{
		SaveSessionCallsCount: 0,
		BuildTokenCallsCount:  0,
		SaveSessionParam:      nil,
	}

	return sut.NewUseCase(authServiceSpy, sessionServiceSpy, mapperSpy), authServiceSpy, sessionServiceSpy, mapperSpy
}

func GetValidInput() *signin.Input {
	return &signin.Input{Password: "fake-password", UserName: "fake-username"}
}

func GetInputWithUserNameInvalid() (*signin.Input, *appErr.ApplicationError) {
	return &signin.Input{UserName: "", Password: "fake-password"},
		appErr.New(appCode.BAD_REQUEST, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "UserName"))
}

func GetInputWithPasswordInvalid() (*signin.Input, *appErr.ApplicationError) {
	return &signin.Input{Password: "", UserName: "fake-username"},
		appErr.New(appCode.BAD_REQUEST, fmt.Errorf(validation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

type authServiceSpy struct {
	Params map[string]string

	CallsCount int

	SuccessResult entity.Session
	errorResult   *appErr.ApplicationError
}

func (a *authServiceSpy) Authenticate(username string, password string) (entity.Session, *appErr.ApplicationError) {
	a.Params["username"] = username
	a.Params["password"] = password

	a.CallsCount++

	return a.SuccessResult, a.errorResult
}

func (a *authServiceSpy) DefineNotAuthenticate(code int, message error) {
	a.errorResult = appErr.New(code, message)
}

func (a *authServiceSpy) DefineAuthenticate() {
	a.SuccessResult, _ = entity.New("fake-id", "fake-first-name", 0)
}

type sessionServiceSpy struct {
	SaveSessionParam       entity.Session
	SaveSessionCallsCount  int
	saveSessionErrorResult error

	BuildTokenParam      entity.Session
	BuildTokenCallsCount int
	getTokenErrorResult  error
}

func (s *sessionServiceSpy) SaveSession(session entity.Session) error {
	s.SaveSessionParam = session

	s.SaveSessionCallsCount++

	return s.saveSessionErrorResult
}

func (s *sessionServiceSpy) BuildToken(session entity.Session) error {
	s.BuildTokenParam = session

	s.BuildTokenCallsCount++

	return s.getTokenErrorResult
}

func (s *sessionServiceSpy) DefineSaveSessionError() {
	s.saveSessionErrorResult = errors.New("any-error")
}

func (s *sessionServiceSpy) DefineBuildTokenError() {
	s.getTokenErrorResult = errors.New("any-error")
}

func (s *sessionServiceSpy) DefineBuildTokenSuccess(session entity.Session) {
	session.SetToken("fake-token")
}

type signinMapperSpy struct {
	Params      map[string]interface{}
	ForceReturn bool

	CallsCount int

	SuccessResult *signin.Output
}

func (m *signinMapperSpy) ToOutput(session entity.Session) *signin.Output {
	m.Params["session"] = session

	m.CallsCount++

	if m.ForceReturn {
		return nil
	}

	m.SuccessResult = &signin.SigninOutput{
		Token: session.GetToken(),
		Session: signin.SessionOutput{
			Id:        session.GetID(),
			FirstName: session.GetFirstName(),
			IsAdmin:   session.GetIsAdmin(),
		},
	}

	return m.SuccessResult
}
