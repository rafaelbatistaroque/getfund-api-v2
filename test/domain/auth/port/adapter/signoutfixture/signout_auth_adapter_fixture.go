package signoutfixture

import (
	"bytes"
	"errors"
	authadapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
	"net/http/httptest"
)

type signoutUsecaseSpy struct {
	Params        map[string]*signout.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*resultapp.ApplicationError
	SuccessResult map[string]*signout.Output
}

func NewSut() (authadapter.AuthAdapter, *signoutUsecaseSpy) {
	signoutSpy := &signoutUsecaseSpy{Params: make(map[string]*signout.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*resultapp.ApplicationError), SuccessResult: make(map[string]*signout.Output)}
	return authadapter.New(nil, signoutSpy), signoutSpy
}

func (s *signoutUsecaseSpy) Execute(input *signout.Input) (*signout.Output, *resultapp.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetSignoutInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetSignoutInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}

func GetSignoutInputSerialized() string {
	return `{"token": "fake-token"}`
}

func (s *signoutUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &resultapp.ApplicationError{Code: resultapp.CODE_SERVER_ERROR, Message: errors.New("fake-error")}
}

func (s *signoutUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signout.Output{Message: "fake-message"}
}
