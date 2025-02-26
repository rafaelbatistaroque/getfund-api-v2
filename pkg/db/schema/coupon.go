package schema

type Coupon struct {
	ID                        uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Code                      string `gorm:"not null;unique;index:idx_coupon_code;column:code"`
	PrizeDrawID               uint   `gorm:"index;column:prize_draw_id"`
	ProductID                 uint   `gorm:"index;column:product_id"`
	CouponTypeApplicabilityID uint   `gorm:"index;column:cupon_type_applicability_id"`
	CreatedAt                 int64  `gorm:"column:created_at"`
	UpdatedAt                 int64  `gorm:"column:updated_at"`

	// Relacionamentos
	CouponTypeApplicability CouponTypeApplicability `gorm:"foreignKey:CouponTypeApplicabilityID;constraint:OnDelete:CASCADE"`
	PrizeDraw               PrizeDraw               `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Product                 Product                 `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	UserCouponApply         []UserCouponApply       `gorm:"foreignKey:CouponID"`
}

func (Coupon) TableName() string {
	return "coupon"
}
