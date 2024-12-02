package authadapter

import (
	signin "getfund-api-v2/internal/domain/auth/usecase/signin"
	"net/http"
)

type AuthAdapter interface {
	Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
	// Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type authAdapter struct {
	signin signin.UseCase
}

func New(usecase signin.UseCase) AuthAdapter {
	return &authAdapter{
		signin: usecase,
	}
}
