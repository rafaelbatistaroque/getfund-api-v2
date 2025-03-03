package activate_user_mapper

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/entity/user_entity"
)

type Mapper interface {
	ToDto(entity *user_entity.User) *auth_dto.ActivationUserDto
}

type activateUserMapper struct {
}

// Constructor
func New() Mapper {
	return &activateUserMapper{}
}

func (m *activateUserMapper) ToDto(entity *user_entity.User) *auth_dto.ActivationUserDto {
	return &auth_dto.ActivationUserDto{
		FirstName: entity.GetFirstName(),
		LastName:  entity.GetLastName(),
		Username:  entity.GetUsername(),
		Password:  entity.GetPassword(),
		IsAdmin:   entity.GetIsAdmin(),
		IsActive:  entity.GetIsActive(),
		CreatedAt: entity.GetCreatedAt().Unix(),
		UpdatedAt: entity.GetUpdatedAt().Unix(),
	}
}
