package signin

import (
	validation "getfund-api-v2/pkg/inputvalidation"
)

type Input = signinInput

type signinInput struct {
	validation.InputValidation
	Password string
	UserName string
}

func (i *signinInput) Validate() {
	if validation.IsNilOrEmpty(i.UserName) {
		i.AppendError("UserName", validation.Err_Msg_PARAMETER_NOT_EMPTY.Error())
	}

	if validation.IsNilOrEmpty(i.Password) {
		i.AppendError("Password", validation.Err_Msg_PARAMETER_NOT_EMPTY.Error())
	}
}
