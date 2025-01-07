package user_repository_fixture

import (
	authuserrepository "getfund-api-v2/internal/domain/auth/adapter/repository"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/test/helper/db_fixture"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSUT() (auth_contract.UserRepository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(&FakeUser{})

	return authuserrepository.New(db), db
}

type FakeUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive int    `json:"is_active"`
}

func (FakeUser) TableName() string {
	return "user"
}

func AddUser(db *gorm.DB, quantity int, diff bool) {
	if diff {
		for range quantity {
			db.Create(FakeUser{ID: uuid.NewString(), Username: uuid.NewString(), IsActive: rand.Intn(2)})
		}
	} else {
		user := uuid.NewString()
		for range quantity {
			db.Create(FakeUser{ID: user, Username: user, IsActive: rand.Intn(2)})
		}
	}

}
