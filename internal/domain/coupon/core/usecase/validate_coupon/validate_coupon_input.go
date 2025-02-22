package validate_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validateCouponInput

type validateCouponInput struct {
	CouponCode          string `json:"coupon_code"`
	SelectedProductId   int    `json:"selected_product_id"`
	SelectedPrizeDrawId int    `json:"selected_prize_draw_id"`
	UserId              int    `json:"user_id"`

	rules validation.Rule
}

func (i *validateCouponInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.CouponCode, "CouponCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Exactly: 8},
		)

	return i.rules.GetResult()
}
