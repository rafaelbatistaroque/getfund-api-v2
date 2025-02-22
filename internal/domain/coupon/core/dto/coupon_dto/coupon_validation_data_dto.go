package coupon_dto

type CouponValidationData struct {
	Product   *ProductData   `json:"product"`
	PrizeDraw *PrizeDrawData `json:"prize_draw"`
}

type ProductData struct {
	TotalPrice int  `json:"total_price"`
	IsActive   bool `json:"is_active"`
}

type PrizeDrawData struct {
	WinnerEntranceId *int `json:"winner_entrance_id"`
}
