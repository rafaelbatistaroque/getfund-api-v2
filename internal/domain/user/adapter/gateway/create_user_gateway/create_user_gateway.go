package create_user_gateway

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

type CreateUserGateway interface {
	CreateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type createUserGateway struct {
	createUser create_user.UseCase
}

func New(createUser create_user.UseCase) CreateUserGateway {
	return &createUserGateway{
		createUser: createUser,
	}
}

func (u *createUserGateway) CreateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input create_user.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, result_app.BAD_REQUEST_CODE, err
	}

	output, err := u.createUser.Execute(&input)
	if err != nil {
		return nil, err.Code, err.Message
	}

	return output, result_app.SUCCESS_CODE, nil
}
