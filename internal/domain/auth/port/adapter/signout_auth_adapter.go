package authadapter

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"
)

func (h *authAdapter) Signout(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token := r.Context().Value(sessionservice.TokenKey{})
	if token == nil || token == "" {
		return nil, resultapp.UNAUTHORIZED_CODE, errors.New("token not found")
	}

	input := &signout.Input{Token: token.(string)}
	output, err := h.signout.Execute(input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, resultapp.SUCCESS_CODE, nil
}
