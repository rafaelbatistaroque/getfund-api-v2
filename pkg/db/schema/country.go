package schema

import "time"

type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
	Code string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
