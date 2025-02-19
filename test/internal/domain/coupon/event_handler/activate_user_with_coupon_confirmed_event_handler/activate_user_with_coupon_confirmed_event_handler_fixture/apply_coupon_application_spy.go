package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	"getfund-api-v2/internal/domain/coupon/core/usecase/apply_coupon"
	"getfund-api-v2/internal/shared/result_app"
)

type ApplyCouponApplicationSpy struct {
	Params        map[string]*apply_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*apply_coupon.Output
}

func NewApplyCoupon() *ApplyCouponApplicationSpy {
	return &ApplyCouponApplicationSpy{
		Params:        make(map[string]*apply_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*apply_coupon.Output)}
}

func (a *ApplyCouponApplicationSpy) Execute(input *apply_coupon.Input) (*apply_coupon.Output, *result_app.ApplicationError) {
	a.Params["Execute:input"] = input

	a.CallsCount["Execute"]++

	return a.SuccessResult["Execute"], a.ErrorResult["Execute"]
}

func GetApplyCouponInput(from coupon_common.CouponData, userId int) *apply_coupon.Input {
	return &apply_coupon.Input{
		Id:          from.Id,
		Code:        from.Code,
		PrizeDrawId: from.PrizeDrawId,
		ProductId:   from.ProductId,
		StartAt:     from.StartAt,
		EndAt:       from.EndAt,
		Discount:    from.Discount,
		UserId:      userId,
	}
}
