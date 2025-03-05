package validate_prizedraw_coupon

type Output = validatePrizeDrawCouponOutput

type validatePrizeDrawCouponOutput struct {
	Message     string `json:"message"`
	CouponId    int    `json:"coupon_id"`
	PrizeDrawId int    `json:"prize_draw_id"`
	ProductId   int    `json:"product_id"`
}
