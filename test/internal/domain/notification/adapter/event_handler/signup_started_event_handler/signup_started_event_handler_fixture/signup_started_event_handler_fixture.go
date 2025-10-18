package signup_process_started_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup/event"
	notification_handler "getfund-api-v2/internal/domain/notification/adapter/event_handler/signup_started_event_handler"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_error "getfund-api-v2/internal/shared/error"
)

type SignupProcessStartedEventHandlerFixture struct {
	SendActivationAccountMailSpy *SendActivationAccountMailSpy
}

type SendActivationAccountMailSpy struct {
	Params        map[string]*send_activation_account_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*send_activation_account_mail.Output
}

func NewSut() (shared_bus.Handler, *SignupProcessStartedEventHandlerFixture) {
	sendActivationAccountMailSpy := &SendActivationAccountMailSpy{
		Params:        make(map[string]*send_activation_account_mail.Input, 1),
		CallsCount:    make(map[string]int, 1),
		ErrorResult:   make(map[string]*shared_error.Error, 1),
		SuccessResult: make(map[string]*send_activation_account_mail.Output, 1),
	}

	return notification_handler.New(sendActivationAccountMailSpy),
		&SignupProcessStartedEventHandlerFixture{
			SendActivationAccountMailSpy: sendActivationAccountMailSpy,
		}
}

func (uc *SendActivationAccountMailSpy) Execute(input *send_activation_account_mail.Input) (*send_activation_account_mail.Output, *shared_error.Error) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++
	uc.DefineSendActivationAccountMailSuccess()
	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *SendActivationAccountMailSpy) DefineSendActivationAccountMailSuccess() {
	uc.SuccessResult["Execute"] = &send_activation_account_mail.Output{}
}

func (uc *SendActivationAccountMailSpy) DefineSendActivationAccountMailError() {
	uc.ErrorResult["Execute"] = &shared_error.Error{}
}

func GetInvalidSignupProcessStartedEvent() *event.SignupStartedEvent {
	return &event.SignupStartedEvent{}
}

func GetValidSignupProcessStartedEvent() *event.SignupStartedEvent {
	signup_event := &event.SignupStartedEvent{}
	data := &event.SignupStartedPayload{
		FirstName:      "fake-first-name",
		Email:          "fake-email",
		ActivationLink: "fake-activation-link",
	}

	payload, _ := json.Marshal(data)
	signup_event.SetPayload(payload)

	return signup_event
}

func GetSendActivationAccountMailInput() *send_activation_account_mail.Input {
	return &send_activation_account_mail.Input{
		FirstName:      "fake-first-name",
		Email:          "fake-email",
		ActivationLink: "fake-activation-link",
	}
}
