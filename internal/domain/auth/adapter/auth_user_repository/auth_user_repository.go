package auth_user_repository

import (
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	model "getfund-api-v2/internal/domain/auth/core/model"

	"gorm.io/gorm"
)

var (
	table_USER = "user"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth_contract.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUserName(username string) (*model.UserModel, error) {
	var user = model.UserModel{}
	err := r.db.
		Table(table_USER).
		Where("username = ? AND is_active = 1", username).
		First(&user)

	if err.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
