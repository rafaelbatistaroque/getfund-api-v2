package prizedraw_contract

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
)

type Repository interface {
	CouponRepository
	GetPrizeDrawById(id int) (*prizedraw_dto.PrizeDrawDto, error)
}

type CouponRepository interface {
	GetCouponByCode(couponCode string) (*prizedraw_dto.CouponDto, error)
	GetCouponById(couponId int) (*prizedraw_dto.CouponDto, error)
	SaveEntranceWithCouponApplied(entrance *prizedraw_dto.EntranceDto, coupon *prizedraw_dto.CouponDto) error
}
