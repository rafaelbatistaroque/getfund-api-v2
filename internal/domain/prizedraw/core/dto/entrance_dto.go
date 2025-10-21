package dto

import "getfund-api-v2/internal/domain/prizedraw/core/entity"

type EntranceDto struct {
	Id          int    `json:"id"`
	LuckyCode   string `json:"lucky_number"`
	UserId      int    `json:"user_id"`
	PrizeDrawId int    `json:"prize_draw_id"`
	PurchaseId  int    `json:"purchase_id"`
	IsDonation  bool   `json:"is_donation"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func ToEntranceDto(entity *entity.Entrance) (dto *EntranceDto) {
	dto = new(EntranceDto)

	dto.Id = entity.GetId()
	dto.LuckyCode = entity.GetLuckyCode()
	dto.UserId = entity.GetUserId()
	dto.PrizeDrawId = entity.GetPrizeDrawId()
	dto.PurchaseId = entity.GetPurchaseId()
	dto.IsDonation = entity.GetIsDonation()
	dto.CreatedAt = entity.GetCreatedAt().Unix()
	dto.UpdatedAt = entity.GetUpdatedAt().Unix()

	return dto
}
