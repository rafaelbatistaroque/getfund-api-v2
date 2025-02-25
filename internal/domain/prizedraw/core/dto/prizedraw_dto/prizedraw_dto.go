package prizedraw_dto

type PrizeDrawDto struct {
	Id               int  `json:"id"`
	WinnerEntranceId *int `json:"winner_entrance_id"`
}
