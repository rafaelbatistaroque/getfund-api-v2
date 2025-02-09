package activate_user

type Output = ActivateUserOutput

type ActivateUserOutput struct {
	Token   string        `json:"token"`
	Session SessionOutput `json:"session"`
}

type SessionOutput struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   bool   `json:"is_admin"`
}
