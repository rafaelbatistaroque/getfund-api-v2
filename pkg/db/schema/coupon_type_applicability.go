package schema

type CouponTypeApplicability struct {
	ID               int     `gorm:"primaryKey;autoIncrement;column:id"`
	CouponTypeCode   string  `gorm:"not null;column:coupon_type_code"`
	StartAt          uint64  `gorm:"not null;column:start_at"`
	EndAt            *uint64 `gorm:"column:end_at"`
	LimitApplication *uint   `gorm:"column:limit_application"`
	LinkedEmail      *string `gorm:"column:linked_email"`
}

func (CouponTypeApplicability) TableName() string {
	return "coupon_type_applicability"
}
