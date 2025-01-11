package session_service_proxy

import (
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/session_service"
)

type sessionServiceProxy struct {
	sessionService session_service.SessionService
	hasher         security.Hasher
	settings       settings.ApplicationSettings
}

func New(sessionService session_service.SessionService, hasher security.Hasher, settings settings.ApplicationSettings) session_service.SessionService {
	return &sessionServiceProxy{
		sessionService: sessionService,
		hasher:         hasher,
		settings:       settings,
	}
}

func (s *sessionServiceProxy) SaveSession(session string) (string, error) {
	sessionEncrypted := s.hasher.Encrypt(session, s.settings.GetSecretKey())

	token := s.hasher.HashAndMerge(sessionEncrypted, s.settings.GetServerSalt())

	sessionData := concatSessionData(token, sessionEncrypted)
	return s.sessionService.SaveSession(sessionData)
}

func concatSessionData(token, sessionEncrypted string) string {
	return token + "@" + sessionEncrypted
}

func (s *sessionServiceProxy) DeleteSession(token string) error {
	return s.sessionService.DeleteSession(token)
}

func (s *sessionServiceProxy) GetSession(token string) (string, error) {
	sessionEncrypted, err := s.sessionService.GetSession(token)
	if err != nil {
		return "", err
	}

	return s.hasher.DecryptMerged(sessionEncrypted, s.settings.GetSecretKey()), nil
}
