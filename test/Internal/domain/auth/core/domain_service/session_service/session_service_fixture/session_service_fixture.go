package session_service_fixture

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/session_service"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/security_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type SessionServiceFixture struct {
	RedisCacheSpy *cache_spy.RedisCacheSpy
	HasherSpy     *security_spy.HasherSpy
	SettingsSpy   *settings_spy.ApplicationSettingsSpy
}

func NewSut() (auth_contract.SessionService, *SessionServiceFixture) {
	redisSpy := cache_spy.New()
	hasherSpy := security_spy.New()
	settingsSpy := settings_spy.New()

	return session_service.New(redisSpy, hasherSpy, settingsSpy),
		&SessionServiceFixture{
			RedisCacheSpy: redisSpy,
			HasherSpy:     hasherSpy,
			SettingsSpy:   settingsSpy,
		}
}

func GetSaveSessionInputNull() *auth_dto.SessionDto { return nil }
func GetSaveSessionInputValid() *auth_dto.SessionDto {
	return &auth_dto.SessionDto{}
}
func GetSaveSessionInputValidSerialized() string {
	return "{\"id\":\"\",\"first_name\":\"\",\"is_admin\":0}"
}

func GetDeleteSessionInputInvalid() string { return "" }
func GetDeleteSessionInputValid() string   { return `fake-token` }

func GetGetSessionInputInvalid() string { return "" }
func GetGetSessionInputValid() string   { return `fake-token` }
