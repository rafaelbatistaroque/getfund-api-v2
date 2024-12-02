package signoutfixture

import (
	"context"
	"errors"
	authadapter "getfund-api-v2/internal/domain/auth/port/adapter"
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"
	"net/http/httptest"
)

type signoutUsecaseSpy struct {
	Params        map[string]*signout.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*resultapp.ApplicationError
	SuccessResult map[string]*signout.Output
}

func NewSut() (authadapter.AuthAdapter, *signoutUsecaseSpy) {
	signoutSpy := &signoutUsecaseSpy{Params: make(map[string]*signout.Input), CallsCount: make(map[string]int), ErrorResult: make(map[string]*resultapp.ApplicationError), SuccessResult: make(map[string]*signout.Output)}
	return authadapter.New(nil, signoutSpy), signoutSpy
}

func (s *signoutUsecaseSpy) Execute(input *signout.Input) (*signout.Output, *resultapp.ApplicationError) {
	s.Params["Execute:input"] = input

	s.CallsCount["Execute"]++

	return s.SuccessResult["Execute"], s.ErrorResult["Execute"]
}

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	session := ""

	switch {
	case bodyString == "":
		session = GetSignoutHeaderSerialized()
	case bodyString == "not-found":
		session = ""
	default:
		session = bodyString
	}

	ctx := context.WithValue(context.Background(), sessionservice.SessionKey{}, session)
	req := httptest.NewRequest("FAKE", "/", nil).WithContext(ctx)
	res := httptest.NewRecorder()

	return res, req
}

func GetSignoutInput() *signout.Input {
	return &signout.Input{Token: "fake-token"}
}

func GetSignoutHeaderSerialized() string {
	return "{\"token\": \"fake-token\", \"session\": {\"id\":\"fake-id\",\"first_name\":\"fake-firstname\",\"is_admin\":true}}"
}

func (s *signoutUsecaseSpy) DefineError() {
	s.ErrorResult["Execute"] = &resultapp.ApplicationError{Code: resultapp.CODE_SERVER_ERROR, Message: errors.New("fake-error")}
}

func (s *signoutUsecaseSpy) DefineSuccess() {
	s.SuccessResult["Execute"] = &signout.Output{Message: "fake-message"}
}
