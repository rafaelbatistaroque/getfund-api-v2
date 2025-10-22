package signin_gateway_fixture

import (
	"errors"
	signin_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signin_auth_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/fixture"
)

type SigninGatewayFixture struct {
	fixture.BaseFixture
	SigninUsecaseSpy *signinUsecaseSpy
}

type signinUsecaseSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*signin.Output
}

func NewSut() (signin_gateway.SigninGateway, *SigninGatewayFixture) {
	signinSpy := &signinUsecaseSpy{
		Params:        make(map[string]*signin.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*signin.Output)}

	return signin_gateway.New(signinSpy),
		&SigninGatewayFixture{
			SigninUsecaseSpy: signinSpy,
		}
}

func (s *signinUsecaseSpy) Execute(input *signin.Input) (*signin.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetSigninInput() *signin.Input {
	return &signin.Input{Username: "fake-username", Password: "fake-password"}
}

func GetSigninInputSerialized() string {
	return `{"username": "fake-username", "password": "fake-password"}`
}

func (s *signinUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signinUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{Token: "fake-token", Session: signin.SessionOutput{ID: 1, FirstName: "fake-firstname", IsAdmin: true}}
}
