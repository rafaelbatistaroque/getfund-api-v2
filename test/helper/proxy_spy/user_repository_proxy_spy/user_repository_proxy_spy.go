package user_repository_proxy_spy

import (
	"errors"
	auth_model "getfund-api-v2/internal/domain/auth/core/model"
)

type UserRepositoryProxySpy struct {
	Params        map[string]string
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *UserRepositoryProxySpy {

	return &UserRepositoryProxySpy{Params: make(map[string]string, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{}, 1), CallsCount: make(map[string]int, 1)}
}

func (r *UserRepositoryProxySpy) GetByUserName(username string) (*auth_model.UserModel, error) {
	r.Params["GetByUserName:username"] = username

	r.CallsCount["GetByUserName"]++

	sucess := r.SuccessResult["GetByUserName"]
	if sucess != nil {
		return sucess.(*auth_model.UserModel), r.ErrorResult["GetByUserName"]
	}
	r.DefineGetByUserNameSuccess()
	return r.SuccessResult["GetByUserName"].(*auth_model.UserModel), r.ErrorResult["GetByUserName"]
}

func (r *UserRepositoryProxySpy) DefineGetByUserNameError() {
	r.ErrorResult["GetByUserName"] = errors.New("fake-error")
}

func (r *UserRepositoryProxySpy) DefineGetByUserNameSuccess() {
	r.SuccessResult["GetByUserName"] = &auth_model.UserModel{Password: "fake-password-hashed", FirstName: "fake-username", Id: "fake-id", IsAdmin: 0}
}
