package user_dto

type UserCriationWithCouponStardedDto struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
}
