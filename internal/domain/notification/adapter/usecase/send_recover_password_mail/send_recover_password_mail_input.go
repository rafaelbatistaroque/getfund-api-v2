package send_recover_password_mail

import inputvalidation "getfund-api-v2/pkg/input_validation"

type Input = sendRecoverPasswordMailInput

type sendRecoverPasswordMailInput struct {
	inputvalidation.InputValidation
	KeyCache string
}

func (i *sendRecoverPasswordMailInput) Validate() {

}
