package contract

import (
	"getfund-api-v2/internal/domain/auth/core/dto"
)

type Repository interface {
	GetAuthenticatedUserByUsername(username string) (*dto.AuthenticatedUserDto, error)
	UpdatePassword(id int, value string) error
	UpdateUsernameHash(id int, username string) error
	Signup(user *dto.ActivationUserDto) (*dto.UserDto, error)
	UserExists(username string) (*dto.UserDto, error)
}
