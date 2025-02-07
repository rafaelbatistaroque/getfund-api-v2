package user_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/user/core/user_dto"
)

type UserRepositorySpy struct {
	Params        map[string]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *UserRepositorySpy {

	return &UserRepositorySpy{Params: make(map[string]any, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]any, 1), CallsCount: make(map[string]int, 1)}
}

func (r *UserRepositorySpy) GetUserByUsername(username string) (*user_dto.UserDto, error) {
	r.Params["GetUserByUsername:username"] = username

	r.CallsCount["GetUserByUsername"]++

	sucess := r.SuccessResult["GetUserByUsername"]
	if sucess != nil {
		return sucess.(*user_dto.UserDto), nil
	}

	return nil, r.ErrorResult["GetUserByUsername"]
}

func (r *UserRepositorySpy) SaveUser(user *user_dto.ActivationUserDto) error {
	r.Params["SaveUser:user"] = user

	r.CallsCount["SaveUser"]++

	return r.ErrorResult["SaveUser"]
}

func (r *UserRepositorySpy) DefineGetUserByUsernameError() {
	r.ErrorResult["GetUserByUsername"] = errors.New("fake-error")
}

func (r *UserRepositorySpy) DefineGetUserByUsernameSuccess() {
	r.SuccessResult["GetUserByUsername"] = &user_dto.UserDto{Id: "fake-id"}
}

func (r *UserRepositorySpy) DefineGetUserByUsernameSuccessNotFound() {
	r.SuccessResult["GetUserByUsername"] = nil
}

func (r *UserRepositorySpy) DefineSaveUserError() {
	r.ErrorResult["SaveUser"] = errors.New("fake-error")
}
