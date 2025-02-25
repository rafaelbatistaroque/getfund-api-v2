package entity

type PrizeDraw struct {
	id               int
	winnerEntranceId *int
}

func PrizeDrawFill(id int, winnerEntranceId *int) *PrizeDraw {
	return &PrizeDraw{
		id:               id,
		winnerEntranceId: winnerEntranceId,
	}
}
func (p *PrizeDraw) GetId() int {
	return p.id
}
func (p *PrizeDraw) HasWinner() bool {
	return p.winnerEntranceId != nil
}
