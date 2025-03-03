package auth_repository_proxy

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
)

const _DEFAULT_ERROR = "error on get authenticated user"

type authRepositoryProxy struct {
	repository auth_contract.Repository
	settings   settings.ApplicationSettings
	hasher     security.Hasher
}

func New(repository auth_contract.Repository, settings settings.ApplicationSettings, hasher security.Hasher) auth_contract.Repository {
	return &authRepositoryProxy{
		repository: repository,
		settings:   settings,
		hasher:     hasher,
	}
}

func (p *authRepositoryProxy) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	usernameHashed, err := p.hasher.HashWithSalt(username, p.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	if usernameHashed == "" {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	authenticatedUser, errRepo := p.repository.GetAuthenticatedUserByUsername(usernameHashed)
	if errRepo != nil {
		return nil, errRepo
	}

	if authenticatedUser == nil {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	return &auth_dto.AuthenticatedUserDto{
		Id:        authenticatedUser.Id,
		FirstName: p.hasher.DecryptMerged(authenticatedUser.FirstName, p.settings.GetSecretKey()),
		IsAdmin:   authenticatedUser.IsAdmin,
		Password:  authenticatedUser.Password,
	}, nil
}

func (p *authRepositoryProxy) UpdatePassword(id int, value string) error {
	passwordHashed := p.hasher.HashAndMerge(value, p.settings.GetServerSalt())

	return p.repository.UpdatePassword(id, passwordHashed)
}

func (p *authRepositoryProxy) CreateUser(user *auth_dto.ActivationUserDto) (*auth_dto.UserDto, error) {
	userHashed := *user

	userHashed.FirstName = p.hasher.Encrypt(user.FirstName, p.settings.GetSecretKey())
	userHashed.LastName = p.hasher.Encrypt(user.LastName, p.settings.GetSecretKey())
	userHashed.Password = p.hasher.HashAndMerge(user.Password, p.settings.GetServerSalt())

	var err error
	userHashed.Username, err = p.hasher.HashWithSalt(user.Username, p.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	userCreated, err := p.repository.CreateUser(&userHashed)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

func (p *authRepositoryProxy) UserExists(username string) (*auth_dto.UserDto, error) {
	usernameHashed, err := p.hasher.HashWithSalt(username, p.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	if usernameHashed == "" {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	userFound, errRepo := p.repository.UserExists(usernameHashed)
	if errRepo != nil {
		return nil, errRepo
	}

	if userFound != nil {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	return nil, nil
}
