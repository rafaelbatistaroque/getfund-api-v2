package signin_mapper_fixture

import (
	model "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
)

func NewSut() (signin_mapper.SigninMapper, *model.SessionModel) {
	return signin_mapper.New(),
		&model.SessionModel{ID: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}

}

func GetUserModel() *model.UserModel {
	return &model.UserModel{Id: "fake-ID", FirstName: "fake-first-name", IsAdmin: 1}
}
