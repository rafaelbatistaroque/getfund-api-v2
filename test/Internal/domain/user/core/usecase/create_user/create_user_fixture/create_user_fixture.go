package create_user_fixture

import (
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	create_user_application "getfund-api-v2/internal/domain/user/core/usecase/create_user/application"
	"getfund-api-v2/test/helper/repository_spy/user_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
)

const (
	_51_Chars = "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"
)

type CreateUserFixture struct {
	RepoSpy   *user_repository_spy.UserRepositorySpy
	HasherSpy *security_spy.HasherSpy
}

func NewSut() (create_user.UseCase, *CreateUserFixture) {
	userRepoSpy := user_repository_spy.New()
	hasherSpy := security_spy.New()

	return create_user_application.New(userRepoSpy, hasherSpy),
		&CreateUserFixture{
			RepoSpy:   userRepoSpy,
			HasherSpy: hasherSpy,
		}
}

type Option func(*create_user.Input)

func GetInput(options ...Option) *create_user.Input {
	input := &create_user.Input{
		FirstName:         "fakeFirstNameValid",
		LastName:          "fakeLastNameValid",
		Password:          "strongPassword123",
		Email:             "fake@mail.com",
		Gender:            "m",
		CountryId:         1,
		UserCategoryId:    1,
		MainSocialNetwork: "@fakeUser",
		RegisteredUrl:     "https://fakeurl.com",
		CouponCode:        "CouponCd",
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

func WithInvalidEmail() Option {
	return func(params *create_user.Input) {
		params.Email = ""
	}
}

func WithEmptyGender() Option {
	return func(params *create_user.Input) {
		params.Gender = ""
	}
}

func WithInvalidGender() Option {
	return func(params *create_user.Input) {
		params.Gender = "z"
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

func WithEmptyMainSocialNetwork() Option {
	return func(params *create_user.Input) {
		params.MainSocialNetwork = ""
	}
}

func WithInvalidMainSocialNetwork(value string) Option {
	return func(params *create_user.Input) {
		params.MainSocialNetwork = value
	}
}

func WithEmptyRegisteredUrl() Option {
	return func(params *create_user.Input) {
		params.RegisteredUrl = ""
	}
}

func WithInvalidCouponCode(value string) Option {
	return func(params *create_user.Input) {
		params.CouponCode = value
	}
}
