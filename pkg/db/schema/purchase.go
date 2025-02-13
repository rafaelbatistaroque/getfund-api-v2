package schema

import "time"

type Purchase struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	CodePurchase string `gorm:"not null"`
	CodeSuccess  string `gorm:"not null"`
	CodeFailure  string `gorm:"not null"`
	CodeFree     *string
	PrizeDrawID  uint `gorm:"index"`
	ProductID    uint `gorm:"index"`
	UserID       uint `gorm:"index"`
	CouponID     *uint
	OpenedAt     int `gorm:"not null;index:idx_purchase_opened_at"`
	ClosedAt     *int
	Successful   *bool

	// Relacionamentos
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Coupon    *Coupon   `gorm:"foreignKey:CouponID;constraint:OnDelete:SET NULL"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
