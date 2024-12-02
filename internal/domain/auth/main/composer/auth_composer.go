package authacomposer

import (
	authService "getfund-api-v2/internal/domain/auth/domainservice/authservice"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	userRepository "getfund-api-v2/internal/domain/auth/port/repository"
	app "getfund-api-v2/internal/domain/auth/usecase/signin/application"
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
	hasher := security.New()
	mapper := mapper.New(hasher, settings)

	composer := adapter.New(
		app.NewUseCase(
			authService.New(
				userRepository.New(db),
				settings,
				hasher,
				mapper),
			sessionServive,
			mapper))

	return AuthComposer{
		Signin: proxy.New(composer.Signin),
	}
}
