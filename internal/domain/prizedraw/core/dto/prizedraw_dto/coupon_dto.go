package prizedraw_dto

type CouponDto struct {
	Id                int               `json:"id"`
	Code              string            `json:"code"`
	TypeApplicability int               `json:"type_applicability"`
	PrizeDrawId       int               `json:"prize_draw_id"`
	ProductId         int               `json:"product_id"`
	StartAt           int64             `json:"start_at"`
	EndAt             *int64            `json:"end_at"`
	Discount          int               `json:"discount"`
	LimitApplication  *int              `json:"limit_application"`
	UserCouponApplies []UserCouponApply `json:"user_coupon_aplies"`
}

type UserCouponApply struct {
	UserId int `json:"user_id"`
}
