package payload

type ActivateUserWithCouponConfirmedPayload struct {
	UserId     int    `json:"iser_id"`
	CouponCode string `json:"coupon_code"`
}
