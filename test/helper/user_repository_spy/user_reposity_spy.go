package user_repository_spy

import (
	"errors"
	model "getfund-api-v2/internal/domain/auth/core/model"
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

func (r *UserRepositorySpy) GetByUserName(username string) (*model.UserModel, error) {
	r.Params["GetByUserName:username"] = username

	r.CallsCount["GetByUserName"]++

	sucess := r.SuccessResult["GetByUserName"]
	if sucess != nil {
		return sucess.(*model.UserModel), r.ErrorResult["GetByUserName"]
	}
	r.DefineGetByUserNameSuccess()
	return r.SuccessResult["GetByUserName"].(*model.UserModel), r.ErrorResult["GetByUserName"]
}

func (r *UserRepositorySpy) DefineGetByUserNameError() {
	r.ErrorResult["GetByUserName"] = errors.New("fake-error")
}

func (r *UserRepositorySpy) DefineGetByUserNameSuccess() {
	r.SuccessResult["GetByUserName"] = &model.UserModel{Password: "fake-password-hashed", FirstName: "fake-username", Id: "fake-id", IsAdmin: 0}
}
