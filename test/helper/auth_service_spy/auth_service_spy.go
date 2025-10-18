package auth_service_spy

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	shared_error "getfund-api-v2/internal/shared/error"
)

type AuthServiceSpy struct {
	Params map[string]string

	CallsCount int

	SuccessResult *auth_dto.SessionDto
	ErrorResult   *shared_error.Error
}

func New() *AuthServiceSpy {
	return &AuthServiceSpy{Params: make(map[string]string), CallsCount: 0, ErrorResult: nil, SuccessResult: nil}
}

func (a *AuthServiceSpy) Authenticate(username string, password string) (*auth_dto.SessionDto, *shared_error.Error) {
	a.Params["username"] = username
	a.Params["password"] = password

	a.CallsCount++

	return a.SuccessResult, a.ErrorResult
}

func (a *AuthServiceSpy) DefineNotAuthenticate(code int, message error) {
	a.ErrorResult = shared_error.New(code, message)
}

func (a *AuthServiceSpy) DefineAuthenticate() {
	a.SuccessResult = &auth_dto.SessionDto{ID: 1, FirstName: "fake-first-name", IsAdmin: false}
}
