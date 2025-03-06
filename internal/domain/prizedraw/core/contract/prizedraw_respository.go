package prizedraw_contract

import "getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"

type Repository interface {
	GetCouponByCode(couponCode string) (*prizedraw_dto.CouponDto, error)
	GetPrizeDrawById(id int) (*prizedraw_dto.PrizeDrawDto, error)
	CreateEntrance(entrance *prizedraw_dto.EntranceDto) (*prizedraw_dto.EntranceDto, error)
}
