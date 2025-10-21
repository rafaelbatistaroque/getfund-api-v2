package repository

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"

	"gorm.io/gorm"
)

type prizedrawRepository struct {
	db *db.GetFund
}

func New(db *db.GetFund) prizedraw_contract.Repository {
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

	userApply := make([]*prizedraw_dto.UserCouponApplyDto, len(coupon.UserCouponApply))
	for i, apply := range coupon.UserCouponApply {
		userApply[i] = &prizedraw_dto.UserCouponApplyDto{
			UserId:   apply.UserID,
			CouponId: apply.CouponID,
		}
	}

	return p.convertCouponToDto(coupon), nil
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

	return &prizedraw_dto.PrizeDrawDto{
		Id:               int(prizeDraw.ID),
		WinnerEntranceId: prizeDraw.WinnerEntranceID,
	}, nil
}

func (p *prizedrawRepository) CreateEntrance(entrance *prizedraw_dto.EntranceDto) error {
	return nil
}

func (p *prizedrawRepository) SaveEntranceWithCouponApplied(entrance *prizedraw_dto.EntranceDto, coupon *prizedraw_dto.CouponDto) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if len(entrance.LuckyCode) > 8 {
			return gorm.ErrInvalidValue
		}

		if err := tx.Create(&schema.Entrance{
			Code:        entrance.LuckyCode,
			UserID:      uint(entrance.UserId),
			PrizeDrawID: uint(entrance.PrizeDrawId),
			PurchaseID:  uint(entrance.PurchaseId),
			IsDonation:  entrance.IsDonation,
			CreatedAt:   entrance.CreatedAt,
			UpdatedAt:   entrance.UpdatedAt,
		}).Error; err != nil {
			return err
		}

		if areAllAppliesFalse(coupon) {
			return errors.New("coupon must have at least one application")
		}

		for _, apply := range coupon.UserCouponApplies {
			if !apply.IsNewApplication {
				continue
			}

			if err := tx.Create(&schema.UserCouponApply{
				UserID:    apply.UserId,
				CouponID:  apply.CouponId,
				CreatedAt: apply.CreatedAt,
				UpdatedAt: apply.UpdatedAt,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *prizedrawRepository) GetCouponById(couponId int) (*prizedraw_dto.CouponDto, error) {
	var coupon = &schema.Coupon{}

	result := p.db.
		Preload("CouponTypeApplicability").
		Preload("UserCouponApply").
		Select("id, code, product_id, prize_draw_id, coupon_type_applicability_id").
		Where("id=?", couponId).
		First(coupon)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("coupon not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return p.convertCouponToDto(coupon), nil
}

func (p *prizedrawRepository) convertCouponToDto(coupon *schema.Coupon) *prizedraw_dto.CouponDto {
	userApply := make([]*prizedraw_dto.UserCouponApplyDto, len(coupon.UserCouponApply))
	for i, apply := range coupon.UserCouponApply {
		userApply[i] = &prizedraw_dto.UserCouponApplyDto{
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
	}
}

func areAllAppliesFalse(coupon *prizedraw_dto.CouponDto) bool {
	allFalse := true

	for _, apply := range coupon.UserCouponApplies {
		if apply.IsNewApplication {
			allFalse = false
			break
		}
	}
	return allFalse
}
