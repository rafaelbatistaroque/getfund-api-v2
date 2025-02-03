package signout_gateway

import (
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

type SignoutGateway interface {
	Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type signoutGateway struct {
	signout signout.UseCase
}

func New(signout signout.UseCase) SignoutGateway {
	return &signoutGateway{
		signout: signout,
	}
}

func (h *signoutGateway) Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token := r.Context().Value(auth_contract.TokenKey{})
	if token == nil || token == "" {
		return nil, result_app.UNAUTHORIZED_CODE, errors.New("token not found")
	}

	input := &signout.Input{Token: token.(string)}
	output, err := h.signout.Execute(input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, result_app.SUCCESS_CODE, nil
}
