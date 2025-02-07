package reset_password

import "github.com/rafaelbatistaroque/validation"

type Input = resetPasswordInput

type resetPasswordInput struct {
	RecoveryCode string `json:"code"`
	RecoveryKey  string
	Password     string `json:"password"`

	rules validation.Rule
}

func (i *resetPasswordInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.RecoveryCode, "RecoveryCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Exactly: 64},
		).
		ApplyRules(i.Password, "Password",
			&validation.RequiredRule{},
			&validation.PasswordRule{MinLength: 8, RequireUpper: true, RequireLower: true, RequireDigit: true},
		).
		ApplyRules(i.RecoveryKey, "RecoveryKey",
			&validation.RequiredRule{},
			&validation.EqualRule{Equal: "recovery_password_" + i.RecoveryCode},
		)

	return i.rules.GetResult()
}
