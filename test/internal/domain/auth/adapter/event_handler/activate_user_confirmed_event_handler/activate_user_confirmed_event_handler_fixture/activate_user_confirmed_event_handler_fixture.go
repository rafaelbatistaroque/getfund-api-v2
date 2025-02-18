package activate_user_confirmed_event_handler_fixture

import (
	"encoding/json"
	auth_payload "getfund-api-v2/internal/domain/auth/adapter/event_handler/activate_user_confirmed_event_handler"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
)

type ActivateUserConfirmedEventHandlerFixture struct {
	SigninSpy *SigninSpy
}

type SigninSpy struct {
	Params        map[string]*signin.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*signin.Output
}

func NewSut() (bus.Handler, *ActivateUserConfirmedEventHandlerFixture) {
	signinSpy := &SigninSpy{
		Params:        make(map[string]*signin.Input, 1),
		CallsCount:    make(map[string]int, 1),
		ErrorResult:   make(map[string]*result_app.ApplicationError, 1),
		SuccessResult: make(map[string]*signin.Output, 1),
	}

	return auth_payload.New(signinSpy),
		&ActivateUserConfirmedEventHandlerFixture{
			SigninSpy: signinSpy,
		}
}

func (uc *SigninSpy) Execute(input *signin.Input) (*signin.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *SigninSpy) DefineSigninSuccessWithValue(output *signin.Output) {
	uc.SuccessResult["Execute"] = output
}

func (uc *SigninSpy) DefineSigninError() {
	uc.ErrorResult["Execute"] = &result_app.ApplicationError{}
}

func GetSigninOutput() *signin.Output {
	return &signin.Output{
		Token: "fake-valid-token",
		Session: signin.SessionOutput{
			ID:        1,
			FirstName: "fake-first-name",
			IsAdmin:   true,
		},
	}
}

func GetInvalidActivateUserConfirmedEvent() *activate_user.ActivateUserConfirmedEvent {
	return &activate_user.ActivateUserConfirmedEvent{}
}

func GetValidActivateUserConfirmedEvent(username, password string) *activate_user.ActivateUserConfirmedEvent {
	event := &activate_user.ActivateUserConfirmedEvent{}
	data := map[string]string{
		"username": username,
		"password": password,
	}

	payload, _ := json.Marshal(data)
	event.SetPayload(payload)

	return event
}

func GetActivateUserConfirmedPayload() *auth_payload.ActivateUserConfirmedPayload {
	var input = &auth_payload.ActivateUserConfirmedPayload{}

	event := GetValidActivateUserConfirmedEvent("", "")
	json.Unmarshal(event.GetPayload(), input)

	return input
}
