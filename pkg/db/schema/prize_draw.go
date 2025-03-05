package schema

type PrizeDraw struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name                string `gorm:"not null;column:name"`
	DetailedDescription string `gorm:"not null;column:detailed_description"`
	PrizeDescription    string `gorm:"not null;column:prize_description"`
	ExpectedAmount      int    `gorm:"not null;column:expected_amount"`
	StartAt             int    `gorm:"not null;index:idx_prize_draw_list;column:start_at"`
	EndAt               *int   `gorm:"column:end_at"`
	PrizeDrawAt         *int   `gorm:"column:prize_draw_at"`
	WinnerEntranceID    *uint  `gorm:"index:idx_prize_draw_winner;column:winner_entrance_id"`
	IsActive            *bool  `gorm:"column:is_active"`
	RetentionRate       *int   `gorm:"column:retention_rate"`
	FreeFunding         *int   `gorm:"column:free_funding"`
	CreatedAt           int64  `gorm:"column:created_at"`
	UpdatedAt           int64  `gorm:"column:updated_at"`
}

func (PrizeDraw) TableName() string {
	return "prize_draw"
}
