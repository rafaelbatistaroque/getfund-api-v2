package dto

type ForgetPasswordDto struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	RecoveryLink string `json:"recovery_link"`
}
