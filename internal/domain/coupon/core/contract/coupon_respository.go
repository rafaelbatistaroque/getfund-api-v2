package coupon_contract

import coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"

type Repository interface {
	GetCouponByCode(couponCode string) (*coupon_dto.CouponDto, error)
}
