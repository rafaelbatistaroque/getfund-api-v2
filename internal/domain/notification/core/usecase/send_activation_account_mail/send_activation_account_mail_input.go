package send_activation_account_mail

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = sendActivationAccountMailInput

type sendActivationAccountMailInput struct {
	FirstName      string `json:"first_name"`
	Email          string `json:"email"`
	ActivationLink string `json:"activation_link"`

	rules validation.Rule
}

func (i *sendActivationAccountMailInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.FirstName, "FirstName",
			&validation.RequiredRule{},
		).
		ApplyRules(i.Email, "Email",
			&validation.RequiredRule{},
		)

	return i.rules.GetResult()
}
