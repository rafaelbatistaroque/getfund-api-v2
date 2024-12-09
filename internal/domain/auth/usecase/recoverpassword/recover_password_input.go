package recoverpassword

import (
	validation "getfund-api-v2/pkg/inputvalidation"
)

type Input = recoverPasswordInput

type recoverPasswordInput struct {
	validation.InputValidation
	Username string `json:"email"`
}

func (i *recoverPasswordInput) Validate() {
	if validation.IsNilOrEmpty(i.Username) {
		i.AppendError("Username", validation.Err_Msg_PARAMETER_NOT_EMPTY.Error())
	}
}
