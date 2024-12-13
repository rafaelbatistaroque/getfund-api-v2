package auth_service

import (
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/adapter/contract"
	authmodel "getfund-api-v2/internal/domain/auth/adapter/model"
	"getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
)

type AuthService interface {
	Authenticate(username string, password string) (*authmodel.SessionModel, *result_app.ApplicationError)
}

type authService struct {
	settings       settings.ApplicationSettings
	userRepository auth_contract.UserRepository
	hasher         security.Hasher
	mapper         signin_mapper.SigninMapper
}

func New(
	userRepository auth_contract.UserRepository,
	settings settings.ApplicationSettings,
	hasher security.Hasher,
	mapper signin_mapper.SigninMapper) AuthService {

	return &authService{
		userRepository: userRepository,
		settings:       settings,
		hasher:         hasher,
		mapper:         mapper,
	}
}

func (a *authService) Authenticate(username string, password string) (*authmodel.SessionModel, *result_app.ApplicationError) {
	usernameHashed, err := a.hasher.HashWithSalt(username, a.settings.GetServerSalt())
	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	user, repoErr := a.userRepository.GetByUserName(usernameHashed)
	if repoErr != nil {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, repoErr)
	}

	if !a.hasher.IsMatch(user.Password, password, a.settings.GetServerSalt()) {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, errors.New("invalid password"))
	}

	return a.mapper.ToSessionModel(user), nil
}
