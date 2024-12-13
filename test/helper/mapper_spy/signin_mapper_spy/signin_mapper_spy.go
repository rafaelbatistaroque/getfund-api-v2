package signin_mapper_spy

import (
	"errors"
	model "getfund-api-v2/internal/domain/auth/adapter/model"
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signin"
)

type Spy struct {
}

type SigninMapperSpy struct {
	Params      map[string]interface{}
	ForceReturn bool

	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *SigninMapperSpy {
	return &SigninMapperSpy{Params: make(map[string]interface{}), ForceReturn: true, ErrorResult: make(map[string]error), CallsCount: make(map[string]int), SuccessResult: make(map[string]interface{})}
}

func (m *SigninMapperSpy) ToOutput(token string, session *model.SessionModel) *signin.Output {
	m.Params["ToOutput:session"] = session
	m.Params["ToOutput:token"] = token

	m.CallsCount["ToOutput"]++

	if m.ForceReturn {
		return nil
	}

	m.SuccessResult["ToOutput"] = &signin.SigninOutput{
		Token: token,
		Session: signin.SessionOutput{
			ID:        session.ID,
			FirstName: session.FirstName,
			IsAdmin:   session.IsAdmin == 1,
		},
	}

	return m.SuccessResult["ToOutput"].(*signin.Output)
}

func (m *SigninMapperSpy) SessionToString(session *model.SessionModel) (string, error) {
	m.Params["SessionToString:session"] = session

	m.CallsCount["SessionToString"]++

	success := m.SuccessResult["SessionToString"]
	if success != nil {
		return m.SuccessResult["SessionToString"].(string), m.ErrorResult["SessionToString"]
	}

	return "", m.ErrorResult["SessionToString"]
}

func (m *SigninMapperSpy) ToSessionModel(user *model.UserModel) *model.SessionModel {
	m.Params["ToSessionModel:user"] = user

	success := m.SuccessResult["ToSessionModel"]
	if success != nil {
		return m.SuccessResult["ToSessionModel"].(*model.SessionModel)
	}

	return nil
}

func (m *SigninMapperSpy) DefineError() {
	m.ErrorResult["SessionToString"] = errors.New("any-error")
}

func (m *SigninMapperSpy) DefineSuccess() {
	m.SuccessResult["SessionToString"] = "fake-success"
}

func (m *SigninMapperSpy) DefineToSessionModelSuccess(user *model.UserModel) {
	m.SuccessResult["ToSessionModel"] = &model.SessionModel{
		ID:      user.Id,
		IsAdmin: user.IsAdmin,
	}
}
