package user_dto

type UserCriationStartedDto struct {
	ActivationCode string `json:"activation_code"`
	FirstName      string `json:"first_name"`
	Email          string `json:"email"`
}
