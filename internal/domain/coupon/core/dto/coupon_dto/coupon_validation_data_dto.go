package coupon_dto

type CouponValidationData struct {
	Product   *ProductData
	PrizeDraw *PrizeDrawData
}

type ProductData struct {
	TotalPrice int  `json:"total_price"`
	IsActive   bool `json:"is_active"`
}

type PrizeDrawData struct {
	WinnerEntranceId int `json:"winner_entrance_id"`
}
