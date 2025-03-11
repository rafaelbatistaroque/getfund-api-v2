package entity

import (
	"errors"
)

const (
	_PRIZE_DRAW_HAS_WINNER        = "prizedraw has winner"
	_COUPON_INVALID_FOR_PRIZEDRAW = "prizedraw is not valid for this coupon"
)

type PrizeDraw struct {
	id               int
	winnerEntranceId *int
}

func FillPrizeDraw(id int, winnerEntranceId *int) *PrizeDraw {
	return &PrizeDraw{
		id:               id,
		winnerEntranceId: winnerEntranceId,
	}
}

func (p *PrizeDraw) GetId() int { return p.id }

func (p *PrizeDraw) Validate(selectedPrizeDrawId int) error {
	if p.winnerEntranceId != nil {
		return errors.New(_PRIZE_DRAW_HAS_WINNER)
	}

	if p.id != selectedPrizeDrawId {
		return errors.New(_COUPON_INVALID_FOR_PRIZEDRAW)
	}

	return nil
}
