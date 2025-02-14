package user_repository_proxy

import (
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/dto/user_dto"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
)

type userRepositoryProxy struct {
	repository user_contract.Repository
	settings   settings.ApplicationSettings
	hasher     security.Hasher
}

func New(repository user_contract.Repository, settings settings.ApplicationSettings, hasher security.Hasher) user_contract.Repository {
	return &userRepositoryProxy{
		repository: repository,
		settings:   settings,
		hasher:     hasher,
	}
}

func (u *userRepositoryProxy) CreateUser(user *user_dto.ActivationUserDto) (*user_dto.UserDto, error) {
	userHashed := *user

	userHashed.FirstName = u.hasher.Encrypt(user.FirstName, u.settings.GetSecretKey())
	userHashed.LastName = u.hasher.Encrypt(user.LastName, u.settings.GetSecretKey())
	userHashed.Email = u.hasher.Encrypt(user.Email, u.settings.GetSecretKey())
	userHashed.MainSocialNetwork = u.hasher.Encrypt(user.MainSocialNetwork, u.settings.GetSecretKey())
	userHashed.RegisteredUrl = u.hasher.Encrypt(user.RegisteredUrl, u.settings.GetSecretKey())
	userHashed.Password = u.hasher.HashAndMerge(user.Password, u.settings.GetServerSalt())

	var err error
	userHashed.Username, err = u.hasher.HashWithSalt(user.Username, u.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	userCreated, err := u.repository.CreateUser(&userHashed)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

func (u *userRepositoryProxy) UserExistsByUsername(username string) (*user_dto.UserDto, error) {
	usernameHashed, err := u.hasher.HashWithSalt(username, u.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	existingUser, errRepo := u.repository.UserExistsByUsername(usernameHashed)
	if errRepo != nil {
		return nil, errRepo
	}

	return existingUser, nil
}
