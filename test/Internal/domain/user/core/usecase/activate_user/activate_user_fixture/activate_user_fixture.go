package activate_user_fixture

import (
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	activate_user_application "getfund-api-v2/internal/domain/user/core/usecase/activate_user/application"
)

type ActivateUserFixture struct {
}

func NewSut() (activate_user.UseCase, *ActivateUserFixture) {

	return activate_user_application.New(),
		&ActivateUserFixture{}
}

type Option func(*activate_user.Input)

func GetInput(options ...Option) *activate_user.Input {
	input := &activate_user.Input{
		ActivationCode: "fakeActivationCodeVl",
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyActivationCode() Option {
	return func(params *activate_user.Input) {
		params.ActivationCode = ""
	}
}

func WithInvalidActivationCodeLength() Option {
	return func(params *activate_user.Input) {
		params.ActivationCode = "invalid"
	}
}
