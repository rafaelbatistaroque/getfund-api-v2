package schema

type Product struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	StripePriceID *string
	Name          string `gorm:"not null"`
	EntranceQty   int    `gorm:"not null"`
	TotalPrice    int    `gorm:"not null"`
	IsActive      bool   `gorm:"not null;index:idx_product_active"`
}
