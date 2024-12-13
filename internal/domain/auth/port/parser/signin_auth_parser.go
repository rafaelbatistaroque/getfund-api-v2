package auth_parser

import (
	"encoding/json"
	signin "getfund-api-v2/internal/domain/auth/adapter/usecase/signin"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

func (h *authParser) Signin(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input signin.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, result_app.BAD_REQUEST_CODE, err
	}

	output, err := h.signin.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, result_app.SUCCESS_CODE, nil
}
