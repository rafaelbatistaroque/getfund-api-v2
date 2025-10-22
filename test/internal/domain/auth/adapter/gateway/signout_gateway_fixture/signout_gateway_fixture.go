package signout_gateway_fixture

import (
	"context"
	"errors"
	signout_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signout_auth_gateway"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/fixture"
	"net/http"
)

type SignoutGatewayFixture struct {
	fixture.BaseFixture
	SignoutUsecaseSpy *signoutUsecaseSpy
}

func (f *SignoutGatewayFixture) GetHttpRequestResponseWithContext() (w http.ResponseWriter, r *http.Request) {
	w, r = f.GetHttpRequestResponse("")
	ctx := context.WithValue(context.Background(), auth_contract.TokenKey{}, GetSignoutHeaderToken())
	r = r.WithContext(ctx)
	return w, r
}

type signoutUsecaseSpy struct {
	Params        map[string]*signout.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*signout.Output
}

func NewSut() (signout_gateway.SignoutGateway, *SignoutGatewayFixture) {
	signoutSpy := &signoutUsecaseSpy{Params: make(map[string]*signout.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*shared_error.Error), SuccessResult: make(map[string]*signout.Output)}

	return signout_gateway.New(signoutSpy),
		&SignoutGatewayFixture{
			SignoutUsecaseSpy: signoutSpy,
		}
}

func (s *signoutUsecaseSpy) Execute(input *signout.Input) (*signout.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetSignoutInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}

func GetSignoutHeaderToken() string {
	return "fake-token"
}

func (s *signoutUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signoutUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signout.Output{Message: "fake-message"}
}
