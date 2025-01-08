package auth_service_fixture

import (
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/test/helper/mapper_spy/signin_mapper_spy"
	"getfund-api-v2/test/helper/proxy_spy/user_repository_proxy_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type AuthServiceFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	UserRepoSpy *user_repository_proxy_spy.UserRepositoryProxySpy
	MapperSpy   *signin_mapper_spy.SigninMapperSpy
}

func NewSut() (auth_service.AuthService, *AuthServiceFixture) {
	settingsSpy := settings_spy.New()
	userRepositoryProxySpy := user_repository_proxy_spy.New()
	hasherSpy := security_spy.New()
	mapperSpy := signin_mapper_spy.New()

	return auth_service.New(userRepositoryProxySpy, settingsSpy, hasherSpy, mapperSpy),
		&AuthServiceFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			UserRepoSpy: userRepositoryProxySpy,
			MapperSpy:   mapperSpy,
		}
}
