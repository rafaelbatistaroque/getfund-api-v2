package value_object

import (
	"time"
)

type CouponTypeApplicability struct {
	id               int
	linkedEmail      *string
	couponTypeCode   string
	startAt          time.Time
	endAt            *time.Time
	limitApplication *int
}

func NewCouponTypeApplicability(id int, couponTypeCode string, linkedEmail *string, startAt time.Time, endAt *time.Time, limitApplication *int) *CouponTypeApplicability {
	return &CouponTypeApplicability{
		id:               id,
		linkedEmail:      linkedEmail,
		couponTypeCode:   couponTypeCode,
		limitApplication: limitApplication,
		startAt:          startAt,
		endAt:            endAt,
	}
}

func (c CouponTypeApplicability) GetId() int                { return c.id }
func (c CouponTypeApplicability) GetCouponTypeCode() string { return c.couponTypeCode }
func (c CouponTypeApplicability) GetLinkedEmail() *string   { return c.linkedEmail }
func (c CouponTypeApplicability) GetStartAt() time.Time     { return c.startAt }
func (c CouponTypeApplicability) GetEndAt() *time.Time      { return c.endAt }
func (c CouponTypeApplicability) GetLimitApplication() *int { return c.limitApplication }
