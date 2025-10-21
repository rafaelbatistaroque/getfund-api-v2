package signin_mapper_fixture

import (
	"getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	model "getfund-api-v2/internal/domain/auth/core/dto"
)

func NewSut() (signin_mapper.SigninMapper, *model.SessionDto) {
	return signin_mapper.New(),
		&model.SessionDto{ID: 1, FirstName: "fake-first-name", IsAdmin: true}

}

func GetauthenticatedUser() *model.AuthenticatedUserDto {
	return &model.AuthenticatedUserDto{Id: 1, FirstName: "fake-first-name", IsAdmin: true}
}
