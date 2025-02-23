package coupon_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
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

func (r *CouponRepositorySpy) GetCouponByCode(couponCode string) (*prizedraw_dto.CouponDto, error) {
	r.Params["GetCouponByCode:couponCode"] = couponCode

	r.CallsCount["GetCouponByCode"]++

	sucess := r.SuccessResult["GetCouponByCode"]
	if sucess != nil {
		return sucess.(*prizedraw_dto.CouponDto), nil
	}

	return nil, r.ErrorResult["GetCouponByCode"]
}

func (r *CouponRepositorySpy) DefineGetCouponByCodeError() {
	r.ErrorResult["GetCouponByCode"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineGetCouponByCodeSuccess(coupon *prizedraw_dto.CouponDto) {
	if coupon != nil {
		r.SuccessResult["GetCouponByCode"] = coupon
		return
	}

	r.SuccessResult["GetCouponByCode"] = &prizedraw_dto.CouponDto{}
}
