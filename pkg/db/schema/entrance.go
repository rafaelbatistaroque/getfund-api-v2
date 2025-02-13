package schema

import "time"

type Entrance struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        string `gorm:"not null;unique"`
	UserID      uint   `gorm:"index"`
	PrizeDrawID uint   `gorm:"index"`
	PurchaseID  uint   `gorm:"index"`
	PaidAmount  int    `gorm:"not null"`
	PaidAt      int    `gorm:"not null"`
	IsDonation  bool   `gorm:"default:false"`

	// Relacionamentos
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Purchase  Purchase  `gorm:"foreignKey:PurchaseID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
