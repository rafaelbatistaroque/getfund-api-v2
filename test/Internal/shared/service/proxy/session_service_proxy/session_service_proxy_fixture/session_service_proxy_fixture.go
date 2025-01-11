package session_service_proxy_fixture

import (
	"getfund-api-v2/internal/shared/service/proxy/session_service_proxy"
	"getfund-api-v2/internal/shared/service/session_service"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/session_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type SessionServiceProxyFixture struct {
	HasherSpy   *security_spy.HasherSpy
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	SessionSpy  *session_spy.SessionServiceSpy
}

func NewSut() (session_service.SessionService, *SessionServiceProxyFixture) {
	settingsSpy := settings_spy.New()
	hasherSpy := security_spy.New()
	sessionServiceSpy := session_spy.New()

	return session_service_proxy.New(sessionServiceSpy, hasherSpy, settingsSpy),
		&SessionServiceProxyFixture{
			HasherSpy:   hasherSpy,
			SettingsSpy: settingsSpy,
			SessionSpy:  sessionServiceSpy,
		}
}

func GetSaveSessionInputValid() string   { return `{"fakeField": "fake-value"}` }
func GetDeleteSessionInputValid() string { return `fake-token` }
func GetGetSessionInputValid() string    { return `fake-token` }
