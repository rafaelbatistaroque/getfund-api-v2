package reset_password

import "github.com/rafaelbatistaroque/validation"

type Input = resetPasswordInput

type resetPasswordInput struct {
	RecoveryCode string `json:"code"`
	Password     string `json:"password"`

	rules validation.Rule
}

func (i *resetPasswordInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.RecoveryCode, "RecoveryCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Length: 64})

	return i.rules.GetResult()
}
