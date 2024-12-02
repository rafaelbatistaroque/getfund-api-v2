package signinmapperfixture

import (
	"getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	model "getfund-api-v2/internal/domain/auth/model"
	"getfund-api-v2/test/spyshared/securityspy"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() (signinmapper.SigninMapper, *model.SessionModel, *securityspy.HasherSpy, *settingsspy.ApplicationSettingsSpy) {
	hasher := securityspy.New()
	settings := settingsspy.New()
	return signinmapper.New(hasher, settings),
		&model.SessionModel{ID: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1},
		hasher,
		settings

}

func GetUserModel() *model.UserModel {
	return &model.UserModel{Id: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}
}
