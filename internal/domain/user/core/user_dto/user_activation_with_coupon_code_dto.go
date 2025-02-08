package user_dto

type UserActivationWithCouponDto struct {
	UserId         string `json:"iser_id"`
	ActivationCode string `json:"activation_code"`
}
