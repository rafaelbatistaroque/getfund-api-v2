package activate_user_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/domain/auth/core/entity"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	activate_user_application "getfund-api-v2/internal/domain/auth/core/usecase/activate_user/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/mapper_spy/activate_user_mapper_spy"
	"getfund-api-v2/test/helper/repository_spy/auth_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type ActivateUserFixture struct {
	CacheSpy    *cache_spy.RedisCacheSpy
	RepoSpy     *auth_repository_spy.AuthRepositorySpy
	MapperSpy   *activate_user_mapper_spy.ActivateUserMapperSpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (activate_user.UseCase, *ActivateUserFixture) {
	cacheSpy := cache_spy.New()
	repoSpy := auth_repository_spy.New()
	mapperSpy := activate_user_mapper_spy.New()
	busSpy := eventbus_spy.New()
	settingsSpy := settings_spy.New()

	return activate_user_application.New(cacheSpy, repoSpy, mapperSpy, busSpy, settingsSpy),
		&ActivateUserFixture{
			CacheSpy:    cacheSpy,
			RepoSpy:     repoSpy,
			MapperSpy:   mapperSpy,
			BusSpy:      busSpy,
			SettingsSpy: settingsSpy,
		}
}

type Option func(*activate_user.Input)

func GetInput(options ...Option) *activate_user.Input {
	fakeValidActivationCode := "fakeActivationCodeVl"
	input := &activate_user.Input{
		ActivationCode:    fakeValidActivationCode,
		ActivationDataKey: "user_activation_" + fakeValidActivationCode,
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyActivationCode() Option {
	return func(params *activate_user.Input) {
		params.ActivationCode = ""
	}
}

func WithInvalidActivationCodeLength() Option {
	return func(params *activate_user.Input) {
		params.ActivationCode = "invalid"
	}
}

func WithEmptyActivationDataKey() Option {
	return func(params *activate_user.Input) {
		params.ActivationDataKey = ""
	}
}

func WithInvalidActivationDataKey() Option {
	return func(params *activate_user.Input) {
		params.ActivationDataKey = "invalid-activation-data-key"
	}
}

func GetActivateUserEntity() *entity.User {
	return entity.NewUser(
		"fake-first-name",
		"fake-last-name",
		"fake@email.com",
		"fakaStrongPass123",
	)
}

func GetUserDataWithCouponSerialized() string {
	return `{"first_name":"fake-first-name","last_name":"fake-last-name","username":"fake@email.com","password":"fakaStrongPass123","cupon_code":"COUPONCD"}`
}

func GetUserDataWithoutCouponSerialized() string {
	return `{"first_name":"fake-first-name","last_name":"fake-last-name","username":"fake@email.com","password":"fakaStrongPass123","cupon_code":""}`
}

func GetOutput() *activate_user.Output {
	userData := GetUserDataWithCoupon()
	output := &activate_user.Output{
		Username: userData.Username,
		Password: userData.Password,
	}

	return output
}

func GetUserDataWithCoupon() *dto.ActivationUserData {
	var user = &dto.ActivationUserData{}
	json.Unmarshal([]byte(GetUserDataWithCouponSerialized()), user)

	return user
}

func (s *ActivateUserFixture) GetActivationUserData() *dto.ActivationUserData {
	var user = dto.ActivationUserData{}
	json.Unmarshal([]byte(s.CacheSpy.SuccessResult["Get"].(string)), &user)

	return &user
}
