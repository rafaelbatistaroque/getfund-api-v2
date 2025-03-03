package auth_payload

type ActivateUserWithCouponConfirmedPayload struct {
	UserId     int    `json:"iser_id"`
	Email      string `json:"email"`
	CouponCode string `json:"coupon_code"`
}
