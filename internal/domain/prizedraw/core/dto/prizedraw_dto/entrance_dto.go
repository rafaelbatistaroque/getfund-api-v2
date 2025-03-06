package prizedraw_dto

type EntranceDto struct {
	Id          int    `json:"id"`
	LuckyCode   string `json:"lucky_number"`
	UserId      int    `json:"user_id"`
	PrizeDrawId int    `json:"prize_draw_id"`
	PurchaseId  int    `json:"purchase_id"`
	IsDonation  bool   `json:"is_donation"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
