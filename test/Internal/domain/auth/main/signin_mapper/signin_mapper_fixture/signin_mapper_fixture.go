package signin_mapper_fixture

import (
	model "getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
)

func NewSut() (auth_contract.SigninMapper, *model.SessionDto) {
	return signin_mapper.New(),
		&model.SessionDto{ID: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}

}

func GetauthenticatedUser() *model.AuthenticatedUserDto {
	return &model.AuthenticatedUserDto{Id: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}
}
