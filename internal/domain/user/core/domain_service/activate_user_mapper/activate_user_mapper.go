package activate_user_mapper

import (
	"getfund-api-v2/internal/domain/user/core/dto/user_dto"
	"getfund-api-v2/internal/domain/user/core/entity/user_entity"
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
	return &user_dto.ActivationUserDto{
		FirstName:         entity.GetFirstName(),
		LastName:          entity.GetLastName(),
		Email:             entity.GetEmail(),
		Username:          entity.GetUsername(),
		Gender:            entity.GetGender(),
		Password:          entity.GetPassword(),
		CountryId:         entity.GetCountryId(),
		UserCategoryId:    entity.GetUserCategoryId(),
		MainSocialNetwork: entity.GetMainSocialNetwork(),
		RegisteredUrl:     entity.GetRegisteredUrl(),
		IsAdmin:           entity.GetIsAdmin(),
		IsActive:          entity.GetIsActive(),
		CreatedAt:         entity.GetCreatedAt().Unix(),
		UpdatedAt:         entity.GetUpdatedAt().Unix(),
	}
}
