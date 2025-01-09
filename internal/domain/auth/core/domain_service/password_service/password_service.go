package password_service

type PasswordService interface {
	IsStrongPassword(password string) bool
}
