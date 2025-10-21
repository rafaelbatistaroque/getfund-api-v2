package dto

type ActivationUserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
