package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	"errors"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	shared_error "getfund-api-v2/internal/shared/error"
)

type ValidatePrizeDrawCouponApplicationSpy struct {
	Params        map[string]*validate_prizedraw_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*shared_error.Error
	SuccessResult map[string]*validate_prizedraw_coupon.Output
}

func NewValidatePrizeDrawCoupon() *ValidatePrizeDrawCouponApplicationSpy {
	return &ValidatePrizeDrawCouponApplicationSpy{
		Params:        make(map[string]*validate_prizedraw_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*shared_error.Error),
		SuccessResult: make(map[string]*validate_prizedraw_coupon.Output)}
}

func (uc *ValidatePrizeDrawCouponApplicationSpy) Execute(input *validate_prizedraw_coupon.Input) (*validate_prizedraw_coupon.Output, *shared_error.Error) {
	uc.Params["Execute:input"] = input

	uc.CallsCount["Execute"]++

	return uc.SuccessResult["Execute"], uc.ErrorResult["Execute"]
}

func (uc *ValidatePrizeDrawCouponApplicationSpy) DefineValidateCouponUsecaseError(withMessage string) {
	uc.ErrorResult["Execute"] = &shared_error.Error{Message: errors.New(withMessage)}
}

func (uc *ValidatePrizeDrawCouponApplicationSpy) DefineValidateCouponUsecaseSuccess() {
	uc.SuccessResult["Execute"] = &validate_prizedraw_coupon.Output{}
}

func (uc *ValidatePrizeDrawCouponApplicationSpy) DefineValidateCouponUsecaseSuccessNull() {
	uc.SuccessResult["Execute"] = nil
}

func (uc *ValidatePrizeDrawCouponApplicationSpy) DefineValidateCouponUsecaseSuccessWithOutput(output *validate_prizedraw_coupon.Output) {
	uc.SuccessResult["Execute"] = output
}

func GetValidateCouponInput() *validate_prizedraw_coupon.Input {

	return &validate_prizedraw_coupon.Input{
		CouponCode:          "fake-coupon-code",
		SelectedProductId:   10,
		SelectedPrizeDrawId: 5,
		UserId:              1,
		Email:               "fake-email",
	}
}
