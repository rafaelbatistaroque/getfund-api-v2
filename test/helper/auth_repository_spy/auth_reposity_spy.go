package auth_repository_spy

import (
	"errors"
	model "getfund-api-v2/internal/domain/auth/core/auth_dto"
)

type AuthRepositorySpy struct {
	Params        map[string]string
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *AuthRepositorySpy {

	return &AuthRepositorySpy{Params: make(map[string]string, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{}, 1), CallsCount: make(map[string]int, 1)}
}

func (r *AuthRepositorySpy) GetAuthenticatedUserByUsername(username string) (*model.AuthenticatedUserDto, error) {
	r.Params["GetAuthenticatedUserByUsername:username"] = username

	r.CallsCount["GetAuthenticatedUserByUsername"]++

	sucess := r.SuccessResult["GetAuthenticatedUserByUsername"]
	if sucess != nil {
		return sucess.(*model.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
	}
	r.DefineGetAuthenticatedUserByUsernameSuccess()
	return r.SuccessResult["GetAuthenticatedUserByUsername"].(*model.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
}

func (r *AuthRepositorySpy) UpdatePassword(id, value string) error {
	r.Params["UpdatePassword:id"] = id
	r.Params["UpdatePassword:value"] = value

	r.CallsCount["UpdatePassword"]++

	return r.ErrorResult["UpdatePassword"]
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameError() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("fake-error")
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameSuccess() {
	r.SuccessResult["GetAuthenticatedUserByUsername"] = &model.AuthenticatedUserDto{Password: "fake-password-hashed", FirstName: "fake-username", Id: "fake-id", IsAdmin: 0}
}

func (r *AuthRepositorySpy) DefineUpdatePasswordError() {
	r.ErrorResult["UpdatePassword"] = errors.New("fake-error")
}
