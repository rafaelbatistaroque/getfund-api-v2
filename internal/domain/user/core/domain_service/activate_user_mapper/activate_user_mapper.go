package activate_user_mapper

import (
	"getfund-api-v2/internal/domain/user/core/entity/user_entity"
	"getfund-api-v2/internal/domain/user/core/user_dto"
)

type Mapper interface {
	ToDto(entity *user_entity.User) *user_dto.ActivationUserDto
}

type activateUserMapper struct {
}

// Constructor
func New() Mapper {
	return &activateUserMapper{}
}

func (m *activateUserMapper) ToDto(entity *user_entity.User) *user_dto.ActivationUserDto {
	return &user_dto.ActivationUserDto{}
}
