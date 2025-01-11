package user_repository_proxy

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
)

type authRepositoryProxy struct {
	authRepository auth_contract.AuthRepository
	settings       settings.ApplicationSettings
	hasher         security.Hasher
}

func New(authRepository auth_contract.AuthRepository, settings settings.ApplicationSettings, hasher security.Hasher) auth_contract.AuthRepository {
	return &authRepositoryProxy{
		authRepository: authRepository,
		settings:       settings,
		hasher:         hasher,
	}
}

func (r *authRepositoryProxy) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	usernameHashed, err := r.hasher.HashWithSalt(username, r.settings.GetServerSalt())
	if err != nil {
		return nil, err
	}

	authenticatedUser, errRepo := r.authRepository.GetAuthenticatedUserByUsername(usernameHashed)
	if errRepo != nil {
		return nil, errRepo
	}

	return &auth_dto.AuthenticatedUserDto{
		Id:        authenticatedUser.Id,
		FirstName: r.hasher.DecryptMerged(authenticatedUser.FirstName, r.settings.GetSecretKey()),
		IsAdmin:   authenticatedUser.IsAdmin,
		Password:  authenticatedUser.Password,
	}, nil
}

func (r *authRepositoryProxy) UpdatePassword(id, value string) error {
	passwordHashed := r.hasher.HashAndMerge(value, r.settings.GetServerSalt())

	return r.authRepository.UpdatePassword(id, passwordHashed)
}
