package recover_password

import (
	validation "getfund-api-v2/pkg/input_validation"
)

type Input = recoverPasswordInput

type recoverPasswordInput struct {
	validation.InputValidation
	Username string `json:"email"`
}

func (i *recoverPasswordInput) Validate() {
	i.Required(i.Username, "Username")
}
