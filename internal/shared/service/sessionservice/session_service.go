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

type SessionKey struct{}

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
	//deletar sessão em cache by token
	return nil
}

func (s *sessionService) GetSession(token string) (string, error) {
	//obter sessão do cache
	//decriptar
	//retornar session serializada
	return "", nil
}
