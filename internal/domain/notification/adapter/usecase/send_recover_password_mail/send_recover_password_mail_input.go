package send_recover_password_mail

import validation "getfund-api-v2/pkg/input_validation"

type Input = sendRecoverPasswordMailInput

type sendRecoverPasswordMailInput struct {
	validation.InputValidation
	KeyCache string
}

func (i *sendRecoverPasswordMailInput) Validate() {
	i.Required(i.KeyCache, "KeyCache")
}
