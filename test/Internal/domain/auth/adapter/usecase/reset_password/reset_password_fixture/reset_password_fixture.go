package reset_password_fixture

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password"
	sut "getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password/application"
)

type ResetPasswordFixture struct {
}

func NewSut() (reset_password.UseCase, *ResetPasswordFixture) {

	return sut.New(), nil
}

type Option func(*reset_password.Input)

func GetInput(options ...Option) *reset_password.Input {
	input := &reset_password.Input{RecoveryCode: "fake-al44iyayjdL7fpW93wEXnkfXotDrx2krPjoREOhIgO7QD4j5MIEtWe5bSxO", Password: "strongPassword123"}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithRecoveryCodeEmpty() Option {
	return func(params *reset_password.Input) {
		params.RecoveryCode = ""
	}
}

func WithRecoveryCodeInvalidLength() Option {
	return func(params *reset_password.Input) {
		params.RecoveryCode = "invalid_recovery_code"
	}
}

func WithPasswordEmpty() Option {
	return func(params *reset_password.Input) {
		params.Password = ""
	}
}

func WithInvalidLengthPassword() Option {
	return func(params *reset_password.Input) {
		params.Password = "-"
	}
}

func WithInvalidUpperPassword() Option {
	return func(params *reset_password.Input) {
		params.Password = "abcdef12"
	}
}

func WithInvalidLowerPassword() Option {
	return func(params *reset_password.Input) {
		params.Password = "ABCDEF12"
	}
}

func WithInvalidDigitPassword() Option {
	return func(params *reset_password.Input) {
		params.Password = "ABCDefgh"
	}
}
