package schema

type Transaction struct {
	ID         string `gorm:"primaryKey;column:id"`
	PurchaseID string `gorm:"column:purchase_id"`
	CodeResult string `gorm:"column:code_result"` //random code_success, code_fail, code_free
	CodeType   string `gorm:"column:code_type"`   //success, fail, free, coupon
	Status     string `gorm:"column:status"`      //pending, finished
	Message    string `gorm:"column:message"`
	ClosedAt   *int   `gorm:"column:closed_at"`
	CreatedAt  int64  `gorm:"column:created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at"`

	Purchase Purchase `gorm:"foreignKey:PurchaseID"`
}
