package apply_prizedraw_coupon

import (
	"fmt"

	validation "github.com/rafaelbatistaroque/validation"
)

type Input = applyCouponInput

type applyCouponInput struct {
	CouponId    int `json:"coupon_id"`
	PrizeDrawId int `json:"prize_draw_id"`
	ProductId   int `json:"product_id"`
	UserId      int `json:"user_id"`

	rules validation.Rule
}

func (i *applyCouponInput) Validate() validation.Validatable {
	if i.CouponId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "CouponId"))
	}

	if i.PrizeDrawId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "PrizeDrawId"))
	}

	if i.ProductId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "ProductId"))
	}

	if i.UserId <= 0 {
		i.rules.AddError(fmt.Errorf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "UserId"))
	}

	return i.rules.GetResult()
}
