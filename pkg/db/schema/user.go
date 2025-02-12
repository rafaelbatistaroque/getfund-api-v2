package schema

type User struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	UserCategoryID    uint   `gorm:"index"`
	CountryID         uint   `gorm:"index"`
	FirstName         string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	Email             string `gorm:"not null;index"`
	Username          string `gorm:"not null;unique"`
	Password          string `gorm:"not null"`
	Gender            string `gorm:"not null"`
	MainSocialNetwork string `gorm:"not null"`
	RegisteredAt      int    `gorm:"not null"`
	RegisteredUrl     string `gorm:"not null"`
	IsAdmin           bool   `gorm:"not null"`
	IsActive          bool   `gorm:"not null"`

	// Relacionamentos
	UserCategory UserCategory `gorm:"foreignKey:UserCategoryID;constraint:OnDelete:CASCADE"`
	Country      Country      `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}
