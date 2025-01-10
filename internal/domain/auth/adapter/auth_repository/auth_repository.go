package auth_repository

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"

	"gorm.io/gorm"
)

var (
	table_USER = "user"
	active     = 1
	first      = 1
)

type authRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth_contract.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	var authenticatedUser = auth_dto.AuthenticatedUserDto{}
	result := r.db.
		Table(table_USER).
		Select("id, first_name, is_admin, password, username").
		Where("is_active=? AND username=?", active, username).
		Limit(first).
		Scan(&authenticatedUser)

	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &authenticatedUser, nil
}

func (r *authRepository) UpdatePassword(id, value string) error {
	result := r.db.Table(table_USER).Where("id=?", id).Update("password", value)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
