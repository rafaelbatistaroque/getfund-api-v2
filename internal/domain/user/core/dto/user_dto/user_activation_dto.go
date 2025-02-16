package user_dto

type ActivationUserDto struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	Gender            string `json:"gender"`
	Password          string `json:"password"`
	CountryId         int    `json:"country_id"`
	UserCategoryId    int    `json:"user_category_id"`
	MainSocialNetwork string `json:"main_social_network"`
	RegisteredUrl     string `json:"registered_url"`
	IsAdmin           bool   `json:"is_admin"`
	IsActive          bool   `json:"is_active"`
	RegisteredAt      int64  `json:"registered_at"`
}
