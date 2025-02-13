package schema

import "time"

type Coupon struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        string `gorm:"not null;unique;index:idx_coupon_code"`
	PrizeDrawID uint   `gorm:"index"`
	ProductID   uint   `gorm:"index"`
	StartAt     int    `gorm:"not null"`
	EndAt       *int
	Discount    int `gorm:"not null"`

	// Relacionamentos
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
