package recover_password_started_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password/event"
	sut "getfund-api-v2/internal/domain/notification/adapter/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/test/helper/cache_spy"
)

type RecoverPasswordStartedEventHandlerFixture struct {
	SendRecoverPasswordMailUsecaseSpy *SendRecoverPasswordMailUsecaseSpy
	CacheSpy                          *cache_spy.RedisCacheSpy
}

type SendRecoverPasswordMailUsecaseSpy struct {
	Params        map[string]*send_recover_password_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*send_recover_password_mail.Output
}

func NewSut() (shared_bus.Handler, *RecoverPasswordStartedEventHandlerFixture) {
	sendRecoverPasswordMailUsecaseSpy := &SendRecoverPasswordMailUsecaseSpy{
		Params:        make(map[string]*send_recover_password_mail.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*send_recover_password_mail.Output)}

	cacheSpy := cache_spy.New()
	return sut.New(sendRecoverPasswordMailUsecaseSpy, cacheSpy),
		&RecoverPasswordStartedEventHandlerFixture{
			SendRecoverPasswordMailUsecaseSpy: sendRecoverPasswordMailUsecaseSpy,
			CacheSpy:                          cacheSpy,
		}
}

func (uc *SendRecoverPasswordMailUsecaseSpy) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *shared_error.Error) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *SendRecoverPasswordMailUsecaseSpy) DefineSendRecoverPasswordMailUsecaseSuccess() {
	uc.SuccessResult["Execute"] = &send_recover_password_mail.Output{}
}

func (uc *SendRecoverPasswordMailUsecaseSpy) DefineSendRecoverPasswordMailUsecaseError() {
	uc.ErrorResult["Execute"] = &shared_error.Error{}
}

func GetInvalidRecoverPasswordStartedEvent() *event.RecoverPasswordStartedEvent {
	return &event.RecoverPasswordStartedEvent{}
}

func GetValidRecoverPasswordStartedEvent(withValue string) *event.RecoverPasswordStartedEvent {
	event := &event.RecoverPasswordStartedEvent{}
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
