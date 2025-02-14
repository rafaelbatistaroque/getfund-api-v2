package payload

type CreateUserProcessPayload struct {
	ActivationDataKey string `json:"activation_data_key"`
	ActivationCode    string `json:"activation_code"`
	ActivationLink    string `json:"activation_link"`
}
