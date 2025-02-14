package send_activation_account_mail

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = sendActivationAccountMailInput

type sendActivationAccountMailInput struct {
	rules validation.Rule
}

func (i *sendActivationAccountMailInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
