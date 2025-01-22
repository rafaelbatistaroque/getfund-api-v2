package create_user_fixture

import (
	"getfund-api-v2/internal/domain/user/core/usercase/create_user"
	create_user_application "getfund-api-v2/internal/domain/user/core/usercase/create_user/application"
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

func WithFirstNameEmpty() Option {
	return func(params *create_user.Input) {
		params.FirstName = ""
	}
}

func WithFirstNameLengthInvalid() Option {
	return func(params *create_user.Input) {
		params.FirstName = "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii" //51
	}
}

func WithLastNameEmpty() Option {
	return func(params *create_user.Input) {
		params.LastName = ""
	}
}

func WithLastNameLengthInvalid() Option {
	return func(params *create_user.Input) {
		params.LastName = "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii" //51
	}
}
