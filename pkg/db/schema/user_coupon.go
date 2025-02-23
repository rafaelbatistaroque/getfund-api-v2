package schema

type UserCoupon struct {
	ID        uint  `gorm:"primaryKey;autoIncrement;column:id"`
	CouponID  uint  `gorm:"not null;index;column:coupon_id"`
	UserID    uint  `gorm:"not null;index;column:user_id"`
	CreatedAt int64 `gorm:"column:created_at"`
}

func (UserCoupon) TableName() string {
	return "user_coupons"
}
