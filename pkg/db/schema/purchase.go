package schema

type Purchase struct {
	ID              uint   `gorm:"primaryKey;autoIncrement;column:id"`
	UserID          uint   `gorm:"index;column:user_id"`
	ProductID       uint   `gorm:"index;column:product_id"`
	Code            string `gorm:"not null;column:code"`
	ItemType        string `gorm:"column:item_type"` //prizedraw, campaign
	ItemID          uint   `gorm:"index;column:item_id"`
	ItemQuantity    uint   `gorm:"column:item_quantity"`
	ProcessCode     string `gorm:"column:code"`      //random code_payment (success or fail), code_free, couponId
	ProcessCodeType string `gorm:"column:code_type"` //payment, free, coupon
	Status          string `gorm:"column:status"`    //pending, success, fail
	Message         string `gorm:"column:message"`
	CreatedAt       int64  `gorm:"column:created_at;<-:create"`
	UpdatedAt       int64  `gorm:"column:updated_at"`

	// Relacionamentos
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Purchase) TableName() string {
	return "purchase"
}
