package coupon_dto

type CouponDto struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     uint64 `json:"start_at"`
	EndAt       int    `json:"end_at"`
	Discount    int    `json:"discount"`
}
