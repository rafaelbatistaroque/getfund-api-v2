package auth_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/dto"
)

type AuthRepositorySpy struct {
	Params        map[string]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *AuthRepositorySpy {

	return &AuthRepositorySpy{
		Params:        make(map[string]any, 1),
		ErrorResult:   make(map[string]error),
		SuccessResult: make(map[string]any, 1),
		CallsCount:    make(map[string]int, 1),
	}
}

func (r *AuthRepositorySpy) GetAuthenticatedUserByUsername(username string) (*dto.AuthenticatedUserDto, error) {
	r.Params["GetAuthenticatedUserByUsername:username"] = username

	r.CallsCount["GetAuthenticatedUserByUsername"]++

	if r.CallsCount["GetAuthenticatedUserByUsername"] == 1 && r.ErrorResult["GetAuthenticatedUserByUsername"] != nil {
		return nil, r.ErrorResult["GetAuthenticatedUserByUsername"]
	}

	sucess := r.SuccessResult["GetAuthenticatedUserByUsername"]
	if sucess != nil {
		return sucess.(*dto.AuthenticatedUserDto), nil
	}

	return nil, r.ErrorResult["GetAuthenticatedUserByUsername"]
}

func (r *AuthRepositorySpy) UpdatePassword(id int, value string) error {
	r.Params["UpdatePassword:id"] = id
	r.Params["UpdatePassword:value"] = value

	r.CallsCount["UpdatePassword"]++

	return r.ErrorResult["UpdatePassword"]
}

func (r *AuthRepositorySpy) UpdateUsernameHash(id int, username string) error {
	r.Params["UpdateUsernameHash:id"] = id
	r.Params["UpdateUsernameHash:username"] = username

	r.CallsCount["UpdateUsernameHash"]++

	return r.ErrorResult["UpdateUsernameHash"]
}

func (r *AuthRepositorySpy) Signup(user *dto.ActivationUserDto) (*dto.UserDto, error) {
	r.Params["Signup:user"] = user

	r.CallsCount["Signup"]++

	sucess := r.SuccessResult["Signup"]
	if sucess != nil {
		return sucess.(*dto.UserDto), nil
	}

	return nil, r.ErrorResult["Signup"]
}

func (r *AuthRepositorySpy) UserExists(username string) (*dto.UserDto, error) {
	r.Params["UserExists:username"] = username

	r.CallsCount["UserExists"]++

	sucess := r.SuccessResult["UserExists"]
	if sucess != nil {
		return sucess.(*dto.UserDto), nil
	}

	return nil, r.ErrorResult["UserExists"]
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameError() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("fake-error")
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameErrorNotFound() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("user not found")
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameSuccess() {
	r.SuccessResult["GetAuthenticatedUserByUsername"] = &dto.AuthenticatedUserDto{Password: "fake-password-hashed", FirstName: "fake-username", Id: 1, IsAdmin: false}
}

func (r *AuthRepositorySpy) DefineGetAuthenticatedUserByUsernameFallback() {
	r.ErrorResult["GetAuthenticatedUserByUsername"] = errors.New("user not found")
	r.SuccessResult["GetAuthenticatedUserByUsername"] = &dto.AuthenticatedUserDto{Password: "fake-password-hashed", FirstName: "fake-username", Id: 1, IsAdmin: false}
}

func (r *AuthRepositorySpy) DefineUpdatePasswordError() {
	r.ErrorResult["UpdatePassword"] = errors.New("fake-error")
}

func (r *AuthRepositorySpy) DefineSignupError() {
	r.ErrorResult["Signup"] = errors.New("fake-error")
}

func (r *AuthRepositorySpy) DefineSignupSuccess() {
	r.SuccessResult["Signup"] = &dto.UserDto{Id: 1}
}

func (r *AuthRepositorySpy) DefineUserExistsError() {
	r.ErrorResult["UserExists"] = errors.New("fake-error")
}

func (r *AuthRepositorySpy) DefineUserExistsSuccessUserFound() {
	r.SuccessResult["UserExists"] = &dto.UserDto{Id: 1}
}

func (r *AuthRepositorySpy) DefineUserExistsSuccess() {
	r.SuccessResult["UserExists"] = nil
}

func (r *AuthRepositorySpy) DefineUpdateUsernameHashError() {
	r.ErrorResult["UpdateUsernameHash"] = errors.New("fake-error")
}
