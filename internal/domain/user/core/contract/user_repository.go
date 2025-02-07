package user_contract

import "getfund-api-v2/internal/domain/user/core/user_dto"

type Repository interface {
	GetUserByUsername(username string) (*user_dto.UserDto, error)
	SaveUser(user *user_dto.ActivationUserDto) error
}
