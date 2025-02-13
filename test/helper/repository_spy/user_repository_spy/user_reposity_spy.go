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

func (r *UserRepositorySpy) UserExistsByUsername(username string) (*user_dto.UserDto, error) {
	r.Params["UserExistsByUsername:username"] = username

	r.CallsCount["UserExistsByUsername"]++

	sucess := r.SuccessResult["UserExistsByUsername"]
	if sucess != nil {
		return sucess.(*user_dto.UserDto), nil
	}

	return nil, r.ErrorResult["UserExistsByUsername"]
}

func (r *UserRepositorySpy) CreateUser(user *user_dto.ActivationUserDto) (*user_dto.UserDto, error) {
	r.Params["CreateUser:user"] = user

	r.CallsCount["CreateUser"]++

	sucess := r.SuccessResult["CreateUser"]
	if sucess != nil {
		return sucess.(*user_dto.UserDto), nil
	}

	return nil, r.ErrorResult["CreateUser"]
}

func (r *UserRepositorySpy) DefineUserExistsByUsernameError() {
	r.ErrorResult["UserExistsByUsername"] = errors.New("fake-error")
}

func (r *UserRepositorySpy) DefineUserExistsByUsernameSuccess() {
	r.SuccessResult["UserExistsByUsername"] = &user_dto.UserDto{Id: 1}
}

func (r *UserRepositorySpy) DefineUserExistsByUsernameSuccessNotFound() {
	r.SuccessResult["UserExistsByUsername"] = nil
}

func (r *UserRepositorySpy) DefineCreateUserError() {
	r.ErrorResult["CreateUser"] = errors.New("fake-error")
}

func (r *UserRepositorySpy) DefineCreateUserSuccess() {
	r.SuccessResult["CreateUser"] = &user_dto.UserDto{Id: 1}
}
