package validate_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validateCouponInput

type validateCouponInput struct {
	CouponCode  string `json:"coupon_code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`

	rules validation.Rule
}

func (i *validateCouponInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
