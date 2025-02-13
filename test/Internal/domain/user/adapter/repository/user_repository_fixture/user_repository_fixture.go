package user_repository_fixture

import (
	user_repository "getfund-api-v2/internal/domain/user/adapter/repository"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/pkg/db/schema"
	"getfund-api-v2/test/helper/db_fixture"

	"gorm.io/gorm"
)

func NewSUT() (user_contract.Repository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(&schema.User{})

	return user_repository.New(db), db
}

func GetEmptyActivationUserDto() *user_dto.ActivationUserDto {
	return &user_dto.ActivationUserDto{}
}
