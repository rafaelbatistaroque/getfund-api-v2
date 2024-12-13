package session_service

import (
	"errors"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
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
	cache    cache_service.Cache
	security security.Hasher
	settings settings.ApplicationSettings
}

func New(cache cache_service.Cache, security security.Hasher, settings settings.ApplicationSettings) SessionService {
	return &sessionService{
		cache:    cache,
		security: security,
		settings: settings,
	}
}

func (s *sessionService) SaveSession(session string) (string, error) {
	if session == "" {
		return "", errors.New("save-session: parameter cannot be null or empty")
	}

	token, sessionEncrypted := encryptSession(s.security, s.settings, session)

	errCache := s.cache.Set(token, sessionEncrypted, time_24_HOURS)
	if errCache != nil {
		return "", errCache
	}

	return token, nil
}

func encryptSession(security security.Hasher, settings settings.ApplicationSettings, session string) (string, string) {
	sessionEncrypted := security.Encrypt(session, settings.GetSecretKey())

	return security.HashAndMerge(sessionEncrypted, settings.GetServerSalt()), sessionEncrypted
}

func (s *sessionService) DeleteSession(token string) error {
	if token == "" {
		return errors.New("delete-session: parameter cannot be null or empty")
	}

	err := s.cache.Delete(token)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionService) GetSession(token string) (string, error) {
	if token == "" {
		return "", errors.New("get-session: parameter cannot be null or empty")
	}

	sessionEncrypted, err := s.cache.Get(token)
	if err != nil {
		return "", err
	}

	return s.security.DecryptMerged(sessionEncrypted, s.settings.GetSecretKey()), nil
}
