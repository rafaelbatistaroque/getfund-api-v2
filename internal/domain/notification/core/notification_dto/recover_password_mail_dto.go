package notification_dto

type RecoverPasswordMailDto struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	RecoveryLink string `json:"recovery_link"`
}
