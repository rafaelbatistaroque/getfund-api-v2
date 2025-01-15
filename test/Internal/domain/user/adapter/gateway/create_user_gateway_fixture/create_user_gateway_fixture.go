package create_user_gateway_fixture

import (
	"bytes"
	"errors"
	user_gateway "getfund-api-v2/internal/domain/user/adapter/gateway"
	"getfund-api-v2/internal/domain/user/core/usercase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type createUserUsecaseSpy struct {
	Params        map[string]*create_user.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*create_user.Output
}

func NewSut() (user_gateway.UserGateway, *createUserUsecaseSpy) {
	createUserSpy := &createUserUsecaseSpy{
		Params:        make(map[string]*create_user.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*create_user.Output)}

	return user_gateway.New(createUserSpy), createUserSpy
}

func (s *createUserUsecaseSpy) Execute(input *create_user.Input) (*create_user.Output, *result_app.ApplicationError) {
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

func GetSigninInput() *create_user.Input {
	return &create_user.Input{}
}

func GetSigninInputSerialized() string {
	return `{}`
}

func (s *createUserUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *createUserUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &create_user.Output{}
}
