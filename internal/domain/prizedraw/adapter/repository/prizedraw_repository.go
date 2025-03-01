package prizedraw_repository

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/pkg/db/schema"

	"gorm.io/gorm"
)

type prizedrawRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) prizedraw_contract.Repository {
	return &prizedrawRepository{db: db}
}

func (p *prizedrawRepository) GetCouponByCode(couponCode string) (*prizedraw_dto.CouponDto, error) {
	var coupon = &schema.Coupon{}

	result := p.db.
		Select("id").
		Where("code=?", couponCode).
		First(coupon)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("coupon not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}

func (p *prizedrawRepository) GetPrizeDrawById(id int) (*prizedraw_dto.PrizeDrawDto, error) {
	return nil, nil
}
