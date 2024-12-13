package signout_fixture

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signout"
	sut "getfund-api-v2/internal/domain/auth/adapter/usecase/signout/application"
	"getfund-api-v2/test/helper/session_spy"
)

func NewSut() (signout.UseCase, *session_spy.SessionServiceSpy) {
	sessionServiceSpy := session_spy.New()

	return sut.New(sessionServiceSpy), sessionServiceSpy
}

func GetInvalidInput() *signout.Input {
	return &signout.Input{Token: ""}
}

func GetValidInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}
