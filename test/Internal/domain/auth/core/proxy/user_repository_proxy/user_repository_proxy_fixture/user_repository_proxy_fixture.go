package user_repository_proxy_fixture

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/proxy/user_repository_proxy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/user_repository_spy"
)

type UserRepositoryProxyFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	UserRepoSpy *user_repository_spy.UserRepositorySpy
}

func NewSut() (auth_contract.UserRepository, *UserRepositoryProxyFixture) {
	settingsSpy := settings_spy.New()
	userRepositorySpy := user_repository_spy.New()
	hasherSpy := security_spy.New()

	return user_repository_proxy.New(userRepositorySpy, settingsSpy, hasherSpy),
		&UserRepositoryProxyFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			UserRepoSpy: userRepositorySpy,
		}
}
