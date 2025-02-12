package auth_contract

import model "getfund-api-v2/internal/domain/auth/core/auth_dto"

type AuthRepository interface {
	GetAuthenticatedUserByUsername(username string) (*model.AuthenticatedUserDto, error)
	UpdatePassword(id int, value string) error
}
