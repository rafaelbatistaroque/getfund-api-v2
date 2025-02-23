package schema

type Purchase struct {
	ID           uint    `gorm:"primaryKey;autoIncrement;column:id"`
	CodePurchase string  `gorm:"not null;column:code_purchase"`
	CodeSuccess  string  `gorm:"not null;column:code_success"`
	CodeFailure  string  `gorm:"not null;column:code_failure"`
	CodeFree     *string `gorm:"column:code_free"`
	PrizeDrawID  uint    `gorm:"index;column:prize_draw_id"`
	ProductID    uint    `gorm:"index;column:product_id"`
	UserID       uint    `gorm:"index;column:user_id"`
	CouponID     *uint   `gorm:"column:coupon_id"`
	OpenedAt     int     `gorm:"not null;index:idx_purchase_opened_at;column:opened_at"`
	ClosedAt     *int    `gorm:"column:closed_at"`
	Successful   *bool   `gorm:"column:successful"`
	CreatedAt    int64   `gorm:"column:created_at"`
	UpdatedAt    int64   `gorm:"column:updated_at"`

	// Relacionamentos
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Coupon    *Coupon   `gorm:"foreignKey:CouponID;constraint:OnDelete:SET NULL"`
}

func (Purchase) TableName() string {
	return "purchase"
}
