package session_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
)

type SessionServiceSpy struct {
	Params     map[string]interface{}
	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *SessionServiceSpy {
	return &SessionServiceSpy{CallsCount: make(map[string]int), Params: make(map[string]interface{}), SuccessResult: make(map[string]interface{}), ErrorResult: make(map[string]error)}
}

func (s *SessionServiceSpy) SaveSession(session *auth_dto.SessionDto) (string, error) {
	s.Params["SaveSession:session"] = session

	s.CallsCount["SaveSession"]++

	success := s.SuccessResult["SaveSession"]
	if success != nil {
		return s.SuccessResult["SaveSession"].(string), s.ErrorResult["SaveSession"]
	}

	return "", s.ErrorResult["SaveSession"]
}

func (s *SessionServiceSpy) DeleteSession(token string) error {
	s.Params["DeleteSession:token"] = token

	s.CallsCount["DeleteSession"]++

	return s.ErrorResult["DeleteSession"]
}

func (s *SessionServiceSpy) GetSession(session string) (string, error) {
	s.Params["GetSession:token"] = session

	s.CallsCount["GetSession"]++

	success := s.SuccessResult["GetSession"]
	if success != nil {
		return s.SuccessResult["GetSession"].(string), s.ErrorResult["GetSession"]
	}

	return "", s.ErrorResult["GetSession"]
}

func (s *SessionServiceSpy) DefineSaveSessionError() {
	s.ErrorResult["SaveSession"] = errors.New("any-error")
}

func (s *SessionServiceSpy) DefineSaveSessionSuccess() {
	s.SuccessResult["SaveSession"] = "fake-success"
}

func (s *SessionServiceSpy) DefineDeleteSessionError() {
	s.ErrorResult["DeleteSession"] = errors.New("any-error")
}

func (s *SessionServiceSpy) DefineDeleteSessionSuccess() {
	s.SuccessResult["DeleteSession"] = "fake-success"
}

func (s *SessionServiceSpy) DefineGetSessionError() {
	s.ErrorResult["GetSession"] = errors.New("any-error")
}

func (s *SessionServiceSpy) DefineGetSessionSuccess() {
	s.SuccessResult["GetSession"] = "fake-data-hashed"
}
