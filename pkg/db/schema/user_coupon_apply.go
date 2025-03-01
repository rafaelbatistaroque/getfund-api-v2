package schema

type UserCouponApply struct {
	ID        uint  `gorm:"primaryKey;autoIncrement;column:id"`
	CouponID  int   `gorm:"not null;index;column:coupon_id"`
	UserID    int   `gorm:"not null;index;column:user_id"`
	CreatedAt int64 `gorm:"column:created_at"`
}

func (UserCouponApply) TableName() string {
	return "user_coupon_apply"
}
