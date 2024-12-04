package authadapter

import (
	"encoding/json"
	signin "getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
)

func (h *authAdapter) Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input signin.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, resultapp.BAD_REQUEST_CODE, err
	}

	output, err := h.signin.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, resultapp.SUCCESS_CODE, nil
}
