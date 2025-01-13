package auth_contract

import "getfund-api-v2/internal/domain/auth/core/auth_dto"

type SessionKey struct{}
type TokenKey struct{}

type SessionService interface {
	SaveSession(session *auth_dto.SessionDto) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}
