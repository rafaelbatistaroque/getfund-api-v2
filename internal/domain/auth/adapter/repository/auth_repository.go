package auth_repository

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"

	"gorm.io/gorm"
)

type authRepository struct {
	db *db.GetFund
}

func New(db *db.GetFund) auth_contract.Repository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	var user = schema.User{}
	result := r.db.
		Select("id, first_name, is_admin, username, password").
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

func (u *authRepository) Signup(dto *auth_dto.ActivationUserDto) (*auth_dto.UserDto, error) {
	var user = schema.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Username:  dto.Username,
		Password:  dto.Password,
		IsAdmin:   dto.IsAdmin,
		IsActive:  dto.IsActive,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}

	result := u.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &auth_dto.UserDto{
		Id: int(user.ID),
	}, nil
}

func (u *authRepository) UserExists(username string) (*auth_dto.UserDto, error) {
	var user = schema.User{}
	result := u.db.
		Select("id").
		Where("username = ?", username).
		First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &auth_dto.UserDto{
		Id: int(user.ID),
	}, nil
}
