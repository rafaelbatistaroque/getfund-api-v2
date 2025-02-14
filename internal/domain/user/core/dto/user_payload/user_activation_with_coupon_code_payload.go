package payload

type UserActivationWithCouponPayload struct {
	UserId         int    `json:"iser_id"`
	ActivationCode string `json:"activation_code"`
}
