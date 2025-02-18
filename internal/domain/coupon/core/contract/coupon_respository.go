package coupon_contract

import coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto"

type Repository interface {
	GetCouponByCouponCode(couponCode string) (*coupon_dto.CouponDto, error)
}
