package auth_gateway

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/session_service"
	"net/http"
)

func (h *authGateway) Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token := r.Context().Value(session_service.TokenKey{})
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
