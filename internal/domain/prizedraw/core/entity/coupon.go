package entity

import (
	"time"
)

type Coupon struct {
	id                int
	code              string
	typeApplicability int
	prizeDrawId       int
	productId         int
	startAt           time.Time
	endAt             *time.Time
	limitApplication  *int
}

func CouponFill(id int, code string, typeApplicability, prizeDrawId, productId int, limitApplication *int, startAt time.Time, endAt *time.Time) *Coupon {
	return &Coupon{
		id:                id,
		code:              code,
		typeApplicability: typeApplicability,
		prizeDrawId:       prizeDrawId,
		productId:         productId,
		startAt:           startAt,
		endAt:             endAt,
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

func (c *Coupon) GetId() int {
	return c.id
}

func (c *Coupon) GetProductId() int {
	return c.productId
}

func (c *Coupon) GetPrizeDrawId() int {
	return c.prizeDrawId
}
