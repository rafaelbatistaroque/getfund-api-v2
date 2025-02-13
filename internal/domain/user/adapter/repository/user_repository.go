package user_repository

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/pkg/db/schema"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user_contract.Repository {
	return &userRepository{db: db}
}

func (u *userRepository) UserExistsByUsername(username string) (*user_dto.UserDto, error) {
	var user = schema.User{}
	result := u.db.
		Select("id").
		Where("username = ?", username).
		First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user_dto.UserDto{
		Id: int(user.ID),
	}, nil
}

func (u *userRepository) SaveUser(user *user_dto.ActivationUserDto) (*user_dto.UserDto, error) {
	return nil, nil
}
