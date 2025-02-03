package signout_gateway_fixture

import (
	"context"
	"errors"
	signout_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signout_auth_gateway"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type SignoutGatewayFixture struct {
	SignoutUsecaseSpy *signoutUsecaseSpy
}

type signoutUsecaseSpy struct {
	Params        map[string]*signout.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*signout.Output
}

func NewSut() (signout_gateway.SignoutGateway, *SignoutGatewayFixture) {
	signoutSpy := &signoutUsecaseSpy{Params: make(map[string]*signout.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*result_app.ApplicationError), SuccessResult: make(map[string]*signout.Output)}

	return signout_gateway.New(signoutSpy),
		&SignoutGatewayFixture{
			SignoutUsecaseSpy: signoutSpy,
		}
}

func (s *signoutUsecaseSpy) Execute(input *signout.Input) (*signout.Output, *result_app.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	token := ""

	switch {
	case bodyString == "":
		token = GetSignoutHeaderToken()
	case bodyString == "not-found":
		token = ""
	default:
		token = bodyString
	}

	ctx := context.WithValue(context.Background(), auth_contract.TokenKey{}, token)
	req := httptest.NewRequest("FAKE", "/", nil).WithContext(ctx)
	res := httptest.NewRecorder()

	return res, req
}

func GetSignoutInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}

func GetSignoutHeaderToken() string {
	return "fake-token"
}

func (s *signoutUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signoutUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signout.Output{Message: "fake-message"}
}
