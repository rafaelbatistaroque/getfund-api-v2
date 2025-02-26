package entity

import (
	vo "getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

const (
	UNIQUE_APPLICATION_BY_EMAIL_TYPE = 1
	UNIQUE_APPLICATION_TYPE          = 2
	LIMIT_APPLICATION_TYPE           = 3
	EXPIRATION_TYPE                  = 4
)

type Coupon struct {
	id                int
	code              string
	couponType        vo.CouponType
	prizeDrawId       int
	productId         int
	startAt           time.Time
	endAt             *time.Time
	limitApplication  *int
	linkedEmail       string
	userCouponApplies []vo.UserCouponApply
}

func CouponFill(id int, code, linkedEmail string, userCouponApplies []vo.UserCouponApply, couponType vo.CouponType, prizeDrawId, productId int, limitApplication *int, startAt time.Time, endAt *time.Time) *Coupon {
	return &Coupon{
		id:                id,
		code:              code,
		couponType:        couponType,
		prizeDrawId:       prizeDrawId,
		productId:         productId,
		startAt:           startAt,
		endAt:             endAt,
		limitApplication:  limitApplication,
		linkedEmail:       linkedEmail,
		userCouponApplies: userCouponApplies,
	}
}

func (c *Coupon) NotStartYet() bool {
	return c.startAt.After(time.Now())
}

func (c *Coupon) GetCouponType() int {
	return int(c.couponType.GetCode())
}

func (c *Coupon) IsExpired() bool {
	return c.endAt != nil && c.endAt.Before(time.Now())
}

func (c *Coupon) ReachedApplicationLimit() bool {
	return c.limitApplication != nil && c.CountApplies() >= *c.limitApplication
}

func (c *Coupon) CountApplies() int {
	return len(c.userCouponApplies)
}

func (c *Coupon) CouponAlreadyAppliedByUser(userId int) bool {
	for _, application := range c.userCouponApplies {
		if application.IsEqual(userId) {
			return true
		}
	}
	return false
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

func (c *Coupon) IsNotSameLinkedEmail(email string) bool {
	return c.linkedEmail != email
}
