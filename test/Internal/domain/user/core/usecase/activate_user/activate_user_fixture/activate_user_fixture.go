package activate_user_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/user/core/entity/activate_user_entity"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	activate_user_application "getfund-api-v2/internal/domain/user/core/usecase/activate_user/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/eventbus_spy"
	"getfund-api-v2/test/helper/mapper_spy/activate_user_mapper_spy"
	"getfund-api-v2/test/helper/repository_spy/user_repository_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type ActivateUserFixture struct {
	CacheSpy    *cache_spy.RedisCacheSpy
	RepoSpy     *user_repository_spy.UserRepositorySpy
	MapperSpy   *activate_user_mapper_spy.ActivateUserMapperSpy
	BusSpy      *eventbus_spy.EventBusSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSut() (activate_user.UseCase, *ActivateUserFixture) {
	cacheSpy := cache_spy.New()
	repoSpy := user_repository_spy.New()
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
		ActivationCode: fakeValidActivationCode,
		ActivationKey:  "user_activation_" + fakeValidActivationCode,
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

func WithEmptyActivationKey() Option {
	return func(params *activate_user.Input) {
		params.ActivationKey = ""
	}
}

func WithInvalidActivationKey() Option {
	return func(params *activate_user.Input) {
		params.ActivationKey = "invalid-activation-key"
	}
}

func GetActivateUserEntity() *activate_user_entity.ActivationUser {
	return activate_user_entity.New(
		"fake-first-name",
		"fake-last-name",
		"fake@email.com",
		"m",
		"fakaStrongPass123",
		"@FakeSocial",
		"https://social.com",
		1,
		1)
}

func GetUserDataWithCouponSerialized() string {
	return `{"first_name":"fake-first-name","last_name":"fake-last-name","email":"fake@email.com","gender":"m","password":"fakaStrongPass123","country_id":1,"user_category_id":1,"main_social_network":"@FakeSocial","registered_url":"https://social.com","cupon_code":"COUPONCD"}`
}

func GetUserDataWithoutCouponSerialized() string {
	return `{"first_name":"fake-first-name","last_name":"fake-last-name","email":"fake@email.com","gender":"m","password":"fakaStrongPass123","country_id":1,"user_category_id":1,"main_social_network":"@FakeSocial","registered_url":"https://social.com","cupon_code":""}`
}

func GetResponseSession() ([]byte, *activate_user.Output) {
	response := &activate_user.Output{
		Token: "fake-token",
		Session: activate_user.SessionOutput{
			ID:        "fake-id",
			FirstName: "fake-first-name",
			IsAdmin:   true,
		},
	}
	responseSerialized, err := json.Marshal(response)
	if err != nil {
		return nil, nil
	}
	return responseSerialized, response
}
