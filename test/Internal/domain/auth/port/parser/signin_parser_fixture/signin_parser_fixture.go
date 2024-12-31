package signin_parser_fixture

import (
	"bytes"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signin"
	parser "getfund-api-v2/internal/domain/auth/port/parser"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type signinUsecaseSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*signin.Output
}

func NewSut() (parser.AuthParser, *signinUsecaseSpy) {
	signinSpy := &signinUsecaseSpy{
		Params:        make(map[string]*signin.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*signin.Output)}

	return parser.New(signinSpy, nil, nil, nil), signinSpy
}

func (s *signinUsecaseSpy) Execute(input *signin.Input) (*signin.Output, *result_app.ApplicationError) {
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

func GetSigninInput() *signin.Input {
	return &signin.Input{UserName: "fake-username", Password: "fake-password"}
}

func GetSigninInputSerialized() string {
	return `{"username": "fake-username", "password": "fake-password"}`
}

func (s *signinUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *signinUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signin.Output{Token: "fake-token", Session: signin.SessionOutput{ID: "fake-id", FirstName: "fake-firstname", IsAdmin: true}}
}
