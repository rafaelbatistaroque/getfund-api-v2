package schema

type Product struct {
	ID            uint    `gorm:"primaryKey;autoIncrement;column:id"`
	StripePriceID *string `gorm:"column:stripe_price_id"`
	Name          string  `gorm:"not null;column:name"`
	EntranceQty   int     `gorm:"not null;column:entrance_qty"`
	TotalPrice    int     `gorm:"not null;column:total_price"`
	IsActive      bool    `gorm:"not null;index:idx_product_active;column:is_active"`
	CreatedAt     int64   `gorm:"column:created_at"`
	UpdatedAt     int64   `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "product"
}
