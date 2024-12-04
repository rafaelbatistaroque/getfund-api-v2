package recoverpassword

import (
	validation "getfund-api-v2/internal/pkg/inputvalidation"
)

type Input = recoverPassword

type recoverPasswordInput struct {
	validation.InputValidation
	UserName string `json:"email"`
}
