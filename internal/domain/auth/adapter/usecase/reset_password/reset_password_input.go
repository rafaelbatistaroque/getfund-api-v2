package reset_password

type Input = resetPasswordInput

type resetPasswordInput struct {
	RecoveryCode string `json:"code"`
	Password     string `json:"password"`
}

func (i *resetPasswordInput) Validate() {

}
