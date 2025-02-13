package schema

import "time"

type FreeFundingUserEntrances struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint `gorm:"index"`
	PrizeDrawID uint `gorm:"index"`

	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
