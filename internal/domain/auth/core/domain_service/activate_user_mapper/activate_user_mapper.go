package activate_user_mapper

import (
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/domain/auth/core/entity"
)

type Mapper interface {
	ToDto(user *entity.User) *dto.ActivationUserDto
}

type activateUserMapper struct {
}

// Constructor
func New() Mapper {
	return &activateUserMapper{}
}

func (m *activateUserMapper) ToDto(user *entity.User) *dto.ActivationUserDto {
	return &dto.ActivationUserDto{
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Username:  user.GetUsername(),
		Password:  user.GetPassword(),
		IsAdmin:   user.GetIsAdmin(),
		IsActive:  user.GetIsActive(),
		CreatedAt: user.GetCreatedAt().Unix(),
		UpdatedAt: user.GetUpdatedAt().Unix(),
	}
}
