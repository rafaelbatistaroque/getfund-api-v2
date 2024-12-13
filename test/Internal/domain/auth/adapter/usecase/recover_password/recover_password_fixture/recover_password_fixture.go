package recover_password_fixture

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password"
	sut "getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password/recover_password_application"
	"getfund-api-v2/internal/shared/result_app"
	inputvalidation "getfund-api-v2/pkg/input_validation"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/code_spy"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/user_repository_spy"
)

type RecoverPasswordFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	UserRepoSpy *user_repository_spy.UserRepositorySpy
	CodeSpy     *code_spy.CodeSpy
	CacheSpy    *cache_spy.RedisCacheSpy
	EventBusSpy *eventbus_spy.EventBusSpy
}

func NewSut() (recover_password.UseCase, *RecoverPasswordFixture) {
	hasherSpy := security_spy.New()
	settingsSpy := settings_spy.New()
	userRepoSpy := user_repository_spy.New()
	codeSpy := code_spy.New()
	cacheSpy := cache_spy.New()
	eventBusSpy := eventbus_spy.New()

	return sut.New(hasherSpy, settingsSpy, userRepoSpy, codeSpy, cacheSpy, eventBusSpy),
		&RecoverPasswordFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			UserRepoSpy: userRepoSpy,
			CodeSpy:     codeSpy,
			CacheSpy:    cacheSpy,
			EventBusSpy: eventBusSpy,
		}
}

func GetInvalidInputWithError() (*recover_password.Input, *result_app.ApplicationError) {
	return &recover_password.Input{Username: ""},
		result_app.New(result_app.BAD_REQUEST_CODE, fmt.Errorf(inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "Username"))
}

func GetValidInput() *recover_password.Input {
	return &recover_password.Input{Username: "fake-username"}
}
