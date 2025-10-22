package reset_password_gateway_fixture

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/reset_password_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/fixture"
)

type ResetPasswordGatewayFixture struct {
	fixture.BaseFixture
	ResetPasswordUsecaseSpy *resetPasswordUsecaseSpy
}

type resetPasswordUsecaseSpy struct {
	Params        map[string]*reset_password.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*reset_password.Output
}

func NewSut() (reset_password_gateway.ResetPasswordGateway, *ResetPasswordGatewayFixture) {
	resetPasswordSpy := &resetPasswordUsecaseSpy{
		Params:        make(map[string]*reset_password.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*reset_password.Output)}

	return reset_password_gateway.New(resetPasswordSpy),
		&ResetPasswordGatewayFixture{
			ResetPasswordUsecaseSpy: resetPasswordSpy}
}

func (s *resetPasswordUsecaseSpy) Execute(input *reset_password.Input) (*reset_password.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
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
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *resetPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &reset_password.Output{Message: "fake-message"}
}
