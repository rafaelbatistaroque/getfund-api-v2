package payload

type UserCriationWithCouponPayload struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
}
