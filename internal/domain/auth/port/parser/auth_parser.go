package auth_parser

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password"
	signin "getfund-api-v2/internal/domain/auth/adapter/usecase/signin"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signout"
	"net/http"
)

type AuthParser interface {
	Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	RecoverPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	ResetPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type authParser struct {
	signin          signin.UseCase
	signout         signout.UseCase
	recoverPassword recover_password.UseCase
	resetPassword   reset_password.UseCase
}

func New(
	signin signin.UseCase,
	signout signout.UseCase,
	recoverPassword recover_password.UseCase,
	resetPassword reset_password.UseCase) AuthParser {

	return &authParser{
		signin:          signin,
		signout:         signout,
		recoverPassword: recoverPassword,
		resetPassword:   resetPassword,
	}
}
