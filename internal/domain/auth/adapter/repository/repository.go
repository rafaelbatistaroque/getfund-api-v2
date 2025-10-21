package repository

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"

	"gorm.io/gorm"
)

type authRepository struct {
	db *db.GetFund
}

func New(db *db.GetFund) contract.Repository {
	return &authRepository{db: db}
}

func (r *authRepository) GetAuthenticatedUserByUsername(username string) (*dto.AuthenticatedUserDto, error) {
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

	return &dto.AuthenticatedUserDto{
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

func (u *authRepository) Signup(auDto *dto.ActivationUserDto) (*dto.UserDto, error) {
	var user = schema.User{
		FirstName: auDto.FirstName,
		LastName:  auDto.LastName,
		Username:  auDto.Username,
		Password:  auDto.Password,
		IsAdmin:   auDto.IsAdmin,
		IsActive:  auDto.IsActive,
		CreatedAt: auDto.CreatedAt,
		UpdatedAt: auDto.UpdatedAt,
	}

	result := u.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.UserDto{
		Id: int(user.ID),
	}, nil
}

func (u *authRepository) UserExists(username string) (*dto.UserDto, error) {
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

	return &dto.UserDto{
		Id: int(user.ID),
	}, nil
}

func (r *authRepository) UpdateUsernameHash(id int, username string) error {
	result := r.db.
		Model(&schema.User{}).
		Where("id=?", id).
		Update("username", username)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
