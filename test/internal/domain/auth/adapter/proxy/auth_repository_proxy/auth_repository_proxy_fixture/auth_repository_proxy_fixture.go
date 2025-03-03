package auth_repository_proxy_fixture

import (
	"getfund-api-v2/internal/domain/auth/adapter/proxy/auth_repository_proxy"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"

	"github.com/google/uuid"
)

type authRepositoryProxyFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	RepoSpy     *auth_repository_spy.AuthRepositorySpy
}

func NewSut() (auth_contract.Repository, *authRepositoryProxyFixture) {
	settingsSpy := settings_spy.New()
	authRepositorySpy := auth_repository_spy.New()
	hasherSpy := security_spy.New()

	return auth_repository_proxy.New(authRepositorySpy, settingsSpy, hasherSpy),
		&authRepositoryProxyFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			RepoSpy:     authRepositorySpy,
		}
}

func GetEmptyActivationUserDto() *auth_dto.ActivationUserDto {
	return &auth_dto.ActivationUserDto{}
}

func GetFilledActivationUserDto() *auth_dto.ActivationUserDto {
	return &auth_dto.ActivationUserDto{
		FirstName: "fake-first-name",
		LastName:  "fake-last-name",
		Username:  "fake@email.com",
		Password:  "fakeStrongPass123",
		IsAdmin:   false,
	}
}

func GetFakeUsername() string {
	return uuid.NewString()
}
