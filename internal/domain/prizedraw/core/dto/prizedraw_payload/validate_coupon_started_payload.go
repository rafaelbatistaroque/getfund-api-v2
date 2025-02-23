package prizedraw_payload

type ValidateCouponStartedPayload struct {
	ProductId   int `json:"product_id"`
	PrizeDrawId int `json:"prize_draw_id"`
}
