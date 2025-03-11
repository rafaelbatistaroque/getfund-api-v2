package validate_prizedraw_coupon

import (
	"fmt"

	validation "github.com/rafaelbatistaroque/validation"
)

type Input = validatePrizeDrawCouponInput

type validatePrizeDrawCouponInput struct {
	CouponCode          string `json:"coupon_code"`
	Email               string `json:"email"`
	SelectedProductId   int    `json:"selected_product_id"`
	SelectedPrizeDrawId int    `json:"selected_prize_draw_id"`
	UserId              int    `json:"user_id"`

	rules validation.Rule
}

func (i *validatePrizeDrawCouponInput) Validate() validation.Validatable {
	i.rules.
		ApplyRules(i.CouponCode, "CouponCode",
			&validation.RequiredRule{},
			&validation.LengthRule{Exactly: 8},
		).
		ApplyRules(i.Email, "Email",
			&validation.RequiredRule{},
			&validation.EmailRule{},
		)

	if i.SelectedProductId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "SelectedProductId"))
	}

	if i.SelectedPrizeDrawId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "SelectedPrizeDrawId"))
	}

	if i.UserId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "UserId"))
	}

	return i.rules.GetResult()
}
