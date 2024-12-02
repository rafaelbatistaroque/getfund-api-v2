package authadapterfixture

import (
	"bytes"
	"errors"
	authadapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/shared/applicationcode"
	"getfund-api-v2/internal/shared/applicationerror"
	"net/http"
	"net/http/httptest"
)

type signinSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*applicationerror.ApplicationError
	SuccessResult map[string]*signin.Output
}

func NewSut() (authadapter.AuthAdapter, *signinSpy) {
	signinSpy := &signinSpy{Params: make(map[string]*signin.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*applicationerror.ApplicationError), SuccessResult: make(map[string]*signin.Output)}
	return authadapter.New(signinSpy), signinSpy
}

func (s *signinSpy) Execute(input *signin.Input) (*signin.Output, *applicationerror.ApplicationError) {
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

func (s *signinSpy) DefineError() {
	s.ErrorResult["Execute"] = &applicationerror.ApplicationError{Code: applicationcode.CODE_SERVER_ERROR, Message: errors.New("fake-error")}
}

func (s *signinSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{Token: "fake-token", Session: signin.SessionOutput{ID: "fake-id", FirstName: "fake-firstname", IsAdmin: true}}
}
