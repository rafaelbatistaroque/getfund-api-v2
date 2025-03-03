package auth_repository_fixture

import (
	auth_repository "getfund-api-v2/internal/domain/auth/adapter/repository"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/pkg/db/schema"
	"getfund-api-v2/test/helper/db_fixture"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSUT() (auth_contract.Repository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(&schema.User{})

	return auth_repository.New(db), db
}

func AddRandomUser(db *gorm.DB, quantity int, diff bool) {
	if diff {
		for range quantity {
			isAdmin := rand.Intn(2) == 1
			db.Create(schema.User{Username: uuid.NewString(), IsActive: isAdmin, Password: uuid.NewString()})
		}
	} else {
		user := uuid.NewString()
		for range quantity {
			isAdmin := rand.Intn(2) == 1
			db.Create(schema.User{Username: user, IsActive: isAdmin, Password: uuid.NewString()})
		}
	}
}

func GetEmptyActivationUserDto() *auth_dto.ActivationUserDto {
	return &auth_dto.ActivationUserDto{}
}

func GetFilledActivationUserDto() *auth_dto.ActivationUserDto {
	return &auth_dto.ActivationUserDto{
		FirstName: "fake-first-name",
		LastName:  "fake-last-name",
		Username:  "fake@email.com",
		Password:  "fakaStrongPass123",
		IsAdmin:   false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
