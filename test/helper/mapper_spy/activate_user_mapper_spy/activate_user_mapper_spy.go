package activate_user_mapper_spy

import (
	"getfund-api-v2/internal/domain/user/core/entity/activate_user_entity"
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

func (m *ActivateUserMapperSpy) ToEntity(data *user_dto.ActivationUserData) *activate_user_entity.ActivationUser {
	m.Params["ToEntity:data"] = data

	m.CallsCount["ToEntity"]++

	if m.ForceReturn {
		return nil
	}

	success := m.SuccessResult["ToEntity"]
	if success != nil {
		return m.SuccessResult["ToEntity"].(*activate_user_entity.ActivationUser)
	}

	return m.SuccessResult["ToEntity"].(*activate_user_entity.ActivationUser)
}

func (m *ActivateUserMapperSpy) ToDto(entity *activate_user_entity.ActivationUser) *user_dto.ActivationUserDto {
	m.Params["ToDto:entity"] = entity

	success := m.SuccessResult["ToDto"]
	if success != nil {
		return m.SuccessResult["ToDto"].(*user_dto.ActivationUserDto)
	}

	return nil
}

func (m *ActivateUserMapperSpy) DefineToEntitySuccess(data *user_dto.ActivationUserData) {
	m.SuccessResult["ToEntity"] = activate_user_entity.New(
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

func (m *ActivateUserMapperSpy) DefineToDtoSuccess(entity *activate_user_entity.ActivationUser) {
	m.SuccessResult["ToDto"] = user_dto.ActivationUserDto{
		FirstName:         entity.GetFirstName(),
		LastName:          entity.GetLastName(),
		Email:             entity.GetEmail(),
		Gender:            entity.GetGender(),
		MainSocialNetwork: entity.GetMainSocialNetwork(),
		RegisteredUrl:     entity.GetRegisteredUrl(),
		CountryId:         entity.GetCountryId(),
		UserCategoryId:    entity.GetUserCategoryId(),
		RegisteredAt:      entity.GetRegisteredAt(),
		IsAdmin:           entity.GetIsAdmin(),
		IsActive:          entity.GetIsActive(),
	}
}
