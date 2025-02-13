package user_respository_proxy_fixture

import (
	"getfund-api-v2/internal/domain/user/adapter/proxy/user_repository_proxy"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/test/helper/repository_spy/user_repository_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type userRepositoryProxyFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	RepoSpy     *user_repository_spy.UserRepositorySpy
}

func NewSut() (user_contract.Repository, *userRepositoryProxyFixture) {
	settingsSpy := settings_spy.New()
	repositorySpy := user_repository_spy.New()
	hasherSpy := security_spy.New()

	return user_repository_proxy.New(repositorySpy, settingsSpy, hasherSpy),
		&userRepositoryProxyFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			RepoSpy:     repositorySpy,
		}
}

func GetEmptyActivationUserDto() *user_dto.ActivationUserDto {
	return &user_dto.ActivationUserDto{}
}

func GetFilledActivationUserDto() *user_dto.ActivationUserDto {
	return &user_dto.ActivationUserDto{
		FirstName:         "fake-first-name",
		LastName:          "fake-last-name",
		Email:             "fake@email.com",
		Username:          "fake@email.com",
		Gender:            "m",
		Password:          "fakeStrongPass123",
		CountryId:         1,
		UserCategoryId:    1,
		MainSocialNetwork: "@FakeSocial",
		RegisteredUrl:     "https://social.com",
		IsAdmin:           false,
	}
}
