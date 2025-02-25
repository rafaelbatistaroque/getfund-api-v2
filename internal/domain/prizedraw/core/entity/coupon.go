package entity

import "time"

type Coupon struct {
	code              string
	typeApplicability int
	prizeDrawId       int
	productId         int
	startAt           time.Time
	endAt             *time.Time
	discount          int
	limitApplication  *int
}

func CouponFill(code string, typeApplicability, prizeDrawId, productId, discount int, limitApplication *int, startAt time.Time, endAt *time.Time) *Coupon {
	return &Coupon{
		code:              code,
		typeApplicability: typeApplicability,
		prizeDrawId:       prizeDrawId,
		productId:         productId,
		startAt:           startAt,
		endAt:             endAt,
		discount:          discount,
		limitApplication:  limitApplication,
	}
}

func (c *Coupon) NotStartYet() bool {
	return c.startAt.After(time.Now())
}

func (c *Coupon) GetTypeApplicability() int {
	return c.typeApplicability
}

func (c *Coupon) IsExpired() bool {
	return c.endAt != nil && c.endAt.Before(time.Now())
}

func (c *Coupon) ReachedApplicationLimit(applicationsCount int) bool {
	return c.limitApplication != nil && applicationsCount >= *c.limitApplication
}
