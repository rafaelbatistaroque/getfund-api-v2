package auth_dto

type AuthenticatedUserDto struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	IsAdmin   int    `json:"is_admin"`
	Password  string `json:"password"`
}
