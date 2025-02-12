package schema

type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
	Code string `gorm:"not null"`
}
