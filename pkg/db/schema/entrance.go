package schema

type Entrance struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Code        string `gorm:"not null;unique;column:code"`
	UserID      uint   `gorm:"index;column:user_id"`
	PrizeDrawID uint   `gorm:"index;column:prize_draw_id"`
	PurchaseID  uint   `gorm:"index;column:purchase_id"`
	PaidAmount  int    `gorm:"not null;column:paid_amount"`
	PaidAt      int    `gorm:"not null;column:paid_at"`
	IsDonation  bool   `gorm:"default:false;column:is_donation"`
	CreatedAt   int64  `gorm:"column:created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at"`

	// Relacionamentos
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	// Purchase Purchase `gorm:"foreignKey:PurchaseID;constraint:OnDelete:CASCADE"`
}

func (Entrance) TableName() string {
	return "entrance"
}
