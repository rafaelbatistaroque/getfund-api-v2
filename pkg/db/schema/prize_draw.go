package schema

import "time"

type PrizeDraw struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement"`
	Name                string `gorm:"not null"`
	DetailedDescription string `gorm:"not null"`
	PrizeDescription    string `gorm:"not null"`
	ExpectedAmount      int    `gorm:"not null"`
	StartAt             int    `gorm:"not null;index:idx_prize_draw_list"`
	EndAt               *int
	PrizeDrawAt         *int
	WinnerEntranceID    *uint `gorm:"index:idx_prize_draw_winner"`
	IsActive            *bool
	RetentionRate       *int
	FreeFunding         *int

	// Relacionamento com Entrance
	WinnerEntrance *Entrance `gorm:"foreignKey:WinnerEntranceID;constraint:OnDelete:SET NULL"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
