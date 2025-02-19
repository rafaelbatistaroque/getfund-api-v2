package apply_coupon

import (
	validation "github.com/rafaelbatistaroque/validation"
)

type Input = applyCouponInput

type applyCouponInput struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     uint64 `json:"start_at"`
	EndAt       int    `json:"end_at"`
	Discount    int    `json:"discount"`
	UserId      int    `json:"user_id"`

	rules validation.Rule
}

func (i *applyCouponInput) Validate() validation.Validatable {
	return i.rules.GetResult()
}
