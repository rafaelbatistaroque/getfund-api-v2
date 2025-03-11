package schema

type Entrance struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Code        string `gorm:"size:8;not null;unique;column:code;<-:create"`
	UserID      uint   `gorm:"index;column:user_id"`
	PrizeDrawID uint   `gorm:"index;column:prize_draw_id"`
	PurchaseID  uint   `gorm:"index;column:purchase_id"`
	IsDonation  bool   `gorm:"default:false;column:is_donation"`
	CreatedAt   int64  `gorm:"column:created_at;<-:create"`
	UpdatedAt   int64  `gorm:"column:updated_at;"`

	// Relacionamentos
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PrizeDraw PrizeDraw `gorm:"foreignKey:PrizeDrawID;constraint:OnDelete:CASCADE"`
	Purchase  Purchase  `gorm:"foreignKey:PurchaseID;constraint:OnDelete:CASCADE"`
}

func (Entrance) TableName() string {
	return "entrance"
}
