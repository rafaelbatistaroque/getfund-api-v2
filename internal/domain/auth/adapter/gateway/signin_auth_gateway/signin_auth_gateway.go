package signin_gateway

import (
	"encoding/json"
	signin "getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
)

type SigninGateway interface {
	Signin(w http.ResponseWriter, r *http.Request) (any, int, error)
}

type signinGateway struct {
	signin signin.UseCase
}

func New(signin signin.UseCase) SigninGateway {
	return &signinGateway{
		signin: signin,
	}
}

func (h *signinGateway) Signin(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var input signin.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, shared_error.BAD_REQUEST_CODE, err
	}

	output, err := h.signin.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, shared_error.SUCCESS_CODE, nil
}
