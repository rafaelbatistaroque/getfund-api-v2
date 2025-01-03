package recover_password_started_event_handler_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	sut "getfund-api-v2/internal/domain/notification/port/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
)

type SendRecoverPasswordMailUsecaseSpy struct {
	Params        map[string]*send_recover_password_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*send_recover_password_mail.Output
}

func NewSut() (bus.Handler, *SendRecoverPasswordMailUsecaseSpy) {
	sendRecoverPasswordMailUsecaseSpy := &SendRecoverPasswordMailUsecaseSpy{
		Params:        make(map[string]*send_recover_password_mail.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*send_recover_password_mail.Output)}

	return sut.New(sendRecoverPasswordMailUsecaseSpy), sendRecoverPasswordMailUsecaseSpy
}

func (uc *SendRecoverPasswordMailUsecaseSpy) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func GetInvalidRecoverPasswordStartedEvent() *event.RecoverPasswordStarted {
	return &event.RecoverPasswordStarted{}
}

func GetValidRecoverPasswordStartedEvent(withValue string) *event.RecoverPasswordStarted {
	event := &event.RecoverPasswordStarted{}
	event.SetPayload([]byte(withValue))

	return event
}

func GetSendRecoverPasswordMailInput(withValue string) *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{KeyCache: withValue}
}
