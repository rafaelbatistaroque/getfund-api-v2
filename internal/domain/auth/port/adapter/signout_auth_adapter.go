package authadapter

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
)

func (h *authAdapter) Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input signout.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.signout.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, resultapp.CODE_SUCCESS, nil
}
