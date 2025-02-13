package user_repository_fixture

import (
	user_repository "getfund-api-v2/internal/domain/user/adapter/repository"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/pkg/db/schema"
	"getfund-api-v2/test/helper/db_fixture"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSUT() (user_contract.Repository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(&schema.User{})

	return user_repository.New(db), db
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
