package signout

import inputvalidation "getfund-api-v2/pkg/input_validation"

type Input = signoutInput

type signoutInput struct {
	inputvalidation.InputValidation
	Token string
}

func (i *signoutInput) Validate() {
	if inputvalidation.IsNilOrEmpty(i.Token) {
		i.AppendError("Token", inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error())
	}
}
