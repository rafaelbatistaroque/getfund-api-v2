package contract

import "getfund-api-v2/internal/domain/auth/core/dto"

type SessionKey struct{}
type TokenKey struct{}

type SessionService interface {
	SaveSession(session *dto.SessionDto) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}
