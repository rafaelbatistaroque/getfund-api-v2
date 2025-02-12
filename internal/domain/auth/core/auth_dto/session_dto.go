package auth_dto

type SessionDto struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   bool   `json:"is_admin"`
}
