package coupon_dto

type UserCriationWithCouponPayloadDto struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
	Status         string `json:"status"`
}
