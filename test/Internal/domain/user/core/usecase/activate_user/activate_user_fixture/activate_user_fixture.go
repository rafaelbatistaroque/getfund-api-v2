package activate_user_fixture

import (
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	activate_user_application "getfund-api-v2/internal/domain/user/core/usecase/activate_user/application"
	"getfund-api-v2/test/helper/cache_spy"
)

type ActivateUserFixture struct {
	CacheSpy *cache_spy.RedisCacheSpy
}

func NewSut() (activate_user.UseCase, *ActivateUserFixture) {
	cacheSpy := cache_spy.New()

	return activate_user_application.New(cacheSpy),
		&ActivateUserFixture{
			CacheSpy: cacheSpy,
		}
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
