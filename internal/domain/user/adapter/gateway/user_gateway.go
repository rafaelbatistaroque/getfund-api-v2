package user_gateway

import (
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"net/http"
)

type UserGateway interface {
	CreateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error)
}

type userGateway struct {
	createUser create_user.UseCase
}

func New(createUser create_user.UseCase) UserGateway {
	return &userGateway{
		createUser: createUser,
	}
}
