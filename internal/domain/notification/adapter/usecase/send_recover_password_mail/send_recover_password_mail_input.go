package send_recover_password_mail

import validation "github.com/rafaelbatistaroque/validation"

type Input = sendRecoverPasswordMailInput

type sendRecoverPasswordMailInput struct {
	KeyCache string

	rules validation.Rule
}

func (i *sendRecoverPasswordMailInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.KeyCache, "KeyCache", &validation.RequiredRule{})

	return i.rules.GetResult()
}
