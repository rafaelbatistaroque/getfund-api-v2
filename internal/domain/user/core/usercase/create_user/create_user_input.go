package create_user

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = createUserInput

type createUserInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Gender            string `json:"gender"`
	Password          string `json:"password"`
	CountryId         int    `json:"country_id"`
	UserCategoryId    int    `json:"user_category_id"`
	MainSocialNetwork string `json:"main_social_network"`
	RegisteredUrl     string `json:"registered_url"`
	CuponCode         string `json:"cupon_code"`

	rules validation.Rule
}

func (i *createUserInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.FirstName, "FirstName",
			&validation.RequiredRule{},
			&validation.LengthRule{Min: 1, Max: 50},
		).
		ApplyRules(i.LastName, "LastName",
			&validation.RequiredRule{},
			&validation.LengthRule{Min: 1, Max: 50},
		).
		ApplyRules(i.Email, "Email",
			&validation.EmailRule{},
		).
		ApplyRules(i.Gender, "Gender",
			&validation.RequiredRule{},
			&validation.ListRule{Contains: []string{"f", "m", "u", "nb"}},
		).
		ApplyRules(i.Password, "Password",
			&validation.RequiredRule{},
			&validation.PasswordRule{MinLength: 8, RequireLower: true, RequireUpper: true, RequireDigit: true},
		).
		ApplyRules(i.MainSocialNetwork, "MainSocialNetwork",
			&validation.RequiredRule{},
			&validation.SocialURLRule{},
		).
		ApplyRules(i.RegisteredUrl, "RegisteredUrl",
			&validation.RequiredRule{},
		)

	return i.rules.GetResult()
}
