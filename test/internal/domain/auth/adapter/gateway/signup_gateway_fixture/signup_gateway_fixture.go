package signup_gateway_fixture

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/signup_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/fixture"
)

type SignupGatewayFixture struct {
	fixture.BaseFixture
	SignupUsecaseSpy *SignupUsecaseSpy
}

type SignupUsecaseSpy struct {
	Params        map[string]*signup.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*signup.Output
}

func NewSut() (signup_gateway.SignupGateway, *SignupGatewayFixture) {
	SignupSpy := &SignupUsecaseSpy{
		Params:        make(map[string]*signup.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*signup.Output)}

	return signup_gateway.New(SignupSpy), &SignupGatewayFixture{SignupUsecaseSpy: SignupSpy}
}

func (s *SignupUsecaseSpy) Execute(input *signup.Input) (*signup.Output, *shared_error.Error) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetSignupInput() *signup.Input {
	return &signup.Input{
		FirstName:            "fake-first-name",
		LastName:             "fake-last-name",
		Username:             "fake@email.com",
		Password:             "fakaStrongPass123",
		PasswordConfirmation: "fakaStrongPass123",
		CouponCode:           "COUPONCD",
	}
}

func GetSigninInputSerialized() string {
	return `{"first_name":"fake-first-name","last_name":"fake-last-name","username":"fake@email.com","password":"fakaStrongPass123","password_confirmation":"fakaStrongPass123","cupon_code":"COUPONCD"}`
}

func (s *SignupUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &shared_error.Error{Code: shared_error.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *SignupUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signup.Output{}
}
