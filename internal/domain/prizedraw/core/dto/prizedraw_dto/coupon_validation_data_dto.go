package prizedraw_dto

type CouponValidationData struct {
	Product   *ProductData   `json:"product"`
	PrizeDraw *PrizeDrawData `json:"prize_draw"`
}

type ProductData struct {
	Id               int  `json:"id"`
	TotalPrice       int  `json:"total_price"`
	IsActive         bool `json:"is_active"`
	EntranceQuantity int  `json:"entrance_qty"`
}

type PrizeDrawData struct {
	Id               int  `json:"id"`
	WinnerEntranceId *int `json:"winner_entrance_id"`
}
