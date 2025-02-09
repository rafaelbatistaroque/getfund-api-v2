package user_dto

type UserActivationWithCouponPayloadDto struct {
	UserId         string `json:"iser_id"`
	ActivationCode string `json:"activation_code"`
}
