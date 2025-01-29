package user_dto

type UserCriationWithCouponDto struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
}
