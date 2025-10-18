package schema

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;column:id"`
	FirstName string `gorm:"not null;column:first_name"`
	LastName  string `gorm:"not null;column:last_name"`
	Username  string `gorm:"not null;unique;column:username"`
	Password  string `gorm:"not null;column:password"`
	IsAdmin   bool   `gorm:"not null;column:is_admin"`
	IsActive  bool   `gorm:"not null;column:is_active"`
	CreatedAt int64  `gorm:"column:created_at;<-:create"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}
