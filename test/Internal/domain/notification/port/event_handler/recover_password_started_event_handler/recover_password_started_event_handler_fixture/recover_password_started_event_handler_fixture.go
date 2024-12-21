package recover_password_started_event_handler_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	sut "getfund-api-v2/internal/domain/notification/port/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/eventbus"
)

type SendRecoverPasswordMailUsecaseSpy struct {
	Params        map[string]*send_recover_password_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*send_recover_password_mail.Output
}

func NewSut() (eventbus.Handler, *SendRecoverPasswordMailUsecaseSpy) {
	sendRecoverPasswordMailUsecaseSpy := &SendRecoverPasswordMailUsecaseSpy{
		Params:        make(map[string]*send_recover_password_mail.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*send_recover_password_mail.Output)}

	return sut.New(sendRecoverPasswordMailUsecaseSpy), sendRecoverPasswordMailUsecaseSpy
}

func (uc *SendRecoverPasswordMailUsecaseSpy) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	return nil, nil
}
