package authadapter

import (
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	signin "getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"net/http"
)

type AuthAdapter interface {
	Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	RecoverPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type authAdapter struct {
	signin          signin.UseCase
	signout         signout.UseCase
	recoverPassword recoverpassword.UseCase
}

func New(
	signin signin.UseCase,
	signout signout.UseCase,
	recoverPassword recoverpassword.UseCase) AuthAdapter {

	return &authAdapter{
		signin:          signin,
		signout:         signout,
		recoverPassword: recoverPassword,
	}
}
