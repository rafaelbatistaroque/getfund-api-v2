package schema

type UserCategory struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string `gorm:"not null;column:name"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (UserCategory) TableName() string {
	return "user_category"
}
