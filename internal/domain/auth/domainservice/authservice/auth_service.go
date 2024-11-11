package authservice

import (
	userRepository "getfund-api-v2/internal/domain/auth/contract/auth_userrepository"
	entity "getfund-api-v2/internal/domain/auth/entity/sessionentity"
	appErr "getfund-api-v2/internal/pkg/applicationerror"
	"getfund-api-v2/internal/pkg/settings"
)

type AuthService interface {
	Authenticate(username string, password string) (entity.Session, *appErr.ApplicationError)
}

type authService struct {
	settings       settings.ApplicationSettings
	userRepository userRepository.UserRepository
}

func New(
	userRepository userRepository.UserRepository,
	settings settings.ApplicationSettings) AuthService {

	return &authService{
		userRepository: userRepository,
		settings:       settings,
	}
}

func (a *authService) Authenticate(username string, password string) (entity.Session, *appErr.ApplicationError) {

	// usernameHash, err := security.HashWithSalt(username, fmt.Sprint(len(username)), a.settings.ServerSalt)

	return nil, nil

}
