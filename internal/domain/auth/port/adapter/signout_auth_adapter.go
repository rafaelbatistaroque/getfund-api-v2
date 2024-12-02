package authadapter

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"
)

func (h *authAdapter) Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input signout.Input

	session := r.Context().Value(sessionservice.SessionKey{}).(string)
	if session == "" {
		return nil, http.StatusInternalServerError, errors.New("session not found")
	}

	if err := json.Unmarshal([]byte(session), &input); err != nil {
		return nil, http.StatusBadRequest, err
	}

	output, err := h.signout.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, resultapp.CODE_SUCCESS, nil
}
