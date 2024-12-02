package authacomposer

import (
	authService "getfund-api-v2/internal/domain/auth/domainservice/authservice"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	userRepository "getfund-api-v2/internal/domain/auth/port/repository"
	signinApplication "getfund-api-v2/internal/domain/auth/usecase/signin/application"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/proxy"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	sessionService "getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"

	"gorm.io/gorm"
)

type AuthComposer struct {
	Signin http.HandlerFunc
}

func GetHandlers(settings settings.ApplicationSettings, cache cacheservice.Cache, sessionServive sessionService.SessionService, db *gorm.DB) AuthComposer {
	//dependencies
	hasher := security.New()
	mapper := mapper.New(hasher, settings)
	userRepository := userRepository.New(db)
	authService := authService.New(
		userRepository,
		settings,
		hasher,
		mapper)

	//applications
	signin := signinApplication.New(authService, sessionServive, mapper)

	//composer
	composer := adapter.New(signin)

	return AuthComposer{
		Signin: proxy.New(composer.Signin),
	}
}
