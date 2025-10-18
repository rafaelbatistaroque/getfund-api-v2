package reset_password_gateway

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
)

const (
	_KEY_CACHE_PREFIX = "recovery_password_"
)

type ResetPasswordGateway interface {
	ResetPassword(w http.ResponseWriter, r *http.Request) (any, int, error)
}

type resetPasswordGateway struct {
	resetPassword reset_password.UseCase
}

func New(resetPassword reset_password.UseCase) ResetPasswordGateway {
	return &resetPasswordGateway{
		resetPassword: resetPassword,
	}
}

func (h *resetPasswordGateway) ResetPassword(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var input reset_password.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, shared_error.BAD_REQUEST_CODE, err
	}

	input.RecoveryKey = _KEY_CACHE_PREFIX + input.RecoveryCode
	result, err := h.resetPassword.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return result, shared_error.SUCCESS_CODE, nil
}
