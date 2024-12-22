package signin

import (
	validation "getfund-api-v2/pkg/input_validation"
)

type Input = signinInput

type signinInput struct {
	validation.InputValidation
	Password string
	UserName string
}

func (i *signinInput) Validate() {
	i.Required(i.UserName, "UserName")
	i.Required(i.Password, "Password")
}
