package authadapter

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
)

func (h *authAdapter) RecoverPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input recoverpassword.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, resultapp.BAD_REQUEST, err
	}

	_, err := h.recoverPassword.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return nil, 0, nil
}
