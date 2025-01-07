package recover_password

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = recoverPasswordInput

type recoverPasswordInput struct {
	Username string `json:"email"`

	rules validation.Rule
}

func (i *recoverPasswordInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.Username, "Username", &validation.RequiredRule{})

	return i.rules.GetResult()
}
