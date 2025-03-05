package prizedraw_payload

type ApplyPrizeDrawCouponStartedPayload struct {
	UserId       int `json:"user_id"`
	CouponId     int `json:"coupon_id"`
	ProductId    int `json:"product_id"`
	PrizeDrawId  int `json:"prize_draw_id"`
	ItemQuantity int `json:"item_quantity"`
}
