package authadapter

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"net/http"
)

func (h *authAdapter) RecoverPassword(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input recoverpassword.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	h.recoverPassword.Execute(&input)

	return nil, 0, nil
}
