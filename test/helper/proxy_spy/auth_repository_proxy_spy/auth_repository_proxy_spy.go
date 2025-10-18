package auth_repository_proxy_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
)

type AuthRepositoryProxySpy struct {
	Params        map[string]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *AuthRepositoryProxySpy {

	return &AuthRepositoryProxySpy{Params: make(map[string]any, 1), ErrorResult: make(map[string]error), SuccessResult: make(map[string]any, 1), CallsCount: make(map[string]int, 1)}
}

func (r *AuthRepositoryProxySpy) GetAuthenticatedUserByUsername(username string) (*auth_dto.AuthenticatedUserDto, error) {
	r.Params["GetAuthenticatedUserByUsername:username"] = username

	r.CallsCount["GetAuthenticatedUserByUsername"]++

	sucess := r.SuccessResult["GetAuthenticatedUserByUsername"]
	if sucess != nil {
		return sucess.(*auth_dto.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
	}
	r.DefineGetAuthenticatedUserByUsernameSuccess()
	return r.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_dto.AuthenticatedUserDto), r.ErrorResult["GetAuthenticatedUserByUsername"]
}

func (r *AuthRepositoryProxySpy) UpdatePassword(id int, value string) error {
	r.Params["UpdatePassword:id"] = id
	r.Params["UpdatePassword:value"] = value

	r.CallsCount["UpdatePassword"]++

	return r.ErrorResult["UpdatePassword"]
}

func (r *AuthRepositoryProxySpy) Signup(user *auth_dto.ActivationUserDto) (*auth_dto.UserDto, error) {
	r.Params["Signup:user"] = user

	r.CallsCount["Signup"]++

	sucess := r.SuccessResult["Signup"]
	if sucess != nil {
		return sucess.(*auth_dto.UserDto), nil
	}

	return nil, r.ErrorResult["Signup"]
}

func (r *AuthRepositoryProxySpy) UserExists(username string) (*auth_dto.UserDto, error) {
	r.Params["UserExists:username"] = username

	r.CallsCount["UserExists"]++

	sucess := r.SuccessResult["UserExists"]
	if sucess != nil {
		return sucess.(*auth_dto.UserDto), nil
	}

	return nil, r.ErrorResult["UserExists"]
}

func (r *AuthRepositoryProxySpy) DefineGetAuthenticatedUserByUsernameError() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("fake-error")
}

func (r *AuthRepositoryProxySpy) DefineGetAuthenticatedUserByUsernameSuccess() {
	r.SuccessResult["GetAuthenticatedUserByUsername"] = &auth_dto.AuthenticatedUserDto{Password: "fake-password-hashed", FirstName: "fake-username", Id: 1, IsAdmin: false}
}
