package user_dto

type UserActivationWithCouponPayloadDto struct {
	UserId         int    `json:"iser_id"`
	ActivationCode string `json:"activation_code"`
}
