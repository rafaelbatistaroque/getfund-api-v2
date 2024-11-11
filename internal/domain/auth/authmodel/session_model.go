package authmodel

type SessionModel struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   int    `json:"is_admin"`
}
