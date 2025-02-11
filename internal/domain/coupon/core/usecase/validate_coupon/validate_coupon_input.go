package validate_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validateCouponInput

type validateCouponInput struct {
	rules validation.Rule
}

func (i *validateCouponInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
