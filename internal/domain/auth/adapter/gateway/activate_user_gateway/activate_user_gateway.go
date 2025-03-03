package activate_user_gateway

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"strings"
)

const (
	_KEY_USER_ACTIVATION_PREFIX = "user_activation_"
)

type ActiveUserGateway interface {
	ActivateUser(w http.ResponseWriter, r *http.Request) (any, int, error)
}

type activeUserGateway struct {
	activateUser activate_user.UseCase
	signin       signin.UseCase
}

func New(activateUser activate_user.UseCase, signin signin.UseCase) ActiveUserGateway {
	return &activeUserGateway{
		activateUser: activateUser,
		signin:       signin,
	}
}

func (u *activeUserGateway) ActivateUser(w http.ResponseWriter, r *http.Request) (any, int, error) {
	activationCode := getActivationCodeParam(r)

	if activationCode == "" {
		return nil, result_app.BAD_REQUEST_CODE, errors.New("activation code is required")
	}

	input := activate_user.Input{
		ActivationCode:    activationCode,
		ActivationDataKey: _KEY_USER_ACTIVATION_PREFIX + activationCode}

	output, err := u.activateUser.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	if output == nil {
		return nil, result_app.SERVER_ERROR_CODE, errors.New("error to activate user")
	}

	outputSignin, errSignin := u.signin.Execute(&signin.Input{
		Username: output.Username,
		Password: output.Password,
	})

	if errSignin != nil {
		return nil, errSignin.Code, errSignin.Message
	}

	return outputSignin, result_app.SUCCESS_CODE, nil
}

func getActivationCodeParam(r *http.Request) string {
	urlParts := strings.Split(r.URL.Path, "/")

	return urlParts[len(urlParts)-1]
}
