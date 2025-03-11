package entity

import (
	"errors"
	vo "getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

const (
	UNIQUE_APPLICATION_BY_EMAIL_TYPE = "unique_by_email"
	UNIQUE_APPLICATION_TYPE          = "unique_application"
	LIMIT_APPLICATION_TYPE           = "limited_application"
	EXPIRATION_TYPE                  = "expiration_time"

	_EXPIRED_COUPON              = "coupon expired"
	_HAS_NOT_START               = "coupon validity has not start yet"
	_COUPON_APPLIED_BY_USER      = "coupon already applied by user"
	_COUPON_NOT_APPLICABLE_EMAIL = "coupon not applicable to this email"
	_COUPON_ALREADY_APPLIED      = "coupon already applied"
	_COUPON_LIMIT_REACHED        = "coupon application limit reached"
	_FOUND_NULL                  = "coupon null"
)

type Coupon struct {
	id          int
	code        string
	prizeDrawId int
	productId   int

	couponApplicability *vo.CouponTypeApplicability
	couponUserApplies   []vo.CouponUserApply

	createdAt time.Time
	updatedAt time.Time
}

func FillCoupon(id int, code string, prizeDrawId, productId int, userCouponApplies []vo.CouponUserApply, couponApplicability *vo.CouponTypeApplicability, createdAt, updatedAt time.Time) *Coupon {
	return &Coupon{
		id:                  id,
		code:                code,
		prizeDrawId:         prizeDrawId,
		productId:           productId,
		couponUserApplies:   userCouponApplies,
		couponApplicability: couponApplicability,
		createdAt:           createdAt,
		updatedAt:           updatedAt,
	}
}

func (c *Coupon) GetId() int                                 { return c.id }
func (c *Coupon) GetCode() string                            { return c.code }
func (c *Coupon) GetProductId() int                          { return c.productId }
func (c *Coupon) GetPrizeDrawId() int                        { return c.prizeDrawId }
func (c *Coupon) GetUserCouponApplies() []vo.CouponUserApply { return c.couponUserApplies }
func (c *Coupon) GetCouponTypeCode() string                  { return c.couponApplicability.GetCouponTypeCode() }
func (c *Coupon) GetCreatedAt() time.Time                    { return c.createdAt }
func (c *Coupon) GetUpdatedAt() time.Time                    { return c.updatedAt }
func (c *Coupon) GetCouponTypeApplicability() *vo.CouponTypeApplicability {
	return c.couponApplicability
}

func (c *Coupon) LinkPrizeDrawIfThereIsNo(prizeDrawId int) {
	if c.prizeDrawId != 0 {
		return
	}

	c.prizeDrawId = prizeDrawId
	c.updatedAt = time.Now()
}

func (c *Coupon) ApplyCoupon(userId int) {
	c.couponUserApplies = append(c.couponUserApplies, *vo.NewUserCouponApply(userId, c.GetId()))
	c.updatedAt = time.Now()
}

func (c *Coupon) Validate(email string, userId int) error {
	if c.notStartYet() {
		return errors.New(_HAS_NOT_START)
	}

	switch c.couponApplicability.GetCouponTypeCode() {
	case UNIQUE_APPLICATION_BY_EMAIL_TYPE:
		if c.isNotSameLinkedEmail(email) {
			return errors.New(_COUPON_NOT_APPLICABLE_EMAIL)
		}
	case UNIQUE_APPLICATION_TYPE:
		if len(c.couponUserApplies) >= 1 {
			return errors.New(_COUPON_ALREADY_APPLIED)
		}
	case LIMIT_APPLICATION_TYPE:
		if c.reachedApplicationLimit() {
			return errors.New(_COUPON_LIMIT_REACHED)
		}
	case EXPIRATION_TYPE:
		if c.isExpired() {
			return errors.New(_EXPIRED_COUPON)
		}
	}

	//validate if user already applied coupon
	if c.couponAlreadyAppliedByUser(userId) {
		return errors.New(_COUPON_APPLIED_BY_USER)
	}

	return nil
}

func (c *Coupon) notStartYet() bool {
	return c.couponApplicability.GetStartAt().After(time.Now())
}

func (c *Coupon) isNotSameLinkedEmail(email string) bool {
	return *c.couponApplicability.GetLinkedEmail() != email
}

func (c *Coupon) reachedApplicationLimit() bool {
	limit := c.couponApplicability.GetLimitApplication()
	return limit != nil && len(c.couponUserApplies) >= *limit
}

func (c *Coupon) isExpired() bool {
	return c.couponApplicability.GetEndAt() != nil && c.couponApplicability.GetEndAt().Before(time.Now())
}

func (c *Coupon) couponAlreadyAppliedByUser(userId int) bool {
	for _, user := range c.couponUserApplies {
		if user.IsEqual(userId) {
			return true
		}
	}
	return false
}
