package signin_mapper_fixture

import (
	model "getfund-api-v2/internal/domain/auth/adapter/model"
	"getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

func NewSut() (signin_mapper.SigninMapper, *model.SessionModel, *security_spy.HasherSpy, *settings_spy.ApplicationSettingsSpy) {
	hasher := security_spy.New()
	settings := settings_spy.New()
	return signin_mapper.New(hasher, settings),
		&model.SessionModel{ID: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1},
		hasher,
		settings

}

func GetUserModel() *model.UserModel {
	return &model.UserModel{Id: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}
}
