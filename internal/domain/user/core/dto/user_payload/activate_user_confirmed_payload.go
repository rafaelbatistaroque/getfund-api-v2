package payload

type ActivateUserConfirmedPayload struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
