package reset_password_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/reset_password/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
)

type ResetPasswordFixture struct {
	CacheSpy *cache_spy.RedisCacheSpy
	RepoSpy  *auth_repository_spy.AuthRepositorySpy
}

func NewSut() (reset_password.UseCase, *ResetPasswordFixture) {
	cacheSpy := cache_spy.New()
	repoSpy := auth_repository_spy.New()

	return sut.New(cacheSpy, repoSpy), &ResetPasswordFixture{
		CacheSpy: cacheSpy,
		RepoSpy:  repoSpy,
	}
}

func GetForgetPasswordFromGetSuccessCache(cacheSpy *cache_spy.RedisCacheSpy) *auth_dto.ForgetPasswordDto {
	expectedParam := &auth_dto.ForgetPasswordDto{}
	json.Unmarshal([]byte(cacheSpy.SuccessResult["Get"].(string)), expectedParam)

	return expectedParam
}

type Option func(*reset_password.Input)

func GetInput(options ...Option) *reset_password.Input {
	fakeCode := "fake-al44iyayjdL7fpW93wEXnkfXotDrx2krPjoREOhIgO7QD4j5MIEtWe5bSxO"
	input := &reset_password.Input{
		RecoveryCode: fakeCode,
		Password:     "strongPassword123",
		RecoveryKey:  "recovery_password_" + fakeCode,
	}

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

func WithRecoveryKeyEmpty() Option {
	return func(params *reset_password.Input) {
		params.RecoveryKey = ""
	}
}

func WithRecoveryKeyInvalid() Option {
	return func(params *reset_password.Input) {
		params.RecoveryKey = "invalid-recovery-key"
	}
}
