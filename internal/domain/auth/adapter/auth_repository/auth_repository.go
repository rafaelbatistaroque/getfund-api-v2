package auth_repository

import (
	"errors"
	model "getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"

	"gorm.io/gorm"
)

var (
	table_USER = "user"
)

type authRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth_contract.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAuthenticatedUserByUsername(username string) (*model.AuthenticatedUserDto, error) {
	var authenticatedUser = model.AuthenticatedUserDto{}
	err := r.db.
		Table(table_USER).
		Where("username = ? AND is_active = 1", username).
		First(&authenticatedUser)

	if err.Error != nil {
		return nil, errors.New("user not found")
	}

	return &authenticatedUser, nil
}

func (r *authRepository) UpdatePassword(id, value string) error {
	return nil
}
