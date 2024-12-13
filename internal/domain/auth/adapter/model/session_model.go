package auth_model

type SessionModel struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   int    `json:"is_admin"`
}
