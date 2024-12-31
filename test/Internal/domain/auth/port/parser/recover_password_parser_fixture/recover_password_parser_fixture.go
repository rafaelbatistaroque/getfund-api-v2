package recover_password_parser_fixture

import (
	"bytes"
	"errors"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/recover_password"
	parser "getfund-api-v2/internal/domain/auth/port/parser"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
	"net/http/httptest"
)

type recoverPasswordUsecaseSpy struct {
	Params        map[string]*recover_password.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*recover_password.Output
}

func NewSut() (parser.AuthParser, *recoverPasswordUsecaseSpy) {
	recoverPasswordSpy := &recoverPasswordUsecaseSpy{
		Params:        make(map[string]*recover_password.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*recover_password.Output)}

	return parser.New(nil, nil, recoverPasswordSpy, nil), recoverPasswordSpy
}

func (s *recoverPasswordUsecaseSpy) Execute(input *recover_password.Input) (*recover_password.Output, *result_app.ApplicationError) {
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
	return `{"email": "fake-username"}`
}

func GetRecoverPasswordInput() *recover_password.Input {
	return &recover_password.Input{Username: "fake-username"}
}

func (s *recoverPasswordUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &result_app.ApplicationError{Code: result_app.SERVER_ERROR_CODE, Message: errors.New("fake-error")}
}

func (s *recoverPasswordUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &recover_password.Output{Message: "fake-message"}
}
