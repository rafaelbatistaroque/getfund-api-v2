package sessionservice

import (
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"time"
)

var (
	time_24_HOURS = 24 * time.Hour
)

type SessionService interface {
	SaveSession(session string) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}

type sessionService struct {
	cache    cacheservice.Cache
	security security.Hasher
	settings settings.ApplicationSettings
}

func New(cache cacheservice.Cache, security security.Hasher, settings settings.ApplicationSettings) SessionService {
	return &sessionService{
		cache:    cache,
		security: security,
		settings: settings,
	}
}

func (s *sessionService) SaveSession(session string) (string, error) {
	sessionEncrypted := s.security.Encrypt(session, s.settings.GetSecretKey())

	token := s.security.HashAndMerge(sessionEncrypted, s.settings.GetServerSalt())

	errCache := s.cache.Set(token, sessionEncrypted, time_24_HOURS)
	if errCache != nil {
		return "", errCache
	}

	return token, nil
}

func (s *sessionService) DeleteSession(token string) error {
	//deletar sessão em cache by token
	return nil
}

func (s *sessionService) GetSession(token string) (string, error) {
	//obter sessão do cache
	//decriptar
	//retornar session serializada
	return "", nil
}
