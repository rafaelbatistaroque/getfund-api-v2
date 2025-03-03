package create_user_fixture

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/create_user"
	create_user_application "getfund-api-v2/internal/domain/auth/core/usecase/create_user/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

const (
	_51_Chars = "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"
)

type CreateUserFixture struct {
	RepoSpy     *auth_repository_spy.AuthRepositorySpy
	HasherSpy   *security_spy.HasherSpy
	CacheSpy    *cache_spy.RedisCacheSpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (create_user.UseCase, *CreateUserFixture) {
	repoSpy := auth_repository_spy.New()
	hasherSpy := security_spy.New()
	cacheSpy := cache_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return create_user_application.New(repoSpy, hasherSpy, cacheSpy, busSpy, settingsSpy),
		&CreateUserFixture{
			RepoSpy:     repoSpy,
			HasherSpy:   hasherSpy,
			CacheSpy:    cacheSpy,
			BusSpy:      busSpy,
			SettingsSpy: settingsSpy,
		}
}

type Option func(*create_user.Input)

func GetInput(options ...Option) *create_user.Input {
	input := &create_user.Input{
		FirstName:            "fakeFirstNameValid",
		LastName:             "fakeLastNameValid",
		Username:             "fake@email.com",
		Password:             "strongPassword123",
		PasswordConfirmation: "strongPassword123",
		CouponCode:           "CouponCd",
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyFirstName() Option {
	return func(params *create_user.Input) {
		params.FirstName = ""
	}
}

func WithInvalidFirstNameLength() Option {
	return func(params *create_user.Input) {
		params.FirstName = _51_Chars //51
	}
}

func WithEmptyLastName() Option {
	return func(params *create_user.Input) {
		params.LastName = ""
	}
}

func WithInvalidLastNameLength() Option {
	return func(params *create_user.Input) {
		params.LastName = _51_Chars
	}
}

func WithEmptyUsername() Option {
	return func(params *create_user.Input) {
		params.Username = ""
	}
}

func WithInvalidUsername() Option {
	return func(params *create_user.Input) {
		params.Username = "invalid-email"
	}
}

func WithEmptyPassword() Option {
	return func(params *create_user.Input) {
		params.Password = ""
	}
}

func WithInvalidPassword(value string) Option {
	return func(params *create_user.Input) {
		params.Password = value
	}
}

func WithEmptyPasswordConfirmation() Option {
	return func(params *create_user.Input) {
		params.PasswordConfirmation = ""
	}
}

func WithInvalidPasswordConfirmation(value string) Option {
	return func(params *create_user.Input) {
		params.PasswordConfirmation = value
	}
}

func WithInvalidCouponCode(value string) Option {
	return func(params *create_user.Input) {
		params.CouponCode = value
	}
}

func WithEmptyCouponCode() Option {
	return func(params *create_user.Input) {
		params.CouponCode = ""
	}
}
