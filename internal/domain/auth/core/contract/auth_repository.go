package auth_contract

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
)

type Repository interface {
	GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error)
	UpdatePassword(id int, value string) error
	UpdateUsernameHash(id int, username string) error
	Signup(user *auth_dto.ActivationUserDto) (*auth_dto.UserDto, error)
	UserExists(username string) (*auth_dto.UserDto, error)
}
