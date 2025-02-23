package schema

type Coupon struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Code         string `gorm:"not null;unique;index:idx_coupon_code;column:code"`
	PrizeDrawID  uint   `gorm:"index;column:prize_draw_id"`
	ProductID    uint   `gorm:"index;column:product_id"`
	StartAt      int    `gorm:"not null;column:start_at"`
	EndAt        *int   `gorm:"column:end_at"`
	Discount     int    `gorm:"not null;column:discount"`
	CouponTypeID uint   `gorm:"not null;index;column:coupon_type_id"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`

	// Relacionamentos
	CouponType CouponType   `gorm:"foreignKey:CouponTypeID;constraint:OnDelete:CASCADE"`
	PrizeDraw  PrizeDraw    `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product    Product      `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	UserCoupon []UserCoupon `gorm:"foreignKey:CouponID"`
}

func (Coupon) TableName() string {
	return "coupon"
}
