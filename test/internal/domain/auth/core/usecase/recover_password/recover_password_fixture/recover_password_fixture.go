package recover_password_fixture

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	sut "getfund-api-v2/internal/domain/auth/core/usecase/recover_password/application"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"

	"github.com/rafaelbatistaroque/validation"
)

type RecoverPasswordFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	RepoSpy     *auth_repository_spy.AuthRepositorySpy
	CacheSpy    *cache_spy.RedisCacheSpy
	EventBusSpy *eventbus_spy.EventBusSpy
}

func NewSut() (recover_password.UseCase, *RecoverPasswordFixture) {
	hasherSpy := security_spy.New()
	settingsSpy := settings_spy.New()
	AuthRepoSpy := auth_repository_spy.New()
	cacheSpy := cache_spy.New()
	eventBusSpy := eventbus_spy.New()

	return sut.New(hasherSpy, settingsSpy, AuthRepoSpy, cacheSpy, eventBusSpy),
		&RecoverPasswordFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			RepoSpy:     AuthRepoSpy,
			CacheSpy:    cacheSpy,
			EventBusSpy: eventBusSpy,
		}
}

func GetInvalidInputWithError() (*recover_password.Input, *shared_error.Error) {
	return &recover_password.Input{Username: ""},
		shared_error.New(shared_error.BAD_REQUEST_CODE, fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Username"))
}

func GetValidInput() *recover_password.Input {
	return &recover_password.Input{Username: "fake-username"}
}
