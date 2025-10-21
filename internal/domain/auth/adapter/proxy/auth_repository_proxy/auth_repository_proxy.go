package auth_repository_proxy

import (
	"errors"
	"getfund-api-v2/internal/config/env"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/shared/security"
)

const _DEFAULT_ERROR = "error on get authenticated user"

type authRepositoryProxy struct {
	repository auth_contract.Repository
	env        env.Variable
	hasher     security.Hasher
}

func New(repository auth_contract.Repository, env env.Variable, hasher security.Hasher) auth_contract.Repository {
	return &authRepositoryProxy{
		repository: repository,
		env:        env,
		hasher:     hasher,
	}
}

func (p *authRepositoryProxy) GetAuthenticatedUserByUsername(username string) (*dto.AuthenticatedUserDto, error) {
	// 1. Tenta encontrar o usuário com a NOVA lógica de hash
	usernameHashed, err := p.hasher.HashWithSalt(username, p.env.GetServerSalt())
	if err != nil {
		return nil, err
	}

	if usernameHashed == "" {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	authenticatedUser, errRepo := p.repository.GetAuthenticatedUserByUsername(usernameHashed)

	// 2. Se não encontrou, tenta com a lógica ANTIGA (legada)
	if authenticatedUser == nil || errRepo != nil {
		usernameHashedLegacy, errLegacy := p.hasher.HashWithSaltLegacy(username, p.env.GetServerSalt())
		if errLegacy != nil {
			return nil, errors.New(_DEFAULT_ERROR)
		}

		authenticatedUser, errRepo = p.repository.GetAuthenticatedUserByUsername(usernameHashedLegacy)
		if errRepo != nil {
			return nil, errRepo
		}

		// Se encontrou com a lógica antiga, atualiza o hash no banco para a nova lógica (MIGRAÇÃO)
		if authenticatedUser != nil {
			errUpdate := p.repository.UpdateUsernameHash(authenticatedUser.Id, usernameHashed)
			if errUpdate != nil {
				return nil, errUpdate
			}
		}
	}

	// 3. Se depois de ambas as tentativas o usuário for nulo, retorna erro.
	if authenticatedUser == nil {
		return nil, errors.New(_DEFAULT_ERROR)
	}

	return &dto.AuthenticatedUserDto{
		Id:        authenticatedUser.Id,
		FirstName: p.hasher.DecryptMerged(authenticatedUser.FirstName, p.env.GetSecretKey()),
		IsAdmin:   authenticatedUser.IsAdmin,
		Password:  authenticatedUser.Password,
	}, nil
}

func (p *authRepositoryProxy) UpdatePassword(id int, value string) error {
	passwordHashed := p.hasher.HashAndMerge(value, p.env.GetServerSalt())

	return p.repository.UpdatePassword(id, passwordHashed)
}

func (p *authRepositoryProxy) UpdateUsernameHash(id int, username string) error {
	return p.repository.UpdateUsernameHash(id, username)
}

func (p *authRepositoryProxy) Signup(user *dto.ActivationUserDto) (*dto.UserDto, error) {
	userHashed := *user

	userHashed.FirstName = p.hasher.Encrypt(user.FirstName, p.env.GetSecretKey())
	userHashed.LastName = p.hasher.Encrypt(user.LastName, p.env.GetSecretKey())
	userHashed.Password = p.hasher.HashAndMerge(user.Password, p.env.GetServerSalt())

	var err error
	userHashed.Username, err = p.hasher.HashWithSalt(user.Username, p.env.GetServerSalt())
	if err != nil {
		return nil, err
	}

	userCreated, err := p.repository.Signup(&userHashed)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

func (p *authRepositoryProxy) UserExists(username string) (*dto.UserDto, error) {
	usernameHashed, err := p.hasher.HashWithSalt(username, p.env.GetServerSalt())
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
