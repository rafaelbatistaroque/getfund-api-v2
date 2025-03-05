package schema

type Purchase struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;column:id"`
	UserID       uint   `gorm:"index;column:user_id"`
	ProductID    uint   `gorm:"index;column:product_id"`
	Code         string `gorm:"not null;column:code"`
	ItemType     string `gorm:"index;column:item_type"` //prizedraw, campaign, coupon
	ItemID       uint   `gorm:"index;column:item_id"`
	ItemQuantity uint   `gorm:"column:item_quantinty"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`

	// Relacionamentos
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Purchase) TableName() string {
	return "purchase"
}
