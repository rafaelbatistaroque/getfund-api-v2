package session_service

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"time"
)

var (
	time_24_HOURS = 24 * time.Hour
)

type sessionService struct {
	cache    cache_service.Cache
	hasher   security.Hasher
	settings settings.ApplicationSettings
}

func New(cache cache_service.Cache, hasher security.Hasher, settings settings.ApplicationSettings) auth_contract.SessionService {
	return &sessionService{
		cache:    cache,
		hasher:   hasher,
		settings: settings,
	}
}

func (s *sessionService) SaveSession(session *auth_dto.SessionDto) (string, error) {
	if session == nil {
		return "", errors.New("save-session: session cannot be null or empty")
	}

	sessionSerialized, _ := json.Marshal(session)
	sessionEncrypted := s.hasher.Encrypt(string(sessionSerialized), s.settings.GetSecretKey())
	token := s.hasher.HashAndMerge(sessionEncrypted, s.settings.GetServerSalt())

	errCache := s.cache.Set(token, sessionEncrypted, time_24_HOURS)
	if errCache != nil {
		return "", errCache
	}

	return token, nil
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

	sessionEncrypted, err := s.cache.Get(token)
	if err != nil {
		return "", err
	}

	return s.hasher.DecryptMerged(sessionEncrypted, s.settings.GetSecretKey()), nil
}
