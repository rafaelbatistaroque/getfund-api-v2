package signoutapplicationfixture

import (
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	sut "getfund-api-v2/internal/domain/auth/usecase/signout/application"
	"getfund-api-v2/test/helper/sessionspy"
)

func NewSut() (signout.UseCase, *sessionspy.SessionServiceSpy) {
	sessionServiceSpy := sessionspy.New()

	return sut.New(sessionServiceSpy), sessionServiceSpy
}

func GetInvalidInput() *signout.Input {
	return &signout.Input{Token: ""}
}

func GetValidInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}
