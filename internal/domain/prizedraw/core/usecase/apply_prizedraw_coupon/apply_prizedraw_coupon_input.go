package apply_prizedraw_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = applyCouponInput

type applyCouponInput struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     int64  `json:"start_at"`
	EndAt       *int64 `json:"end_at"`
	Discount    int    `json:"discount"`
	UserId      int    `json:"user_id"`

	rules validation.Rule
}

func (i *applyCouponInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
