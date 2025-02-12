package schema

type FreeFundingUserEntrances struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint `gorm:"index"`
	PrizeDrawID uint `gorm:"index"`

	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
}
