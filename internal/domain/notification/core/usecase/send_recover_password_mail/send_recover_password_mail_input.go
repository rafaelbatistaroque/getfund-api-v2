package send_recover_password_mail

import validation "github.com/rafaelbatistaroque/validation"

type Input = sendRecoverPasswordMailInput

type sendRecoverPasswordMailInput struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	RecoveryLink string `json:"recovery_link"`

	rules validation.Rule
}

func (i *sendRecoverPasswordMailInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.Username, "Username", &validation.RequiredRule{}).
		ApplyRules(i.FirstName, "FirstName", &validation.RequiredRule{}).
		ApplyRules(i.RecoveryLink, "RecoveryLink", &validation.RequiredRule{})

	return i.rules.GetResult()
}
