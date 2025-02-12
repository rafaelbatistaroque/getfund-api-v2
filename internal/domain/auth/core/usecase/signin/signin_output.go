package signin

type Output = SigninOutput

type SigninOutput struct {
	Token   string        `json:"token"`
	Session SessionOutput `json:"session"`
}

type SessionOutput struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   bool   `json:"is_admin"`
}
