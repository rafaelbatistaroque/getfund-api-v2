package user_dto

type UserCriationPayloadDto struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
