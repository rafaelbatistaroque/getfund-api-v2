package coupon_repository_spy

import (
	"errors"
	coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto"
)

type CouponRepositorySpy struct {
	Params        map[string]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *CouponRepositorySpy {

	return &CouponRepositorySpy{
		Params:        make(map[string]any, 1),
		ErrorResult:   make(map[string]error),
		SuccessResult: make(map[string]any, 1),
		CallsCount:    make(map[string]int, 1)}
}

func (r *CouponRepositorySpy) GetCouponByCouponCode(couponCode string) (*coupon_dto.CouponDto, error) {
	r.Params["GetCouponByCouponCode:couponCode"] = couponCode

	r.CallsCount["GetCouponByCouponCode"]++

	sucess := r.SuccessResult["GetCouponByCouponCode"]
	if sucess != nil {
		return sucess.(*coupon_dto.CouponDto), nil
	}

	return nil, r.ErrorResult["GetCouponByCouponCode"]
}

func (r *CouponRepositorySpy) DefineGetCouponByCouponCodeError() {
	r.ErrorResult["GetCouponByCouponCode"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineGetCouponByCouponCodeSuccess() {
	r.SuccessResult["GetCouponByCouponCode"] = &coupon_dto.CouponDto{}
}
