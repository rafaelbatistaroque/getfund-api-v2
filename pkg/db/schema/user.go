package schema

// Modelo de usuário
type User struct {
	ID                uint   `gorm:"primaryKey;autoIncrement;column:id"`
	UserCategoryID    uint   `gorm:"index;column:user_category_id"`
	CountryID         uint   `gorm:"index;column:country_id"`
	FirstName         string `gorm:"not null;column:first_name"`
	LastName          string `gorm:"not null;column:last_name"`
	Email             string `gorm:"not null;index;column:email"`
	Username          string `gorm:"not null;unique;column:username"`
	Password          string `gorm:"not null;column:password"`
	Gender            string `gorm:"not null;column:gender"`
	MainSocialNetwork string `gorm:"not null;column:main_social_network"`
	RegisteredUrl     string `gorm:"not null;column:registered_url"`
	IsAdmin           bool   `gorm:"not null;column:is_admin"`
	IsActive          bool   `gorm:"not null;column:is_active"`
	CreatedAt         int64  `gorm:"column:created_at"`
	UpdatedAt         int64  `gorm:"column:updated_at"`

	// Relacionamentos
	UserCategory UserCategory `gorm:"foreignKey:UserCategoryID;constraint:OnDelete:CASCADE"`
	Country      Country      `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "user"
}
