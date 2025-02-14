package payload

type CreateUserProcessPayload struct {
	ActivationCode string `json:"activation_code"`
	ActivationLink string `json:"activation_link"`
}
