package auth_gateway

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	signin "getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"net/http"
)

type AuthGateway interface {
	Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	RecoverPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	ResetPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type authGateway struct {
	signin          signin.UseCase
	signout         signout.UseCase
	recoverPassword recover_password.UseCase
	resetPassword   reset_password.UseCase
}

func New(
	signin signin.UseCase,
	signout signout.UseCase,
	recoverPassword recover_password.UseCase,
	resetPassword reset_password.UseCase) AuthGateway {

	return &authGateway{
		signin:          signin,
		signout:         signout,
		recoverPassword: recoverPassword,
		resetPassword:   resetPassword,
	}
}
