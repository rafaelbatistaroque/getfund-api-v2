package create_user_process_with_coupon_started_event_handler_fixture

import (
	"encoding/json"
	"errors"
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	coupon_payload "getfund-api-v2/internal/domain/coupon/adapter/event_handler/create_user_process_with_coupon_started_event_handler"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
)

type CreateUserWithCouponStartedEventHandlerFixture struct {
	CacheSpy   *cache_spy.RedisCacheSpy
	UseCaseSpy *ValidateCouponApplicationSpy
}

type ValidateCouponApplicationSpy struct {
	Params        map[string]*validate_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*validate_coupon.Output
}

func NewSut() (bus.Handler, *CreateUserWithCouponStartedEventHandlerFixture) {
	cacheSpy := cache_spy.New()
	usecase := &ValidateCouponApplicationSpy{
		Params:        make(map[string]*validate_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*validate_coupon.Output)}

	return coupon_payload.New(usecase, cacheSpy),
		&CreateUserWithCouponStartedEventHandlerFixture{
			CacheSpy:   cacheSpy,
			UseCaseSpy: usecase,
		}
}

func (uc *ValidateCouponApplicationSpy) Execute(input *validate_coupon.Input) (*validate_coupon.Output, *result_app.ApplicationError) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *ValidateCouponApplicationSpy) DefineValidateCouponUsecaseError(withMessage string) {
	uc.ErrorResult["Execute"] = &result_app.ApplicationError{Message: errors.New(withMessage)}
}

func (uc *ValidateCouponApplicationSpy) DefineValidateCouponUsecaseSuccess() {
	uc.SuccessResult["Execute"] = &validate_coupon.Output{}
}

func GetInvalidCreateUserProcessWithCouponStartedEvent() *create_user.CreateUserProcessWithCouponStartedEvent {
	return &create_user.CreateUserProcessWithCouponStartedEvent{}
}

func GetValidateCouponInput() *validate_coupon.Input {
	return &validate_coupon.Input{
		CouponCode: "fake-coupon-code",
	}
}

func GetCachedCouponData(code, message string, isValid bool) coupon_common.CacheCouponData {
	return coupon_common.CacheCouponData{
		IsValid:      isValid,
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func GetValidCreateUserProcessWithCouponStartedEvent() *create_user.CreateUserProcessWithCouponStartedEvent {
	payload, _ := json.Marshal(map[string]string{
		"coupon_code":         "fake-coupon-code",
		"activation_data_key": "fake-activation-data-key",
	})

	event := &create_user.CreateUserProcessWithCouponStartedEvent{}
	event.SetPayload(payload)

	return event
}
