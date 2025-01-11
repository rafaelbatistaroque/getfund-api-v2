package session_service_fixture

import (
	"getfund-api-v2/internal/shared/service/session_service"
	"getfund-api-v2/test/helper/cache_spy"
)

type SessionServiceFixture struct {
	RedisCacheSpy *cache_spy.RedisCacheSpy
}

func NewSut() (session_service.SessionService, *SessionServiceFixture) {
	redisSpy := cache_spy.New()

	return session_service.New(redisSpy),
		&SessionServiceFixture{
			RedisCacheSpy: redisSpy,
		}
}

func GetSaveSessionInputEmpty() string   { return "" }
func GetSaveSessionInputInvalid() string { return "invalid" }
func GetSaveSessionInputValid() string   { return `fake-token@fake-session-hashed` }

func GetDeleteSessionInputInvalid() string { return "" }
func GetDeleteSessionInputValid() string   { return `fake-token` }

func GetGetSessionInputInvalid() string { return "" }
func GetGetSessionInputValid() string   { return `fake-token` }
