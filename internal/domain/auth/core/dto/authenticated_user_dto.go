package dto

type AuthenticatedUserDto struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   bool   `json:"is_admin"`
	Password  string `json:"password"`
}
