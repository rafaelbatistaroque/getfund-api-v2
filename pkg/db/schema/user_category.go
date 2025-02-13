package schema

import "time"

type UserCategory struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
