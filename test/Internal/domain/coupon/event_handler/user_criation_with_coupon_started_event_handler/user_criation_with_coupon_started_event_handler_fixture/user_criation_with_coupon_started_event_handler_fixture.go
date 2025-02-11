package user_criation_with_coupon_started_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/coupon/adapter/event_handler/user_criation_with_coupon_started_event_handler"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
	"getfund-api-v2/test/helper/cache_spy"
)

type UserCriationWithCouponStartedEventHandlerFixture struct {
	CacheSpy   *cache_spy.RedisCacheSpy
	UseCaseSpy *ValidateCouponApplicationSpy
}

type ValidateCouponApplicationSpy struct {
	Params        map[string]*validate_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*validate_coupon.Output
}

func NewSut() (bus.Handler, *UserCriationWithCouponStartedEventHandlerFixture) {
	cacheSpy := cache_spy.New()
	usecase := &ValidateCouponApplicationSpy{
		Params:        make(map[string]*validate_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*validate_coupon.Output)}

	return user_criation_with_coupon_started_event_handler.New(usecase, cacheSpy),
		&UserCriationWithCouponStartedEventHandlerFixture{
			CacheSpy:   cacheSpy,
			UseCaseSpy: usecase,
		}
}

func (uc *ValidateCouponApplicationSpy) Execute(input *validate_coupon.Input) (*validate_coupon.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func GetInvalidUserCriationWithCouponStarted() *event.UserCriationWithCouponStarted {
	return &event.UserCriationWithCouponStarted{}
}

func GetValidateCouponInput() *validate_coupon.Input {
	return &validate_coupon.Input{
		CouponCode: "fake-coupon-code",
	}
}

func GetValidUserCriationWithCouponStarted() *event.UserCriationWithCouponStarted {
	payload, _ := json.Marshal(map[string]string{
		"coupon_code": "fake-coupon-code",
	})

	event := &event.UserCriationWithCouponStarted{}
	event.SetPayload(payload)

	return event
}
