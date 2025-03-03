package signin

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = signinInput

type signinInput struct {
	Password string
	Username string

	rules validation.Rule
}

func (i *signinInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.Username, "UserName", &validation.RequiredRule{}).
		ApplyRules(i.Password, "Password", &validation.RequiredRule{})

	return i.rules.GetResult()
}
