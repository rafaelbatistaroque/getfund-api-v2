package validate_prizedraw_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validateCouponInput

type validateCouponInput struct {
	CouponCode          string `json:"coupon_code"`
	Email               string `json:"email"`
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
		).
		ApplyRules(i.Email, "Email",
			&validation.RequiredRule{},
			&validation.EmailRule{},
		)

	return i.rules.GetResult()
}
