package signout

import inputvalidation "getfund-api-v2/pkg/input_validation"

type Input = signoutInput

type signoutInput struct {
	inputvalidation.InputValidation
	Token string
}

func (i *signoutInput) Validate() {
	i.Required(i.Token, "Token")
}
