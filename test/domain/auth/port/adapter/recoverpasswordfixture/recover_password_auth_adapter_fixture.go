package recoverpasswordfixture

import (
	"bytes"
	"errors"
	adapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/resultapp"
	"net/http"
	"net/http/httptest"
)

type recoverPasswordUsecaseSpy struct {
	Params        map[string]*recoverpassword.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*resultapp.ApplicationError
	SuccessResult map[string]*recoverpassword.Output
}

func NewSut() (adapter.AuthAdapter, *recoverPasswordUsecaseSpy) {
	recoverPasswordSpy := &recoverPasswordUsecaseSpy{Params: make(map[string]*recoverpassword.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*resultapp.ApplicationError), SuccessResult: make(map[string]*recoverpassword.Output)}
	return adapter.New(nil, nil, recoverPasswordSpy), recoverPasswordSpy
}

func (s *recoverPasswordUsecaseSpy) Execute(input *recoverpassword.Input) (*recoverpassword.Output, *resultapp.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetRecoverPasswordInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetRecoverPasswordInputSerialized() string {
	return `{"email": "fake-username"}`
}

func GetRecoverPasswordInput() *recoverpassword.Input {
	return &recoverpassword.Input{UserName: "fake-username"}
}

func (s *recoverPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &resultapp.ApplicationError{Code: resultapp.CODE_SERVER_ERROR, Message: errors.New("fake-error")}
}

func (s *recoverPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &recoverpassword.Output{Message: "fake-message"}
}
