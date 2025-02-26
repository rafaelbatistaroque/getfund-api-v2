package apply_prizedraw_coupon

import (
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
	return i.rules.GetResult()
}
