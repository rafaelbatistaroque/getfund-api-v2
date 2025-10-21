package prizedraw_repository_spy

import (
	"errors"
	"getfund-api-v2/internal/domain/prizedraw/core/dto"
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

func (r *CouponRepositorySpy) GetCouponByCode(couponCode string) (*dto.CouponDto, error) {
	r.Params["GetCouponByCode:couponCode"] = couponCode

	r.CallsCount["GetCouponByCode"]++

	sucess := r.SuccessResult["GetCouponByCode"]
	if sucess != nil {
		return sucess.(*dto.CouponDto), nil
	}

	return nil, r.ErrorResult["GetCouponByCode"]
}

func (r *CouponRepositorySpy) GetCouponById(couponId int) (*dto.CouponDto, error) {
	r.Params["GetCouponById:couponId"] = couponId

	r.CallsCount["GetCouponById"]++

	sucess := r.SuccessResult["GetCouponById"]
	if sucess != nil {
		return sucess.(*dto.CouponDto), nil
	}

	return nil, r.ErrorResult["GetCouponById"]
}

func (r *CouponRepositorySpy) DefineGetCouponByCodeError() {
	r.ErrorResult["GetCouponByCode"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineGetCouponByCodeSuccess(coupon *dto.CouponDto) {
	if coupon != nil {
		r.SuccessResult["GetCouponByCode"] = coupon
		return
	}

	r.SuccessResult["GetCouponByCode"] = &dto.CouponDto{}
}

func (r *CouponRepositorySpy) DefineGetCouponByIdError() {
	r.ErrorResult["GetCouponById"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineGetCouponByIdSuccess(coupon *dto.CouponDto) {
	if coupon != nil {
		r.SuccessResult["GetCouponById"] = coupon
		return
	}

	r.SuccessResult["GetCouponById"] = &dto.CouponDto{}
}

func (r *CouponRepositorySpy) GetPrizeDrawById(id int) (*dto.PrizeDrawDto, error) {
	r.Params["GetPrizeDrawById:id"] = id

	r.CallsCount["GetPrizeDrawById"]++

	sucess := r.SuccessResult["GetPrizeDrawById"]
	if sucess != nil {
		return sucess.(*dto.PrizeDrawDto), nil
	}

	return nil, r.ErrorResult["GetPrizeDrawById"]
}

func (r *CouponRepositorySpy) DefineGetPrizeDrawByIdError() {
	r.ErrorResult["GetPrizeDrawById"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineGetPrizeDrawByIdSuccess(prizeDraw *dto.PrizeDrawDto) {
	if prizeDraw != nil {
		r.SuccessResult["GetPrizeDrawById"] = prizeDraw
		return
	}

	r.SuccessResult["GetPrizeDrawById"] = &dto.PrizeDrawDto{}
}

func (r *CouponRepositorySpy) SaveEntranceWithCouponApplied(entrance *dto.EntranceDto, coupon *dto.CouponDto) error {
	r.Params["SaveEntranceWithCouponApplied:entrance"] = entrance
	r.Params["SaveEntranceWithCouponApplied:coupon"] = coupon

	r.CallsCount["SaveEntranceWithCouponApplied"]++

	return r.ErrorResult["SaveEntranceWithCouponApplied"]
}

func (r *CouponRepositorySpy) DefineSaveEntranceWithCouponAppliedError() {
	r.ErrorResult["SaveEntranceWithCouponApplied"] = errors.New("fake-error")
}

func (r *CouponRepositorySpy) DefineSaveEntranceWithCouponAppliedSuccess() {
	r.ErrorResult["SaveEntranceWithCouponApplied"] = nil
}
