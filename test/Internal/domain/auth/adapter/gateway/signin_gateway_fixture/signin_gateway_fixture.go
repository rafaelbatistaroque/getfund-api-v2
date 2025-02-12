package signin_gateway_fixture

import (
	"bytes"
	"errors"
	signin_gateway "getfund-api-v2/internal/domain/auth/adapter/gateway/signin_auth_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type SigninGatewayFixture struct {
	SigninUsecaseSpy *signinUsecaseSpy
}

type signinUsecaseSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*signin.Output
}

func NewSut() (signin_gateway.SigninGateway, *SigninGatewayFixture) {
	signinSpy := &signinUsecaseSpy{
		Params:        make(map[string]*signin.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*signin.Output)}

	return signin_gateway.New(signinSpy),
		&SigninGatewayFixture{
			SigninUsecaseSpy: signinSpy,
		}
}

func (s *signinUsecaseSpy) Execute(input *signin.Input) (*signin.Output, *result_app.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetSigninInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetSigninInput() *signin.Input {
	return &signin.Input{UserName: "fake-username", Password: "fake-password"}
}

func GetSigninInputSerialized() string {
	return `{"username": "fake-username", "password": "fake-password"}`
}

func (s *signinUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signinUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{Token: "fake-token", Session: signin.SessionOutput{ID: 1, FirstName: "fake-firstname", IsAdmin: true}}
}
