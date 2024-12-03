package sessionservicefixture

import (
	"getfund-api-v2/internal/shared/service/sessionservice"
	"getfund-api-v2/test/spyshared/cachespy"
	"getfund-api-v2/test/spyshared/securityspy"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() (sessionservice.SessionService, *securityspy.HasherSpy, *settingsspy.ApplicationSettingsSpy, *cachespy.RedisCacheSpy) {
	settingsSpy := settingsspy.New()
	hasherSpy := securityspy.New()
	redisSpy := cachespy.New()

	return sessionservice.New(redisSpy, hasherSpy, settingsSpy),
		hasherSpy,
		settingsSpy,
		redisSpy
}

func GetSaveSessionInputInvalid() string { return "" }
func GetSaveSessionInputValid() string   { return `{"fakeField": "fake-value"}` }

func GetDeleteSessionInputValid() string { return `{"fakeField": "fake-value"}` }
