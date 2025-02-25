package prizedraw_dto

type ValidationCouponData struct {
	Product    *ProductData `json:"product"`
	CouponDara *CouponData  `json:"coupon_data"`
}

type ProductData struct {
	Id               int  `json:"id"`
	TotalPrice       int  `json:"total_price"`
	IsActive         bool `json:"is_active"`
	EntranceQuantity int  `json:"entrance_qty"`
}

type CouponData struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
	StartAt     int64  `json:"start_at"`
	EndAt       *int64 `json:"end_at"`
	Discount    int    `json:"discount"`
}
