package prizedraw_dto

type CouponTypeApplicabilityDto struct {
	Id               int     `json:"id"`
	CouponTypeCode   string  `json:"coupon_type_code"`
	StartAt          int64   `json:"start_at"`
	EndAt            *int64  `json:"end_at"`
	LimitApplication *int    `json:"limit_application"`
	LinkedEmail      *string `json:"linked_email"`
}
