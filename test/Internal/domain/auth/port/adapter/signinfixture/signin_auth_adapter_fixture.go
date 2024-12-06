package signinfixture

import (
	"bytes"
	"errors"
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
	"net/http/httptest"
)

type signinUsecaseSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*resultapp.ApplicationError
	SuccessResult map[string]*signin.Output
}

func NewSut() (adapter.AuthAdapter, *signinUsecaseSpy) {
	signinSpy := &signinUsecaseSpy{Params: make(map[string]*signin.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*resultapp.ApplicationError), SuccessResult: make(map[string]*signin.Output)}
	return adapter.New(signinSpy, nil, nil), signinSpy
}

func (s *signinUsecaseSpy) Execute(input *signin.Input) (*signin.Output, *resultapp.ApplicationError) {
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
	s.ErrorResult["Execute"] = &resultapp.ApplicationError{Code: resultapp.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signinUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{Token: "fake-token", Session: signin.SessionOutput{ID: "fake-id", FirstName: "fake-firstname", IsAdmin: true}}
}
