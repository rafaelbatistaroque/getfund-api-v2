package user_criation_with_coupon_code_started_event_handler_fixture

import (
	"getfund-api-v2/internal/domain/coupon/adapter/event_handler/user_criation_with_coupon_code_started_event_handler"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon_code"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
	"getfund-api-v2/test/helper/cache_spy"
)

type UserCriationWithCouponCodeStartedEventHandlerFixture struct {
	CacheSpy   *cache_spy.RedisCacheSpy
	UseCaseSpy *ValidateCouponCodeApplicationSpy
}

type ValidateCouponCodeApplicationSpy struct {
	Params        map[string]*validate_coupon_code.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*validate_coupon_code.Output
}

func NewSut() (bus.Handler, *UserCriationWithCouponCodeStartedEventHandlerFixture) {
	cacheSpy := cache_spy.New()
	usecase := &ValidateCouponCodeApplicationSpy{
		Params:        make(map[string]*validate_coupon_code.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*validate_coupon_code.Output)}

	return user_criation_with_coupon_code_started_event_handler.New(usecase, cacheSpy),
		&UserCriationWithCouponCodeStartedEventHandlerFixture{
			CacheSpy:   cacheSpy,
			UseCaseSpy: usecase,
		}
}

func (uc *ValidateCouponCodeApplicationSpy) Execute(input *validate_coupon_code.Input) (*validate_coupon_code.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func GetInvalidUserCriationWithCouponCodeStarted() *event.UserCriationWithCouponCodeStarted {
	return &event.UserCriationWithCouponCodeStarted{}
}
