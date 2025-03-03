package auth_payload

type CreateUserProcessWithCouponPayload struct {
	ActivationDataKey string `json:"activation_data_key"`
	CouponCode        string `json:"coupon_code"`
}
