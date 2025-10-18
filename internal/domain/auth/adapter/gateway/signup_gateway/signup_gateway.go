package signup_gateway

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
)

type SignupGateway interface {
	Signup(w http.ResponseWriter, r *http.Request) (any, int, error)
}

type signupGateway struct {
	signup signup.UseCase
}

func New(Signup signup.UseCase) SignupGateway {
	return &signupGateway{
		signup: Signup,
	}
}

func (u *signupGateway) Signup(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var input signup.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, shared_error.BAD_REQUEST_CODE, err
	}

	output, err := u.signup.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, shared_error.SUCCESS_CODE, nil
}
