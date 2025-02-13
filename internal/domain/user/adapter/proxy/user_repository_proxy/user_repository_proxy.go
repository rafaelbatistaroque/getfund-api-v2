package user_repository_proxy

import (
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/user_dto"
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
	user.FirstName = u.hasher.Encrypt(user.FirstName, u.settings.GetSecretKey())
	user.LastName = u.hasher.Encrypt(user.LastName, u.settings.GetSecretKey())
	user.Email = u.hasher.Encrypt(user.Email, u.settings.GetSecretKey())
	user.MainSocialNetwork = u.hasher.Encrypt(user.MainSocialNetwork, u.settings.GetSecretKey())
	user.RegisteredUrl = u.hasher.Encrypt(user.RegisteredUrl, u.settings.GetSecretKey())
	user.Password = u.hasher.HashAndMerge(user.Password, u.settings.GetServerSalt())
	usernameHashed, errHash := u.hasher.HashWithSalt(user.Username, u.settings.GetServerSalt())
	if errHash != nil {
		return nil, errHash
	}

	user.Username = usernameHashed

	userCreated, errRepo := u.repository.CreateUser(user)
	if errRepo != nil {
		return nil, errRepo
	}

	return userCreated, nil
}

func (u *userRepositoryProxy) UserExistsByUsername(username string) (*user_dto.UserDto, error) {
	return nil, nil
}
