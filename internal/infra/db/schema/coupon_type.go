package schema

type CouponType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Code        uint   `gorm:"not null;unique;column:code"`
	Description string `gorm:"not null;column:description"`
	CreatedAt   int64  `gorm:"column:created_at;<-:create"`
	UpdatedAt   int64  `gorm:"column:updated_at"`
}

func (CouponType) TableName() string {
	return "coupon_type"
}
