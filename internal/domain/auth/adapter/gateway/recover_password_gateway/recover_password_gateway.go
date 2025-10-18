package recover_password_gateway

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
)

type RecoverPasswordGateway interface {
	RecoverPassword(w http.ResponseWriter, r *http.Request) (any, int, error)
}

type recoverPasswordGateway struct {
	recoverPassword recover_password.UseCase
}

func New(recoverPassword recover_password.UseCase) RecoverPasswordGateway {
	return &recoverPasswordGateway{
		recoverPassword: recoverPassword,
	}
}

func (h *recoverPasswordGateway) RecoverPassword(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var input recover_password.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, shared_error.BAD_REQUEST_CODE, err
	}

	result, err := h.recoverPassword.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return result, shared_error.SUCCESS_CODE, nil
}
