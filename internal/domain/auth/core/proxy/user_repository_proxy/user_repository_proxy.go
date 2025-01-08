package user_repository_proxy

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	auth_model "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
)

type userRepositoryProxy struct {
	userRepository auth_contract.UserRepository
	settings       settings.ApplicationSettings
	hasher         security.Hasher
}

func New(userRepository auth_contract.UserRepository, settings settings.ApplicationSettings, hasher security.Hasher) auth_contract.UserRepository {
	return &userRepositoryProxy{
		userRepository: userRepository,
		settings:       settings,
		hasher:         hasher,
	}
}

func (r *userRepositoryProxy) GetByUserName(username string) (*auth_model.UserModel, error) {
	r.hasher.HashWithSalt(username, r.settings.GetServerSalt())
	return nil, nil
}
