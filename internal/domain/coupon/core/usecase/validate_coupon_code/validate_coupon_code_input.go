package validate_coupon_code

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validateCouponCodeInput

type validateCouponCodeInput struct {
	rules validation.Rule
}

func (i *validateCouponCodeInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
