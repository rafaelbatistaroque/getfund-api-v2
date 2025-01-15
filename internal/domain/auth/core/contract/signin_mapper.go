package auth_contract

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
)

type SigninMapper interface {
	ToOutput(token string, session *auth_dto.SessionDto) *signin.Output
	ToSessionModel(authenticatedUser *auth_dto.AuthenticatedUserDto) *auth_dto.SessionDto
}
