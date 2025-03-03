package activate_user

// Apenas mensagem
type Output = ActivateUserOutput

type ActivateUserOutput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
