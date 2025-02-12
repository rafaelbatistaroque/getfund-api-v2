package auth_service_spy

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/shared/result_app"
)

type AuthServiceSpy struct {
	Params map[string]string

	CallsCount int

	SuccessResult *auth_dto.SessionDto
	ErrorResult   *result_app.ApplicationError
}

func New() *AuthServiceSpy {
	return &AuthServiceSpy{Params: make(map[string]string), CallsCount: 0, ErrorResult: nil, SuccessResult: nil}
}

func (a *AuthServiceSpy) Authenticate(username string, password string) (*auth_dto.SessionDto, *result_app.ApplicationError) {
	a.Params["username"] = username
	a.Params["password"] = password

	a.CallsCount++

	return a.SuccessResult, a.ErrorResult
}

func (a *AuthServiceSpy) DefineNotAuthenticate(code int, message error) {
	a.ErrorResult = result_app.New(code, message)
}

func (a *AuthServiceSpy) DefineAuthenticate() {
	a.SuccessResult = &auth_dto.SessionDto{ID: 1, FirstName: "fake-first-name", IsAdmin: false}
}
