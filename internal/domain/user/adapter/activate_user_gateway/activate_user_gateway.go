package activate_user_gateway

import (
	"errors"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"strings"
)

const (
	_PATH_ACTIVATE_USER = "/user/activate/"
)

type ActiveUserGateway interface {
	ActivateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type activeUserGateway struct {
	activateUser activate_user.UseCase
}

func New(activateUser activate_user.UseCase) ActiveUserGateway {
	return &activeUserGateway{
		activateUser: activateUser,
	}
}

func (u *activeUserGateway) ActivateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	activationCode := strings.TrimPrefix(r.URL.Path, _PATH_ACTIVATE_USER)

	if activationCode == "" {
		return nil, result_app.BAD_REQUEST_CODE, errors.New("activation code is required")
	}

	input := activate_user.Input{ActivationCode: activationCode}

	_, err := u.activateUser.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return nil, 0, nil
}
