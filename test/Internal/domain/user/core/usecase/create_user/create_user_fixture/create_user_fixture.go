package create_user_fixture

import (
	"getfund-api-v2/internal/domain/user/core/usercase/create_user"
	create_user_application "getfund-api-v2/internal/domain/user/core/usercase/create_user/application"
)

const (
	_51_Chars = "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"
)

type CreateUserFixture struct {
}

func NewSut() (create_user.UseCase, *CreateUserFixture) {

	return create_user_application.New(), &CreateUserFixture{}
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
		RegisteredUrl:     "http://fakeurl.com",
		CuponCode:         "fakeCuponCode",
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
