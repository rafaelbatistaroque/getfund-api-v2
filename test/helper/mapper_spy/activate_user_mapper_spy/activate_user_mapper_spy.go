package activate_user_mapper_spy

import (
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/domain/auth/core/entity"
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

func (m *ActivateUserMapperSpy) ToDto(entity *entity.User) *dto.ActivationUserDto {
	m.Params["ToDto:entity"] = entity

	m.CallsCount["ToDto"]++

	success := m.SuccessResult["ToDto"]
	if success != nil {
		return m.SuccessResult["ToDto"].(*dto.ActivationUserDto)
	}

	return nil
}

func (m *ActivateUserMapperSpy) DefineToDtoSuccess(entity *entity.User) {
	m.SuccessResult["ToDto"] = &dto.ActivationUserDto{
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
