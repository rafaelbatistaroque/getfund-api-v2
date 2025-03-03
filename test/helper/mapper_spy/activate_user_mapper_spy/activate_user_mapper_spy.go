package activate_user_mapper_spy

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/entity/user_entity"
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

func (m *ActivateUserMapperSpy) ToDto(entity *user_entity.User) *auth_dto.ActivationUserDto {
	m.Params["ToDto:entity"] = entity

	m.CallsCount["ToDto"]++

	success := m.SuccessResult["ToDto"]
	if success != nil {
		return m.SuccessResult["ToDto"].(*auth_dto.ActivationUserDto)
	}

	return nil
}

func (m *ActivateUserMapperSpy) DefineToDtoSuccess(entity *user_entity.User) {
	m.SuccessResult["ToDto"] = &auth_dto.ActivationUserDto{
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
