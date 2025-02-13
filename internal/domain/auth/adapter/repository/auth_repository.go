package auth_repository

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/pkg/db/schema"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth_contract.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	var user = schema.User{}
	result := r.db.
		Select("id, first_name, is_admin, username").
		Where("is_active = ? AND username = ?", true, username).
		First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &auth_dto.AuthenticatedUserDto{
		Id:        int(user.ID),
		FirstName: user.FirstName,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
	}, nil
}

func (r *authRepository) UpdatePassword(id int, value string) error {
	result := r.db.
		Model(&schema.User{}).
		Where("id=?", id).
		Update("password", value)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
