package activate_user

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = activateUserInput

type activateUserInput struct {
	ActivationCode string `json:"activation_code"`

	rules validation.Rule
}

func (i *activateUserInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.ActivationCode, "ActivationCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Exactly: 20},
		)

	return i.rules.GetResult()
}
