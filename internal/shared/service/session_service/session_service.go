package session_service

import (
	"errors"
	"getfund-api-v2/internal/shared/service/cache_service"
	"strings"
	"time"
)

var (
	time_24_HOURS = 24 * time.Hour
)

type SessionKey struct{}
type TokenKey struct{}

type SessionService interface {
	SaveSession(session string) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}

type sessionService struct {
	cache cache_service.Cache
}

func New(cache cache_service.Cache) SessionService {
	return &sessionService{
		cache: cache,
	}
}

func (s *sessionService) SaveSession(session string) (string, error) {
	if session == "" {
		return "", errors.New("save-session: parameter cannot be null or empty")
	}

	if notContain(session, "@") {
		return "", errors.New("save-session: parameter invalid")
	}

	sessionData := strings.Split(session, "@")

	errCache := s.cache.Set(sessionData[0], sessionData[1], time_24_HOURS)
	if errCache != nil {
		return "", errCache
	}

	return sessionData[0], nil
}

func notContain(value, param string) bool {
	return !strings.Contains(value, param)
}

func (s *sessionService) DeleteSession(token string) error {
	if token == "" {
		return errors.New("delete-session: parameter cannot be null or empty")
	}

	return s.cache.Delete(token)
}

func (s *sessionService) GetSession(token string) (string, error) {
	if token == "" {
		return "", errors.New("get-session: parameter cannot be null or empty")
	}

	return s.cache.Get(token)
}
