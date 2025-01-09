package auth_repository_proxy_spy

import (
	"errors"
	auth_model "getfund-api-v2/internal/domain/auth/core/auth_dto"
)

type AuthRepositoryProxySpy struct {
	Params        map[string]string
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *AuthRepositoryProxySpy {

	return &AuthRepositoryProxySpy{Params: make(map[string]string, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{}, 1), CallsCount: make(map[string]int, 1)}
}

func (r *AuthRepositoryProxySpy) GetAuthenticatedUserByUsername(username string) (*auth_model.AuthenticatedUserDto, error) {
	r.Params["GetAuthenticatedUserByUsername:username"] = username

	r.CallsCount["GetAuthenticatedUserByUsername"]++

	sucess := r.SuccessResult["GetAuthenticatedUserByUsername"]
	if sucess != nil {
		return sucess.(*auth_model.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
	}
	r.DefineGetAuthenticatedUserByUsernameSuccess()
	return r.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_model.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
}

func (r *AuthRepositoryProxySpy) UpdatePassword(id, value string) error {
	r.Params["UpdatePassword:id"] = id
	r.Params["UpdatePassword:value"] = value

	r.CallsCount["UpdatePassword"]++

	return r.ErrorResult["UpdatePassword"]
}

func (r *AuthRepositoryProxySpy) DefineGetAuthenticatedUserByUsernameError() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("fake-error")
}

func (r *AuthRepositoryProxySpy) DefineGetAuthenticatedUserByUsernameSuccess() {
	r.SuccessResult["GetAuthenticatedUserByUsername"] = &auth_model.AuthenticatedUserDto{Password: "fake-password-hashed", FirstName: "fake-username", Id: "fake-id", IsAdmin: 0}
}
