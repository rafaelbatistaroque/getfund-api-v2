package auth_entry_point_composer

import (
	"getfund-api-v2/internal/domain/auth/adapter/gateway/activate_user_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/create_user_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/recover_password_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	signin_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signin_auth_gateway"
	signout_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signout_auth_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/middleware/auth_middleware"
	"getfund-api-v2/internal/domain/auth/adapter/proxy/auth_repository_proxy"
	authRepository "getfund-api-v2/internal/domain/auth/adapter/repository"
	"getfund-api-v2/internal/domain/auth/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/session_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	activate_user_application "getfund-api-v2/internal/domain/auth/core/usecase/activate_user/application"
	create_user_application "getfund-api-v2/internal/domain/auth/core/usecase/create_user/application"
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

type authEntryPointComposer struct {
	Signin          http.HandlerFunc
	Signout         http.HandlerFunc
	RecoverPassword http.HandlerFunc
	ResetPassword   http.HandlerFunc
	CreateUser      http.HandlerFunc
	ActivateUser    http.HandlerFunc

	MiddlewareAutenticate      middlewareFunc
	MiddlewareAutenticateAdmin middlewareFunc
}

func Get(
	settings settings.ApplicationSettings,
	cache cache_service.Cache,
	db *gorm.DB,
	eventBus bus.EventBus) authEntryPointComposer {

	//dependencies
	hasher := security.New()
	signin_mapper := signin_mapper.New()
	activate_user_mapper := activate_user_mapper.New()
	repositoryProxy := auth_repository_proxy.New(authRepository.New(db), settings, hasher)
	authService := auth_service.New(repositoryProxy, settings, hasher, signin_mapper)
	sessionService := session_service.New(cache, hasher, settings)

	//applications
	signin := signin_application.New(authService, sessionService, signin_mapper)
	signout := signout_application.New(sessionService)
	recoverPassword := recover_password_application.New(hasher, settings, repositoryProxy, cache, eventBus)
	resetPassword := reset_password_application.New(cache, repositoryProxy)
	create_user_application := create_user_application.New(repositoryProxy, hasher, cache, eventBus, settings)
	activate_user_application := activate_user_application.New(cache, repositoryProxy, activate_user_mapper, eventBus, settings)

	//gateways
	signinGateway := signin_gateway.New(signin)
	signoutGateway := signout_gateway.New(signout)
	resetPasswordGateway := reset_password_gateway.New(resetPassword)
	recoverPasswordGateway := recover_password_gateway.New(recoverPassword)
	create_user_gateway := create_user_gateway.New(create_user_application)
	activate_user_gateway := activate_user_gateway.New(activate_user_application, signin)

	//Middlewares
	auth_middleware := auth_middleware.New(sessionService)

	//Event Handler

	return authEntryPointComposer{
		Signin:          response_proxy.New(signinGateway.Signin),
		Signout:         response_proxy.New(signoutGateway.Signout),
		RecoverPassword: response_proxy.New(recoverPasswordGateway.RecoverPassword),
		ResetPassword:   response_proxy.New(resetPasswordGateway.ResetPassword),
		ActivateUser:    response_proxy.New(activate_user_gateway.ActivateUser),
		CreateUser:      response_proxy.New(create_user_gateway.CreateUser),

		MiddlewareAutenticate:      auth_middleware.Authenticate,
		MiddlewareAutenticateAdmin: auth_middleware.AuthenticateAdmin,
	}
}
