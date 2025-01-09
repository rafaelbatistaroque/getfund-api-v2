package auth_service

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
)

type AuthService interface {
	Authenticate(username string, password string) (*auth_dto.SessionDto, *result_app.ApplicationError)
}

type authService struct {
	settings       settings.ApplicationSettings
	authRepository auth_contract.AuthRepository
	hasher         security.Hasher
	mapper         signin_mapper.SigninMapper
}

func New(
	authRepository auth_contract.AuthRepository,
	settings settings.ApplicationSettings,
	hasher security.Hasher,
	mapper signin_mapper.SigninMapper) AuthService {

	return &authService{
		authRepository: authRepository,
		settings:       settings,
		hasher:         hasher,
		mapper:         mapper,
	}
}

func (a *authService) Authenticate(username string, password string) (*auth_dto.SessionDto, *result_app.ApplicationError) {
	authenticatedUser, repoErr := a.authRepository.GetAuthenticatedUserByUsername(username)
	if repoErr != nil {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, repoErr)
	}

	if !a.hasher.IsMatch(authenticatedUser.Password, password, a.settings.GetServerSalt()) {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, errors.New("invalid password"))
	}

	return a.mapper.ToSessionModel(authenticatedUser), nil
}
