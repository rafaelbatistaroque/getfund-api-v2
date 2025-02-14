package payload

type UserCriationPayload struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
