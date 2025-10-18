package auth_composer

import (
	"getfund-api-v2/internal/config/env"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/activate_user_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/recover_password_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	signin_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signin_auth_gateway"
	signout_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signout_auth_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/signup_gateway"
	"getfund-api-v2/internal/domain/auth/adapter/middleware/auth_middleware"
	"getfund-api-v2/internal/domain/auth/adapter/proxy/auth_repository_proxy"
	authRepository "getfund-api-v2/internal/domain/auth/adapter/repository"
	"getfund-api-v2/internal/domain/auth/core/domain_service/activate_user_mapper"
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/session_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	activate_user_application "getfund-api-v2/internal/domain/auth/core/usecase/activate_user/application"
	recover_password_application "getfund-api-v2/internal/domain/auth/core/usecase/recover_password/application"
	reset_password_application "getfund-api-v2/internal/domain/auth/core/usecase/reset_password/application"
	signin_application "getfund-api-v2/internal/domain/auth/core/usecase/signin/application"
	signout_application "getfund-api-v2/internal/domain/auth/core/usecase/signout/application"
	signup_application "getfund-api-v2/internal/domain/auth/core/usecase/signup/application"
	"getfund-api-v2/internal/infra/db"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	shared_response_proxy "getfund-api-v2/internal/shared/proxy"
	"getfund-api-v2/internal/shared/security"
	"net/http"
)

type middlewareFunc = func(http.Handler) http.Handler

type authComposer struct {
	Signin          http.HandlerFunc
	Signout         http.HandlerFunc
	RecoverPassword http.HandlerFunc
	ResetPassword   http.HandlerFunc
	Signup          http.HandlerFunc
	ActivateUser    http.HandlerFunc

	MiddlewareAutenticate      middlewareFunc
	MiddlewareAutenticateAdmin middlewareFunc
}

func Compose(
	env env.Variable,
	cache cache.Contract,
	db *db.GetFund,
	eventBus shared_bus.EventBus) authComposer {

	//dependencies
	hasher := security.New()
	signin_mapper := signin_mapper.New()
	activate_user_mapper := activate_user_mapper.New()
	repositoryProxy := auth_repository_proxy.New(authRepository.New(db), env, hasher)
	authService := auth_service.New(repositoryProxy, env, hasher, signin_mapper)
	sessionService := session_service.New(cache, hasher, env)

	//applications
	signin := signin_application.New(authService, sessionService, signin_mapper)
	signout := signout_application.New(sessionService)
	recoverPassword := recover_password_application.New(hasher, env, repositoryProxy, cache, eventBus)
	resetPassword := reset_password_application.New(cache, repositoryProxy)
	signupApplication := signup_application.New(repositoryProxy, hasher, cache, eventBus, env)
	activateUserApplication := activate_user_application.New(cache, repositoryProxy, activate_user_mapper, eventBus, env)

	//gateways
	signinGateway := signin_gateway.New(signin)
	signoutGateway := signout_gateway.New(signout)
	resetPasswordGateway := reset_password_gateway.New(resetPassword)
	recoverPasswordGateway := recover_password_gateway.New(recoverPassword)
	signupGateway := signup_gateway.New(signupApplication)
	activateUserGateway := activate_user_gateway.New(activateUserApplication, signin)

	//Middlewares
	authMiddleware := auth_middleware.New(sessionService)

	//Event Handler

	return authComposer{
		Signin:          shared_response_proxy.New(signinGateway.Signin),
		Signout:         shared_response_proxy.New(signoutGateway.Signout),
		RecoverPassword: shared_response_proxy.New(recoverPasswordGateway.RecoverPassword),
		ResetPassword:   shared_response_proxy.New(resetPasswordGateway.ResetPassword),
		ActivateUser:    shared_response_proxy.New(activateUserGateway.ActivateUser),
		Signup:          shared_response_proxy.New(signupGateway.Signup),

		MiddlewareAutenticate:      authMiddleware.Authenticate,
		MiddlewareAutenticateAdmin: authMiddleware.AuthenticateAdmin,
	}
}
