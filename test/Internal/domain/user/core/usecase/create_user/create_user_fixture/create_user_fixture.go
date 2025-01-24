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
	input := &create_user.Input{FirstName: "fakeFirstNameValid", Password: "strongPassword123"}

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
