package entity

import "time"

type Entrance struct {
	id          int
	luckyCode   string
	userId      int
	prizeDrawId int
	purchaseId  int
	isDonation  bool

	createdAt time.Time
	updatedAt time.Time
}

func NewEntrance(luckyCode string, userId int, prizeDrawId int, purchaseId int, isDonation bool) *Entrance {
	now := time.Now()
	return &Entrance{
		luckyCode:   getValidLuckCode(luckyCode),
		userId:      getValidId(userId),
		prizeDrawId: getValidId(prizeDrawId),
		purchaseId:  getValidId(purchaseId),
		isDonation:  isDonation,
		createdAt:   now,
		updatedAt:   now,
	}
}

func (e *Entrance) GetId() int                   { return e.id }
func (e *Entrance) GetLuckyCode() string         { return e.luckyCode }
func (e *Entrance) GetUserId() int               { return e.userId }
func (e *Entrance) GetPrizeDrawId() int          { return e.prizeDrawId }
func (e *Entrance) GetPurchaseId() int           { return e.purchaseId }
func (e *Entrance) GetIsDonation() bool          { return e.isDonation }
func (e *Entrance) GetCreatedAt() time.Time      { return e.createdAt }
func (e *Entrance) GetUpdatedAt() time.Time      { return e.updatedAt }
func (e *Entrance) SetUpdatedAt(value time.Time) { e.updatedAt = value }

func getValidLuckCode(value string) string {
	if len(value) != 8 {
		panic("error on create entrance entity")
	}

	return value
}

func getValidId(value int) int {
	if value <= 0 {
		panic("error on create entrance entity")
	}

	return value
}
