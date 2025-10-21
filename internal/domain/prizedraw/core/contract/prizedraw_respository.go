package contract

import (
	"getfund-api-v2/internal/domain/prizedraw/core/dto"
)

type Repository interface {
	CouponRepository
	GetPrizeDrawById(id int) (*dto.PrizeDrawDto, error)
}

type CouponRepository interface {
	GetCouponByCode(couponCode string) (*dto.CouponDto, error)
	GetCouponById(couponId int) (*dto.CouponDto, error)
	SaveEntranceWithCouponApplied(entrance *dto.EntranceDto, coupon *dto.CouponDto) error
}
