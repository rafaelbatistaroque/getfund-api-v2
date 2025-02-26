package prizedraw_dto

type CouponDto struct {
	Id                      int                         `json:"id"`
	Code                    string                      `json:"code"`
	PrizeDrawId             int                         `json:"prize_draw_id"`
	ProductId               int                         `json:"product_id"`
	CouponTypeApplicability *CouponTypeApplicabilityDto `json:"coupon_type_applicability"`
	UserCouponApplies       []UserCouponApplyDto        `json:"user_coupon_aplies"`
}
