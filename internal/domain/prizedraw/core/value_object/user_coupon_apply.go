package value_object

import "time"

type CouponUserApply struct {
	userId           int
	couponId         int
	isNewApplication bool

	createdAt time.Time
	updatedAt time.Time
}

func NewUserCouponApply(userId, couponId int) *CouponUserApply {
	return &CouponUserApply{
		userId:           userId,
		couponId:         couponId,
		isNewApplication: true,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}
}

func FillUserCouponApply(userId, couponId int, createdAt, updatedAt time.Time) *CouponUserApply {
	return &CouponUserApply{
		userId:           userId,
		couponId:         couponId,
		isNewApplication: true,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
	}
}

func (u CouponUserApply) GetUserId() int            { return u.userId }
func (u CouponUserApply) GetCouponId() int          { return u.couponId }
func (u CouponUserApply) GetIsNewApplication() bool { return u.isNewApplication }
func (u CouponUserApply) GetCreatedAt() time.Time   { return u.createdAt }
func (u CouponUserApply) GetUpdatedAt() time.Time   { return u.updatedAt }

func (u CouponUserApply) IsEqual(userId int) bool {
	return u.userId == userId
}
