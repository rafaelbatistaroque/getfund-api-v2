package signup

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = SignupInput

type SignupInput struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Username             string `json:"username"` //user email
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	CouponCode           string `json:"cupon_code"`

	rules validation.Rule
}

func (i *SignupInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.FirstName, "FirstName",
			&validation.RequiredRule{},
			&validation.LengthRule{Min: 1, Max: 50},
		).
		ApplyRules(i.LastName, "LastName",
			&validation.RequiredRule{},
			&validation.LengthRule{Min: 1, Max: 50},
		).
		ApplyRules(i.Username, "Username",
			&validation.RequiredRule{},
			&validation.EmailRule{},
		).
		ApplyRules(i.Password, "Password",
			&validation.RequiredRule{},
			&validation.PasswordRule{MinLength: 8, RequireLower: true, RequireUpper: true, RequireDigit: true},
		).
		ApplyRules(i.PasswordConfirmation, "PasswordConfirmation",
			&validation.RequiredRule{},
			&validation.EqualRule{Equal: i.Password},
		)

	if i.CouponCode != "" {
		i.rules.ApplyRules(i.CouponCode, "CouponCode",
			&validation.LengthRule{Exactly: 8})
	}

	return i.rules.GetResult()
}
