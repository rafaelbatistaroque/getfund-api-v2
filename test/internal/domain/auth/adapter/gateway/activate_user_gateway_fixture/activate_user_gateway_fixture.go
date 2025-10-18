package activate_user_gateway_fixture

import (
	"errors"
	"fmt"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/activate_user_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
	"net/http/httptest"
)

type ActivateUserUsecaseFixture struct {
	ActivateUserUsecaseSpy *activateUserUsecaseSpy
	SigninUsecaseSpy       *signinUsecaseSpy
}

type activateUserUsecaseSpy struct {
	Params        map[string]*activate_user.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*activate_user.Output
}

type signinUsecaseSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*signin.Output
}

func NewSut() (activate_user_gateway.ActiveUserGateway, *ActivateUserUsecaseFixture) {
	activateUserUsecase := &activateUserUsecaseSpy{
		Params:        make(map[string]*activate_user.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*activate_user.Output)}

	signinUsecase := &signinUsecaseSpy{
		Params:        make(map[string]*signin.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*signin.Output)}

	return activate_user_gateway.New(activateUserUsecase, signinUsecase),
		&ActivateUserUsecaseFixture{
			ActivateUserUsecaseSpy: activateUserUsecase,
			SigninUsecaseSpy:       signinUsecase,
		}
}

func (s *activateUserUsecaseSpy) Execute(input *activate_user.Input) (*activate_user.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func (s *signinUsecaseSpy) Execute(input *signin.Input) (*signin.Output, *shared_error.Error) {
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
	fakeActivationCode := "fake-activation-code"
	return &activate_user.Input{
		ActivationCode:    fakeActivationCode,
		ActivationDataKey: "user_activation_" + fakeActivationCode}
}

func (s *activateUserUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *activateUserUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &activate_user.Output{}
}

func (s *activateUserUsecaseSpy) DefineSuccessWithValue() {
	s.SuccessResult["Execute"] = &activate_user.Output{
		Username: "fake-username",
		Password: "fake-password",
	}
}

func (s *signinUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signinUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{}
}
