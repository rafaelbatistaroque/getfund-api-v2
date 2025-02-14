package create_user_process_started_event_handler_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/event_handler/create_user_process_started_event_handler"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
)

type CreateUserProcessStartedEventHandlerFixture struct {
	CacheSpy *cache_spy.RedisCacheSpy
}

type SendActivationAccountMailSpy struct {
	Params        map[string]*send_activation_account_mail.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*send_activation_account_mail.Output
}

func NewSut() (bus.Handler, *CreateUserProcessStartedEventHandlerFixture) {
	cacheSpy := cache_spy.New()

	return create_user_process_started_event_handler.New(cacheSpy),
		&CreateUserProcessStartedEventHandlerFixture{
			CacheSpy: cacheSpy,
		}
}

func (uc *SendActivationAccountMailSpy) Execute(input *send_activation_account_mail.Input) (*send_activation_account_mail.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *SendActivationAccountMailSpy) DefineSendActivationAccountMailSuccess() {
	uc.SuccessResult["Execute"] = &send_activation_account_mail.Output{}
}

func (uc *SendActivationAccountMailSpy) DefineSendActivationAccountMailError() {
	uc.ErrorResult["Execute"] = &result_app.ApplicationError{}
}

func GetInvalidCreateUserProcessStartedEvent() *create_user.CreateUserProcessStartedEvent {
	return &create_user.CreateUserProcessStartedEvent{}
}

func GetValidCreateUserProcessStartedEvent(withValue string) *create_user.CreateUserProcessStartedEvent {
	event := &create_user.CreateUserProcessStartedEvent{}
	event.SetPayload([]byte(withValue))

	return event
}
