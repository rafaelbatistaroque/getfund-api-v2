package auth_parser

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

func (h *authParser) ResetPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input reset_password.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, result_app.BAD_REQUEST_CODE, err
	}

	result, err := h.resetPassword.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return result, result_app.SUCCESS_CODE, nil
}
