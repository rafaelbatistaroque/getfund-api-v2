package user_dto

type UserCriationStartedDto struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
