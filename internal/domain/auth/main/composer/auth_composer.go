package authacomposer

import (
	authService "getfund-api-v2/internal/domain/auth/domainservice/authservice"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	userRepository "getfund-api-v2/internal/domain/auth/port/repository"
	recoverPasswordApplication "getfund-api-v2/internal/domain/auth/usecase/recoverpassword/application"
	signinApplication "getfund-api-v2/internal/domain/auth/usecase/signin/application"
	signoutApplication "getfund-api-v2/internal/domain/auth/usecase/signout/application"
	"getfund-api-v2/internal/proxy"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/internal/shared/service/codeservice"
	sessionService "getfund-api-v2/internal/shared/service/sessionservice"
	"getfund-api-v2/pkg/eventbus"
	"net/http"

	"gorm.io/gorm"
)

type AuthComposer struct {
	Signin          http.HandlerFunc
	Signout         http.HandlerFunc
	RecoverPassword http.HandlerFunc
}

func GetHandlers(
	settings settings.ApplicationSettings,
	cache cacheservice.Cache,
	sessionServive sessionService.SessionService,
	db *gorm.DB,
	eventBus eventbus.EventBus,
	codeService codeservice.CodeService) AuthComposer {

	//dependencies
	hasher := security.New()
	mapper := mapper.New(hasher, settings)
	userRepository := userRepository.New(db)
	authService := authService.New(userRepository, settings, hasher, mapper)

	//applications
	signin := signinApplication.New(authService, sessionServive, mapper)
	signout := signoutApplication.New(sessionServive)
	recoverPassword := recoverPasswordApplication.New(hasher, settings, userRepository, codeService, cache, eventBus)

	//composer
	composer := adapter.New(signin, signout, recoverPassword)

	return AuthComposer{
		Signin:          proxy.New(composer.Signin),
		Signout:         proxy.New(composer.Signout),
		RecoverPassword: proxy.New(composer.RecoverPassword),
	}
}
