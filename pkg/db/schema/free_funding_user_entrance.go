package schema

type FreeFundingUserEntrances struct {
	ID          uint  `gorm:"primaryKey;autoIncrement;column:id"`
	UserID      uint  `gorm:"index;column:user_id"`
	PrizeDrawID uint  `gorm:"index;column:prize_draw_id"`
	CreatedAt   int64 `gorm:"column:created_at;<-:create"`
	UpdatedAt   int64 `gorm:"column:updated_at"`

	// Relacionamentos
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
}

func (FreeFundingUserEntrances) TableName() string {
	return "free_funding_user_entrance"
}
