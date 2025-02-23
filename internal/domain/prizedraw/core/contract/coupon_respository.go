package coupon_contract

import "getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"

type Repository interface {
	GetCouponByCode(couponCode string) (*prizedraw_dto.CouponDto, error)
}
