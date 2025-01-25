package user_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/user/core/user_dto"
)

type UserRepositorySpy struct {
	Params        map[string]string
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *UserRepositorySpy {

	return &UserRepositorySpy{Params: make(map[string]string, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{}, 1), CallsCount: make(map[string]int, 1)}
}

func (r *UserRepositorySpy) GetUserByUsername(username string) (*user_dto.UserDto, error) {
	r.Params["GetUserByUsername:username"] = username

	r.CallsCount["GetUserByUsername"]++

	sucess := r.SuccessResult["GetUserByUsername"]
	if sucess != nil {
		return sucess.(*user_dto.UserDto), r.ErrorResult["GetUserByUsername"]
	}
	r.DefineGetUserByUsernameSuccess()
	return r.SuccessResult["GetUserByUsername"].(*user_dto.UserDto), r.ErrorResult["GetUserByUsername"]
}

func (r *UserRepositorySpy) DefineGetUserByUsernameError() {
	r.ErrorResult["GetUserByUsername"] = errors.New("fake-error")
}

func (r *UserRepositorySpy) DefineGetUserByUsernameSuccess() {
	r.SuccessResult["GetUserByUsername"] = &user_dto.UserDto{Id: "fake-id"}
}
