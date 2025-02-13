package recover_password_started_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	sut "getfund-api-v2/internal/domain/notification/adapter/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
)

type RecoverPasswordStartedEventHandlerFixture struct {
	SendRecoverPasswordMailUsecaseSpy *SendRecoverPasswordMailUsecaseSpy
	CacheSpy                          *cache_spy.RedisCacheSpy
}

type SendRecoverPasswordMailUsecaseSpy struct {
	Params        map[string]*send_recover_password_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*send_recover_password_mail.Output
}

func NewSut() (bus.Handler, *RecoverPasswordStartedEventHandlerFixture) {
	sendRecoverPasswordMailUsecaseSpy := &SendRecoverPasswordMailUsecaseSpy{
		Params:        make(map[string]*send_recover_password_mail.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*send_recover_password_mail.Output)}

	cacheSpy := cache_spy.New()
	return sut.New(sendRecoverPasswordMailUsecaseSpy, cacheSpy),
		&RecoverPasswordStartedEventHandlerFixture{
			SendRecoverPasswordMailUsecaseSpy: sendRecoverPasswordMailUsecaseSpy,
			CacheSpy:                          cacheSpy,
		}
}

func (uc *SendRecoverPasswordMailUsecaseSpy) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *SendRecoverPasswordMailUsecaseSpy) DefineSendRecoverPasswordMailUsecaseSuccess() {
	uc.SuccessResult["Execute"] = &send_recover_password_mail.Output{}
}

func (uc *SendRecoverPasswordMailUsecaseSpy) DefineSendRecoverPasswordMailUsecaseError() {
	uc.ErrorResult["Execute"] = &result_app.ApplicationError{}
}

func GetInvalidRecoverPasswordStartedEvent() *recover_password.RecoverPasswordStartedEvent {
	return &recover_password.RecoverPasswordStartedEvent{}
}

func GetValidRecoverPasswordStartedEvent(withValue string) *recover_password.RecoverPasswordStartedEvent {
	event := &recover_password.RecoverPasswordStartedEvent{}
	event.SetPayload([]byte(withValue))

	return event
}

func GetValidCacheData() string {
	return `{"username":"fake-username", "first_name":"fake-first-name", "recovery_link":"fake-recovery_link"}`
}

func GetSendRecoverPasswordMailInput() *send_recover_password_mail.Input {
	var input = &send_recover_password_mail.Input{}

	json.Unmarshal([]byte(GetValidCacheData()), input)

	return input
}
