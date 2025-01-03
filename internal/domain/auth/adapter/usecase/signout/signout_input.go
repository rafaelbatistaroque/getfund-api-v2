package signout

import "github.com/rafaelbatistaroque/validation"

type Input = signoutInput

type signoutInput struct {
	Token string

	rules validation.Rule
}

func (i *signoutInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.Token, "Token", &validation.RequiredRule{})

	return i.rules.GetResult()
}
