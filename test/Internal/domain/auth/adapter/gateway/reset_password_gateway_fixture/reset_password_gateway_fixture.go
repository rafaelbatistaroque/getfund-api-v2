package reset_password_gateway_fixture

import (
	"bytes"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/test/helper/cache_spy"
	"net/http"
	"net/http/httptest"
)

type ResetPasswordGatewayFixture struct {
	ResetPasswordUsecaseSpy *resetPasswordUsecaseSpy
	CacheSpy                *cache_spy.RedisCacheSpy
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

	cacheSpy := cache_spy.New()

	return reset_password_gateway.New(resetPasswordSpy),
		&ResetPasswordGatewayFixture{
			ResetPasswordUsecaseSpy: resetPasswordSpy,
			CacheSpy:                cacheSpy}
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
	return &reset_password.Input{RecoveryCode: "fake-recovery-code", Password: "fake-password"}
}

func (s *resetPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *resetPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &reset_password.Output{Message: "fake-message"}
}
