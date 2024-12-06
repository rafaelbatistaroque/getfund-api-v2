package userrepositoryfixture

import (
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	authuserrepository "getfund-api-v2/internal/domain/auth/port/repository"
	"getfund-api-v2/test/helper/dbfixture"
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewSUT() (auth_contract.UserRepository, *gorm.DB) {
	db := dbfixture.NewMemoryDB(&FakeUser{})

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
