package session_service_fixture

import (
	"getfund-api-v2/internal/shared/service/session_service"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

func NewSut() (session_service.SessionService, *security_spy.HasherSpy, *settings_spy.ApplicationSettingsSpy, *cache_spy.RedisCacheSpy) {
	settingsSpy := settings_spy.New()
	hasherSpy := security_spy.New()
	redisSpy := cache_spy.New()

	return session_service.New(redisSpy, hasherSpy, settingsSpy),
		hasherSpy,
		settingsSpy,
		redisSpy
}

func GetSaveSessionInputInvalid() string { return "" }
func GetSaveSessionInputValid() string   { return `{"fakeField": "fake-value"}` }

func GetDeleteSessionInputValid() string   { return `{"fakeField": "fake-value"}` }
func GetDeleteSessionInputInvalid() string { return "" }

func GetGetSessionInputValid() string   { return `{"fakeField": "fake-value"}` }
func GetGetSessionInputInvalid() string { return "" }
