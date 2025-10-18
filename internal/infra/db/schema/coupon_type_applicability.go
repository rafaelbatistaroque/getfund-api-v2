package schema

type CouponTypeApplicability struct {
	ID               int     `gorm:"primaryKey;autoIncrement;column:id"`
	CouponTypeCode   string  `gorm:"not null;column:coupon_type_code"`
	StartAt          int64   `gorm:"not null;column:start_at"`
	EndAt            *int64  `gorm:"column:end_at"`
	LimitApplication *int    `gorm:"column:limit_application"`
	LinkedEmail      *string `gorm:"column:linked_email"`
}

func (CouponTypeApplicability) TableName() string {
	return "coupon_type_applicability"
}
