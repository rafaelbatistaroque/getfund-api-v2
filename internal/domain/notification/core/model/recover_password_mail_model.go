package notification_model

type RecoverPasswordMailModel struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	RecoveryLink string `json:"recovery_link"`
}
