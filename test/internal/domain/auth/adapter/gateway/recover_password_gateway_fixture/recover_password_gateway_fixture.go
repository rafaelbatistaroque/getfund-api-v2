package recover_password_gateway_fixture

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/recover_password_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/fixture"
)

type RecoverPasswordGatewayFixture struct {
	fixture.BaseFixture
	RecoverPasswordUsecaseSpy *recoverPasswordUsecaseSpy
}

type recoverPasswordUsecaseSpy struct {
	Params        map[string]*recover_password.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*recover_password.Output
}

func NewSut() (recover_password_gateway.RecoverPasswordGateway, *RecoverPasswordGatewayFixture) {
	recoverPasswordSpy := &recoverPasswordUsecaseSpy{
		Params:        make(map[string]*recover_password.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*recover_password.Output)}

	return recover_password_gateway.New(recoverPasswordSpy),
		&RecoverPasswordGatewayFixture{
			RecoverPasswordUsecaseSpy: recoverPasswordSpy,
		}
}

func (s *recoverPasswordUsecaseSpy) Execute(input *recover_password.Input) (*recover_password.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetRecoverPasswordInputSerialized() string {
	return `{"email": "fake-username"}`
}

func GetRecoverPasswordInput() *recover_password.Input {
	return &recover_password.Input{Username: "fake-username"}
}

func (s *recoverPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *recoverPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &recover_password.Output{Message: "fake-message"}
}
