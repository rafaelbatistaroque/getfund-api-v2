package authadapter

import (
	"encoding/json"
	signin "getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/shared/applicationcode"
	"net/http"
)

type SigninAdapter interface {
	Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type signinAdapter struct {
	signin signin.UseCase
}

func New(usecase signin.UseCase) SigninAdapter {
	return &signinAdapter{
		signin: usecase,
	}
}

func (h *signinAdapter) Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input signin.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.signin.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, applicationcode.CODE_SUCCESS, nil
}
