package session_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/dto"
)

type SessionServiceSpy struct {
	Params     map[string]any
	CallsCount map[string]int

	SuccessResult map[string]any
	ErrorResult   map[string]error
}

func New() *SessionServiceSpy {
	return &SessionServiceSpy{CallsCount: make(map[string]int), Params: make(map[string]any), SuccessResult: make(map[string]any), ErrorResult: make(map[string]error)}
}

func (s *SessionServiceSpy) SaveSession(session *dto.SessionDto) (string, error) {
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
