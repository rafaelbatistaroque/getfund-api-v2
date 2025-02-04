package activate_user

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = activateUserInput

type activateUserInput struct {
	rules validation.Rule
}

func (i *activateUserInput) Validate() validation.Validatable {
	return nil
}
