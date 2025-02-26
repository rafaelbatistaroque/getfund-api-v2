package entity

import (
	vo "getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

const (
	UNIQUE_APPLICATION_BY_EMAIL_TYPE = "unique_by_email"
	UNIQUE_APPLICATION_TYPE          = "unique_application"
	LIMIT_APPLICATION_TYPE           = "limited_application"
	EXPIRATION_TYPE                  = "expiration_time"
)

type Coupon struct {
	id          int
	code        string
	prizeDrawId int
	productId   int

	couponApplicability *vo.CouponTypeApplicability
	couponUserApplies   []vo.CouponUserApply
}

func CouponFill(id int, code string, prizeDrawId, productId int, userCouponApplies []vo.CouponUserApply, couponApplicability *vo.CouponTypeApplicability) *Coupon {
	return &Coupon{
		id:                  id,
		code:                code,
		prizeDrawId:         prizeDrawId,
		productId:           productId,
		couponUserApplies:   userCouponApplies,
		couponApplicability: couponApplicability,
	}
}

func (c *Coupon) GetId() int                { return c.id }
func (c *Coupon) GetProductId() int         { return c.productId }
func (c *Coupon) GetPrizeDrawId() int       { return c.prizeDrawId }
func (c *Coupon) GetCouponTypeCode() string { return c.couponApplicability.GetCouponTypeCode() }
func (c *Coupon) CountApplies() int         { return len(c.couponUserApplies) }

func (c *Coupon) NotStartYet() bool {
	return c.couponApplicability.GetStartAt().After(time.Now())
}

func (c *Coupon) IsExpired() bool {
	return c.couponApplicability.GetEndAt() != nil && c.couponApplicability.GetEndAt().Before(time.Now())
}

func (c *Coupon) ReachedApplicationLimit() bool {
	limit := c.couponApplicability.GetLimitApplication()
	return limit != nil && len(c.couponUserApplies) >= *limit
}

func (c *Coupon) IsNotSameLinkedEmail(email string) bool {
	return *c.couponApplicability.GetLinkedEmail() != email
}

func (c *Coupon) CouponAlreadyAppliedByUser(userId int) bool {
	for _, user := range c.couponUserApplies {
		if user.IsEqual(userId) {
			return true
		}
	}
	return false
}
