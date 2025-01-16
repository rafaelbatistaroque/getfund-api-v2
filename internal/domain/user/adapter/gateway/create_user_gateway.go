package user_gateway

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/user/core/usercase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

func (u *userGateway) CreateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var input create_user.Input

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, result_app.BAD_REQUEST_CODE, err
	}

	u.createUser.Execute(&input)

	return nil, 0, nil
}
