package schema

import "time"

type Coupon struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Code         string `gorm:"not null;unique;index:idx_coupon_code"`
	PrizeDrawID  uint   `gorm:"index"`
	ProductID    uint   `gorm:"index"`
	StartAt      int    `gorm:"not null"`
	EndAt        *int
	Discount     int          `gorm:"not null"`
	CouponTypeID uint         `gorm:"not null;index"`
	UserCoupon   []UserCoupon `gorm:"foreignKey:CouponID"`

	// Relacionamentos
	CouponType CouponType `gorm:"foreignKey:CouponTypeID;constraint:OnDelete:CASCADE"`
	PrizeDraw  PrizeDraw  `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product    Product    `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CouponType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        uint   `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserCoupon struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CouponID  uint `gorm:"not null;index"`
	UserID    uint `gorm:"not null;index"`
	CreatedAt time.Time
}
