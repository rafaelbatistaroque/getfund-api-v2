package schema

type Profile struct {
	ID                uint   `gorm:"primaryKey;autoIncrement;column:id"`
	SocialName        string `gorm:"not null;column:social_name"`
	UseSocialName     bool   `gorm:"not null;column:use_social_name"`
	AlternativeEmail  string `gorm:"not null;index;column:alternative_email"`
	Gender            string `gorm:"not null;column:gender"`
	MainSocialNetwork string `gorm:"not null;column:main_social_network"`
	RegisteredUrl     string `gorm:"not null;column:registered_url"`
	CreatedAt         int64  `gorm:"column:created_at;<-:create"`
	UpdatedAt         int64  `gorm:"column:updated_at"`

	UserCategoryID uint `gorm:"index;column:user_category_id"`
	CountryID      uint `gorm:"index;column:country_id"`

	// Relacionamentos
	UserCategory UserCategory `gorm:"foreignKey:UserCategoryID;constraint:OnDelete:CASCADE"`
	Country      Country      `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
}

func (Profile) TableName() string {
	return "profile"
}
