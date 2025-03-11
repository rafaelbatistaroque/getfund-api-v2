package prizedraw_dto

import (
	"getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

type CouponTypeApplicabilityDto struct {
	Id               int     `json:"id"`
	CouponTypeCode   string  `json:"coupon_type_code"`
	StartAt          int64   `json:"start_at"`
	EndAt            *int64  `json:"end_at"`
	LimitApplication *int    `json:"limit_application"`
	LinkedEmail      *string `json:"linked_email"`
}

func (c CouponTypeApplicabilityDto) ToEntity() *value_object.CouponTypeApplicability {
	var endAt *time.Time
	if c.EndAt != nil {
		endAtTime := time.Unix(*c.EndAt, 0)
		endAt = &endAtTime
	}

	return value_object.NewCouponTypeApplicability(
		c.Id,
		c.CouponTypeCode,
		c.LinkedEmail,
		time.Unix(c.StartAt, 0),
		endAt,
		c.LimitApplication,
	)
}

func ToCouponTypeApplicabilityDto(entity value_object.CouponTypeApplicability) *CouponTypeApplicabilityDto {
	var endAt *int64
	if entity.GetEndAt() != nil {
		endAtUnix := entity.GetEndAt().Unix()
		endAt = &endAtUnix
	}

	return &CouponTypeApplicabilityDto{
		Id:               entity.GetId(),
		CouponTypeCode:   entity.GetCouponTypeCode(),
		StartAt:          entity.GetStartAt().Unix(),
		EndAt:            endAt,
		LimitApplication: entity.GetLimitApplication(),
		LinkedEmail:      entity.GetLinkedEmail(),
	}
}
