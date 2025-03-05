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
		Preload("CouponTypeApplicability").
		Preload("UserCouponApply").
		Select("id, code, product_id, prize_draw_id, coupon_type_applicability_id").
		Where("code=?", couponCode).
		First(coupon)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("coupon not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	userApply := make([]prizedraw_dto.UserCouponApplyDto, len(coupon.UserCouponApply))
	for i, apply := range coupon.UserCouponApply {
		userApply[i] = prizedraw_dto.UserCouponApplyDto{
			UserId:   apply.UserID,
			CouponId: apply.CouponID,
		}
	}

	return &prizedraw_dto.CouponDto{
		Id:                int(coupon.ID),
		Code:              coupon.Code,
		PrizeDrawId:       coupon.PrizeDrawID,
		ProductId:         coupon.ProductID,
		UserCouponApplies: userApply,
		CouponTypeApplicability: &prizedraw_dto.CouponTypeApplicabilityDto{
			Id:               coupon.CouponTypeApplicability.ID,
			CouponTypeCode:   coupon.CouponTypeApplicability.CouponTypeCode,
			LimitApplication: coupon.CouponTypeApplicability.LimitApplication,
			LinkedEmail:      coupon.CouponTypeApplicability.LinkedEmail,
			StartAt:          coupon.CouponTypeApplicability.StartAt,
			EndAt:            coupon.CouponTypeApplicability.EndAt,
		},
	}, nil
}

func (p *prizedrawRepository) GetPrizeDrawById(id int) (*prizedraw_dto.PrizeDrawDto, error) {
	var prizeDraw = &schema.PrizeDraw{}

	result := p.db.
		Select("id, winner_entrance_id").
		Where("id=?", id).
		First(prizeDraw)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("prize draw not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}
