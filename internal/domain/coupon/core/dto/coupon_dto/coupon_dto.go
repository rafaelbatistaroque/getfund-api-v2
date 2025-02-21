package coupon_dto

type CouponDto struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     int64  `json:"start_at"`
	EndAt       *int64 `json:"end_at"`
	Discount    int    `json:"discount"`
}
