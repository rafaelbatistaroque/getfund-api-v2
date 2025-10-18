package signup_gateway_fixture

import (
	"bytes"
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/gateway/signup_gateway"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup"
	shared_error "getfund-api-v2/internal/shared/error"
	"net/http"
	"net/http/httptest"
)

type SignupUsecaseSpy struct {
	Params        map[string]*signup.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*signup.Output
}

func NewSut() (signup_gateway.SignupGateway, *SignupUsecaseSpy) {
	SignupSpy := &SignupUsecaseSpy{
		Params:        make(map[string]*signup.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*signup.Output)}

	return signup_gateway.New(SignupSpy), SignupSpy
}

func (s *SignupUsecaseSpy) Execute(input *signup.Input) (*signup.Output, *shared_error.Error) {
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

func GetHttpWithBody(body any) (w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(body)
	return httptest.NewRecorder(), httptest.NewRequest("FAKE", "/", bytes.NewBuffer(json))
}

func GetHttpDefault() (w http.ResponseWriter, r *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest("FAKE", "/", bytes.NewBufferString("{}"))
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
