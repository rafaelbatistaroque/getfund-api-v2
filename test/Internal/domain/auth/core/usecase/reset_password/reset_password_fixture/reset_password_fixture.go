package reset_password_fixture

import (
	"encoding/json"
	auth_model "getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/reset_password/application"
	"getfund-api-v2/test/helper/auth_repository_spy"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type ResetPasswordFixture struct {
	CacheSpy    *cache_spy.RedisCacheSpy
	AuthRepoSpy *auth_repository_spy.AuthRepositorySpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	HasherSpy   *security_spy.HasherSpy
}

func NewSut() (reset_password.UseCase, *ResetPasswordFixture) {
	cacheSpy := cache_spy.New()
	AuthRepoSpy := auth_repository_spy.New()
	settingsSpy := settings_spy.New()
	hasherSpy := security_spy.New()

	return sut.New(cacheSpy, AuthRepoSpy, settingsSpy, hasherSpy), &ResetPasswordFixture{
		CacheSpy:    cacheSpy,
		AuthRepoSpy: AuthRepoSpy,
		SettingsSpy: settingsSpy,
		HasherSpy:   hasherSpy,
	}
}

func GetForgetPasswordFromGetSuccessCache(cacheSpy *cache_spy.RedisCacheSpy) *auth_model.ForgetPasswordDto {
	expectedParam := &auth_model.ForgetPasswordDto{}
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
