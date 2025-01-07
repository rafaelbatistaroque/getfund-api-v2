package reset_password_fixture

import (
	"encoding/json"
	auth_model "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/reset_password/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/user_repository_spy"
)

type ResetPasswordFixture struct {
	CacheSpy      *cache_spy.RedisCacheSpy
	RepositorySpy *user_repository_spy.UserRepositorySpy
}

func NewSut() (reset_password.UseCase, *ResetPasswordFixture) {
	cacheSpy := cache_spy.New()
	repo := user_repository_spy.New()

	return sut.New(cacheSpy, repo), &ResetPasswordFixture{
		CacheSpy:      cacheSpy,
		RepositorySpy: repo,
	}
}

func GetForgetPasswordFromGetSuccessCache(cacheSpy *cache_spy.RedisCacheSpy) *auth_model.ForgetPasswordModel {
	expectedParam := &auth_model.ForgetPasswordModel{}
	json.Unmarshal([]byte(cacheSpy.SuccessResult["Get"].(string)), expectedParam)

	return expectedParam
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
