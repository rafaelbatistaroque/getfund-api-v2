package authservicefixture

import (
	"getfund-api-v2/internal/domain/auth/domainservice/authservice"
	"getfund-api-v2/test/helper/mapperspy/signinmapperspy"
	"getfund-api-v2/test/helper/securityspy"
	"getfund-api-v2/test/helper/settingsspy"
	"getfund-api-v2/test/helper/userrepositoryspy"
)

func NewSut() (authservice.AuthService, *settingsspy.ApplicationSettingsSpy, *userrepositoryspy.UserRepositorySpy, *securityspy.HasherSpy, *signinmapperspy.SigninMapperSpy) {
	settingsSpy := settingsspy.New()
	userRepositorySpy := userrepositoryspy.New()
	hasherSpy := securityspy.New()
	mapperSpy := signinmapperspy.New()

	return authservice.New(userRepositorySpy, settingsSpy, hasherSpy, mapperSpy),
		settingsSpy,
		userRepositorySpy,
		hasherSpy,
		mapperSpy
}
