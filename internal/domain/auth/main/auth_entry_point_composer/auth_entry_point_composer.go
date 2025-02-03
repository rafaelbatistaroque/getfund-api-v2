package auth_entry_point_composer

import (
	"getfund-api-v2/internal/domain/auth/adapter/gateway/recover_password_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	signin_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signin_auth_gateway"
	signout_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signout_auth_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/middleware/auth_middleware"
	"getfund-api-v2/internal/domain/auth/adapter/proxy/user_repository_proxy"
	authRepository "getfund-api-v2/internal/domain/auth/adapter/repository"
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/session_service"
	mapper "getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	recover_password_application "getfund-api-v2/internal/domain/auth/core/usecase/recover_password/application"
	reset_password_application "getfund-api-v2/internal/domain/auth/core/usecase/reset_password/application"
	signin_application "getfund-api-v2/internal/domain/auth/core/usecase/signin/application"
	signout_application "getfund-api-v2/internal/domain/auth/core/usecase/signout/application"
	"getfund-api-v2/internal/proxy/response_proxy"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"net/http"

	"gorm.io/gorm"
)

type middlewareFunc = func(http.Handler) http.Handler

type AuthEntryPointComposer struct {
	Signin                     http.HandlerFunc
	Signout                    http.HandlerFunc
	RecoverPassword            http.HandlerFunc
	ResetPassword              http.HandlerFunc
	MiddlewareAutenticate      middlewareFunc
	MiddlewareAutenticateAdmin middlewareFunc
}

func Get(
	settings settings.ApplicationSettings,
	cache cache_service.Cache,
	db *gorm.DB,
	eventBus bus.EventBus) AuthEntryPointComposer {

	//dependencies
	hasher := security.New()
	mapper := mapper.New()
	authRepositoryProxy := user_repository_proxy.New(authRepository.New(db), settings, hasher)
	authService := auth_service.New(authRepositoryProxy, settings, hasher, mapper)
	sessionService := session_service.New(cache, hasher, settings)

	//applications
	signin := signin_application.New(authService, sessionService, mapper)
	signout := signout_application.New(sessionService)
	recoverPassword := recover_password_application.New(hasher, settings, authRepositoryProxy, cache, eventBus)
	resetPassword := reset_password_application.New(cache, authRepositoryProxy)

	//gateways
	signinGateway := signin_gateway.New(signin)
	signoutGateway := signout_gateway.New(signout)
	resetPasswordGateway := reset_password_gateway.New(resetPassword)
	recoverPasswordGateway := recover_password_gateway.New(recoverPassword)

	//Middlewares
	auth_middleware := auth_middleware.New(sessionService)

	return AuthEntryPointComposer{
		Signin:                     response_proxy.New(signinGateway.Signin),
		Signout:                    response_proxy.New(signoutGateway.Signout),
		RecoverPassword:            response_proxy.New(recoverPasswordGateway.RecoverPassword),
		ResetPassword:              response_proxy.New(resetPasswordGateway.ResetPassword),
		MiddlewareAutenticate:      auth_middleware.Authenticate,
		MiddlewareAutenticateAdmin: auth_middleware.AuthenticateAdmin,
	}
}
