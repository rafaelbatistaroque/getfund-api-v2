package activate_user_mapper_spy

import (
	"getfund-api-v2/internal/domain/user/core/entity/user_entity"
	"getfund-api-v2/internal/domain/user/core/user_dto"
)

type ActivateUserMapperSpy struct {
	Params      map[string]any
	ForceReturn bool

	CallsCount map[string]int

	SuccessResult map[string]any
	ErrorResult   map[string]error
}

func New() *ActivateUserMapperSpy {
	return &ActivateUserMapperSpy{Params: make(map[string]any), ForceReturn: true, ErrorResult: make(map[string]error), CallsCount: make(map[string]int), SuccessResult: make(map[string]any)}
}

func (m *ActivateUserMapperSpy) ToDto(entity *user_entity.User) *user_dto.ActivationUserDto {
	m.Params["ToDto:entity"] = entity

	m.CallsCount["ToDto"]++

	success := m.SuccessResult["ToDto"]
	if success != nil {
		return m.SuccessResult["ToDto"].(*user_dto.ActivationUserDto)
	}

	return nil
}

func (m *ActivateUserMapperSpy) DefineToDtoSuccess(entity *user_entity.User) {
	m.SuccessResult["ToDto"] = &user_dto.ActivationUserDto{
		FirstName:         entity.GetFirstName(),
		LastName:          entity.GetLastName(),
		Email:             entity.GetEmail(),
		Gender:            entity.GetGender(),
		MainSocialNetwork: entity.GetMainSocialNetwork(),
		RegisteredUrl:     entity.GetRegisteredUrl(),
		CountryId:         entity.GetCountryId(),
		UserCategoryId:    entity.GetUserCategoryId(),
		IsAdmin:           entity.GetIsAdmin(),
		IsActive:          entity.GetIsActive(),
	}
}
