package user_repository

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/dto/user_dto"
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

func (u *userRepository) CreateUser(dto *user_dto.ActivationUserDto) (*user_dto.UserDto, error) {
	var user = schema.User{
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Email:             dto.Email,
		Username:          dto.Username,
		Password:          dto.Password,
		Gender:            dto.Gender,
		MainSocialNetwork: dto.MainSocialNetwork,
		RegisteredUrl:     dto.RegisteredUrl,
		IsAdmin:           dto.IsAdmin,
		IsActive:          dto.IsActive,
		UserCategoryID:    uint(dto.UserCategoryId),
		CountryID:         uint(dto.CountryId),
	}

	result := u.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user_dto.UserDto{
		Id: int(user.ID),
	}, nil
}
