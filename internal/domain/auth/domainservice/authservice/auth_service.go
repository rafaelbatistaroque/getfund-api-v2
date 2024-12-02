package authservice

import (
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	"getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	authmodel "getfund-api-v2/internal/domain/auth/model"
	appCode "getfund-api-v2/internal/shared/applicationcode"
	appErr "getfund-api-v2/internal/shared/applicationerror"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
)

type AuthService interface {
	Authenticate(username string, password string) (*authmodel.SessionModel, *appErr.ApplicationError)
}

type authService struct {
	settings       settings.ApplicationSettings
	userRepository auth_contract.UserRepository
	hasher         security.Hasher
	mapper         signinmapper.SigninMapper
}

func New(
	userRepository auth_contract.UserRepository,
	settings settings.ApplicationSettings,
	hasher security.Hasher,
	mapper signinmapper.SigninMapper) AuthService {

	return &authService{
		userRepository: userRepository,
		settings:       settings,
		hasher:         hasher,
		mapper:         mapper,
	}
}

func (a *authService) Authenticate(username string, password string) (*authmodel.SessionModel, *appErr.ApplicationError) {
	usernameHashed, err := a.hasher.HashWithSalt(username, a.settings.GetServerSalt())
	if err != nil {
		return nil, appErr.New(appCode.CODE_SERVER_ERROR, err)
	}

	user, repoErr := a.userRepository.GetByUserName(usernameHashed)
	if repoErr != nil {
		return nil, appErr.New(appCode.CODE_UNAUTHORIZED, repoErr)
	}

	if !a.hasher.IsMatch(user.Password, password, a.settings.GetServerSalt()) {
		return nil, appErr.New(appCode.CODE_UNAUTHORIZED, errors.New("wrong password"))
	}

	return a.mapper.ToSessionModel(user), nil
}
