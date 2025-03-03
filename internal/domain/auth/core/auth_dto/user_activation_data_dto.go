package auth_dto

type ActivationUserData struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CouponCode string `json:"cupon_code"`
}
