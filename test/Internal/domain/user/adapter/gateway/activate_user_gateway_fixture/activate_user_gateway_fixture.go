package activate_user_gateway_fixture

import (
	"errors"
	"fmt"
	"getfund-api-v2/internal/domain/user/adapter/activate_user_gateway"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type ActivateUserUsecaseFixture struct {
	ActivateUserUsecaseSpy *activateUserUsecaseSpy
}

type activateUserUsecaseSpy struct {
	Params        map[string]*activate_user.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*activate_user.Output
}

func NewSut() (activate_user_gateway.ActiveUserGateway, *ActivateUserUsecaseFixture) {
	activateUserGateway := &activateUserUsecaseSpy{
		Params:        make(map[string]*activate_user.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*activate_user.Output)}

	return activate_user_gateway.New(activateUserGateway),
		&ActivateUserUsecaseFixture{
			ActivateUserUsecaseSpy: activateUserGateway}
}

func (s *activateUserUsecaseSpy) Execute(input *activate_user.Input) (*activate_user.Output, *result_app.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(activationCode string) (w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("/user/activate/%s", activationCode)
	req := httptest.NewRequest("FAKE", url, nil)
	res := httptest.NewRecorder()

	return res, req
}

func GetActivateUserInput() *activate_user.Input {
	return &activate_user.Input{ActivationCode: "fake-activation-code"}
}

func (s *activateUserUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *activateUserUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &activate_user.Output{}
}
