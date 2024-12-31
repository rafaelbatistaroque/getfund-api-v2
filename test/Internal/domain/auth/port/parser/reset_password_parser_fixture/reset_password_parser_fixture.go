package reset_password_parser_fixture

import (
	"bytes"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password"
	parser "getfund-api-v2/internal/domain/auth/port/parser"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type resetPasswordUsecaseSpy struct {
	Params        map[string]*reset_password.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*reset_password.Output
}

func NewSut() (parser.AuthParser, *resetPasswordUsecaseSpy) {
	resetPasswordSpy := &resetPasswordUsecaseSpy{
		Params:        make(map[string]*reset_password.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*reset_password.Output)}

	return parser.New(nil, nil, nil, resetPasswordSpy), resetPasswordSpy
}

func (s *resetPasswordUsecaseSpy) Execute(input *reset_password.Input) (*reset_password.Output, *result_app.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetRecoverPasswordInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetRecoverPasswordInputSerialized() string {
	return `{"code": "fake-code", "password": "fake-password"}`
}

func GetRecoverPasswordInput() *recover_password.Input {
	return &recover_password.Input{Username: "fake-username"}
}

func (s *resetPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *resetPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &reset_password.Output{Message: "fake-message"}
}
