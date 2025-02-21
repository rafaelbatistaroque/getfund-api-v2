package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	"errors"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
)

type ValidateCouponApplicationSpy struct {
	Params        map[string]*validate_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*validate_coupon.Output
}

func NewValidateCoupon() *ValidateCouponApplicationSpy {
	return &ValidateCouponApplicationSpy{
		Params:        make(map[string]*validate_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*validate_coupon.Output)}
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

func GetValidateCouponInput() *validate_coupon.Input {
	return &validate_coupon.Input{
		CouponCode:  "fake-coupon-code",
		PrizeDrawId: 1,
		ProductId:   1,
	}
}
