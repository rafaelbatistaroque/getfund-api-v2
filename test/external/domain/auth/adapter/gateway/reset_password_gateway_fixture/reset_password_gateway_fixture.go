package reset_password_gateway_fixture

import (
	"bytes"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type ResetPasswordGatewayFixture struct {
	ResetPasswordUsecaseSpy *resetPasswordUsecaseSpy
}

type resetPasswordUsecaseSpy struct {
	Params        map[string]*reset_password.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*reset_password.Output
}

func NewSut() (reset_password_gateway.ResetPasswordGateway, *ResetPasswordGatewayFixture) {
	resetPasswordSpy := &resetPasswordUsecaseSpy{
		Params:        make(map[string]*reset_password.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*reset_password.Output)}

	return reset_password_gateway.New(resetPasswordSpy),
		&ResetPasswordGatewayFixture{
			ResetPasswordUsecaseSpy: resetPasswordSpy}
}

func (s *resetPasswordUsecaseSpy) Execute(input *reset_password.Input) (*reset_password.Output, *result_app.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetResetPasswordInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetResetPasswordInputSerialized() string {
	return `{"code": "fake-recovery-code", "password": "fake-password"}`
}

func GetResetPasswordInput() *reset_password.Input {
	return &reset_password.Input{
		RecoveryCode: "fake-recovery-code",
		Password:     "fake-password",
		RecoveryKey:  "recovery_password_fake-recovery-code",
	}
}

func (s *resetPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *resetPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &reset_password.Output{Message: "fake-message"}
}
