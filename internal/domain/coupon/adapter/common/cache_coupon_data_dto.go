package coupon_common

type CacheCouponData struct {
	IsValid    bool        `json:"is_valid"`
	CouponData *CouponData `json:"coupon_data"`
	Error      *ErrorData  `json:"error"`
}

type CouponData struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     uint64 `json:"start_at"`
	EndAt       int    `json:"end_at"`
	Discount    int    `json:"discount"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
