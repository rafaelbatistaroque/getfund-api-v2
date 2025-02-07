package activate_user

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = activateUserInput

type activateUserInput struct {
	ActivationCode string `json:"activation_code"`
	ActivationKey  string

	rules validation.Rule
}

func (i *activateUserInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.ActivationCode, "ActivationCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Exactly: 20},
		).
		ApplyRules(i.ActivationKey, "ActivationKey",
			&validation.RequiredRule{},
			&validation.EqualRule{Equal: "user_activation_" + i.ActivationCode},
		)

	return i.rules.GetResult()
}
