package schema

type Country struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string `gorm:"not null;column:name"`
	Code      string `gorm:"not null;column:code"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (Country) TableName() string {
	return "country"
}
