package auth_service_fixture

import (
	"getfund-api-v2/internal/domain/auth/adapter/domain_service/auth_service"
	"getfund-api-v2/test/helper/mapper_spy/signin_mapper_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/user_repository_spy"
)

func NewSut() (auth_service.AuthService, *settings_spy.ApplicationSettingsSpy, *user_repository_spy.UserRepositorySpy, *security_spy.HasherSpy, *signin_mapper_spy.SigninMapperSpy) {
	settingsSpy := settings_spy.New()
	userRepositorySpy := user_repository_spy.New()
	hasherSpy := security_spy.New()
	mapperSpy := signin_mapper_spy.New()

	return auth_service.New(userRepositorySpy, settingsSpy, hasherSpy, mapperSpy),
		settingsSpy,
		userRepositorySpy,
		hasherSpy,
		mapperSpy
}
