package auth_repository_proxy_fixture

import (
	"getfund-api-v2/internal/domain/auth/adapter/proxy/user_repository_proxy"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type authRepositoryProxyFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	AuthRepoSpy *auth_repository_spy.AuthRepositorySpy
}

func NewSut() (auth_contract.AuthRepository, *authRepositoryProxyFixture) {
	settingsSpy := settings_spy.New()
	authRepositorySpy := auth_repository_spy.New()
	hasherSpy := security_spy.New()

	return user_repository_proxy.New(authRepositorySpy, settingsSpy, hasherSpy),
		&authRepositoryProxyFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			AuthRepoSpy: authRepositorySpy,
		}
}
