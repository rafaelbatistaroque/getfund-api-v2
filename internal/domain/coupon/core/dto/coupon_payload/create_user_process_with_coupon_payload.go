package coupon_payload

type CreateUserProcessWithCouponPayload struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
	ErrorStatus    string `json:"error_status"`
}
