package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
)

type ApplyPrizeDrawCouponApplicationSpy struct {
	Params        map[string]*apply_prizedraw_coupon.Input
	CallsCount    map[string]int
	ErrorResult   map[string]*result_app.ApplicationError
	SuccessResult map[string]*apply_prizedraw_coupon.Output
}

func NewApplyPrizeDrawCoupon() *ApplyPrizeDrawCouponApplicationSpy {
	return &ApplyPrizeDrawCouponApplicationSpy{
		Params:        make(map[string]*apply_prizedraw_coupon.Input),
		CallsCount:    make(map[string]int),
		ErrorResult:   make(map[string]*result_app.ApplicationError),
		SuccessResult: make(map[string]*apply_prizedraw_coupon.Output)}
}

func (a *ApplyPrizeDrawCouponApplicationSpy) Execute(input *apply_prizedraw_coupon.Input) (*apply_prizedraw_coupon.Output, *result_app.ApplicationError) {
	a.Params["Execute:input"] = input

	a.CallsCount["Execute"]++

	return a.SuccessResult["Execute"], a.ErrorResult["Execute"]
}

func GetApplyCouponInput(from *prizedraw_dto.CouponDto, userId int) *apply_prizedraw_coupon.Input {
	return &apply_prizedraw_coupon.Input{
		CouponId:    from.Id,
		PrizeDrawId: from.PrizeDrawId,
		ProductId:   from.ProductId,
		UserId:      userId,
		IsUserAdmin: false,
	}
}
