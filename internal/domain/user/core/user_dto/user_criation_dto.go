package user_dto

type UserCriationDto struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
