package auth_composer

import (
	"getfund-api-v2/internal/domain/auth/adapter/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password/recover_password_application"
	signin_application "getfund-api-v2/internal/domain/auth/adapter/usecase/signin/application"
	signout_application "getfund-api-v2/internal/domain/auth/adapter/usecase/signout/application"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
	parser "getfund-api-v2/internal/domain/auth/port/parser"
	userRepository "getfund-api-v2/internal/domain/auth/port/repository"
	"getfund-api-v2/internal/proxy"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/internal/shared/service/code_service"
	sessionService "getfund-api-v2/internal/shared/service/session_service"
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
	cache cache_service.Cache,
	sessionServive sessionService.SessionService,
	db *gorm.DB,
	eventBus eventbus.EventBus,
	codeService code_service.CodeService) AuthComposer {

	//dependencies
	hasher := security.New()
	mapper := mapper.New(hasher, settings)
	userRepository := userRepository.New(db)
	authService := auth_service.New(userRepository, settings, hasher, mapper)

	//applications
	signin := signin_application.New(authService, sessionServive, mapper)
	signout := signout_application.New(sessionServive)
	recoverPassword := recover_password_application.New(hasher, settings, userRepository, codeService, cache, eventBus)

	//parser
	parser := parser.New(signin, signout, recoverPassword)

	return AuthComposer{
		Signin:          proxy.New(parser.Signin),
		Signout:         proxy.New(parser.Signout),
		RecoverPassword: proxy.New(parser.RecoverPassword),
	}
}
