package auth_service

import (
	"errors"
	"getfund-api-v2/internal/config/env"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	"getfund-api-v2/internal/domain/auth/core/dto"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/security"
)

type AuthService interface {
	Authenticate(username string, password string) (*dto.SessionDto, *shared_error.Error)
}

type authService struct {
	env            env.Variable
	authRepository auth_contract.Repository
	hasher         security.Hasher
	mapper         signin_mapper.SigninMapper
}

func New(
	authRepository auth_contract.Repository,
	env env.Variable,
	hasher security.Hasher,
	mapper signin_mapper.SigninMapper) AuthService {

	return &authService{
		authRepository: authRepository,
		env:            env,
		hasher:         hasher,
		mapper:         mapper,
	}
}

func (a *authService) Authenticate(username string, password string) (*dto.SessionDto, *shared_error.Error) {
	authenticatedUser, repoErr := a.authRepository.GetAuthenticatedUserByUsername(username)
	if repoErr != nil {
		return nil, shared_error.New(shared_error.UNAUTHORIZED_CODE, repoErr)
	}

	if !a.hasher.IsMatch(authenticatedUser.Password, password, a.env.GetServerSalt()) {
		return nil, shared_error.New(shared_error.UNAUTHORIZED_CODE, errors.New("invalid password"))
	}

	return a.mapper.ToSessionModel(authenticatedUser), nil
}
