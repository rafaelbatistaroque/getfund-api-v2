package activate_user_mapper

import (
	"getfund-api-v2/internal/domain/user/core/entity/activate_user_entity"
	"getfund-api-v2/internal/domain/user/core/user_dto"
)

type Mapper interface {
	ToEntity(data *user_dto.ActivationUserData) *activate_user_entity.ActivationUser
	ToDto(entity *activate_user_entity.ActivationUser) *user_dto.ActivationUserDto
}

type activateUserMapper struct {
}

// Constructor
func New() Mapper {
	return &activateUserMapper{}
}

func (m *activateUserMapper) ToEntity(data *user_dto.ActivationUserData) *activate_user_entity.ActivationUser {
	return activate_user_entity.New(
		data.FirstName,
		data.LastName,
		data.Email,
		data.Gender,
		data.Password,
		data.MainSocialNetwork,
		data.RegisteredUrl,
		data.CountryId,
		data.UserCategoryId)
}

func (m *activateUserMapper) ToDto(entity *activate_user_entity.ActivationUser) *user_dto.ActivationUserDto {
	return &user_dto.ActivationUserDto{}
}
