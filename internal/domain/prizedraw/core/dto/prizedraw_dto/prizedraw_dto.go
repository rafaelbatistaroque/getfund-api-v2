package prizedraw_dto

import "getfund-api-v2/internal/domain/prizedraw/core/entity"

type PrizeDrawDto struct {
	Id               int  `json:"id"`
	WinnerEntranceId *int `json:"winner_entrance_id"`
}

func (p *PrizeDrawDto) ToEntity() *entity.PrizeDraw {
	return entity.FillPrizeDraw(p.Id, p.WinnerEntranceId)
}
