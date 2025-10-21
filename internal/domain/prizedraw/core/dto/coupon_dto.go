package dto

import (
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

type CouponDto struct {
	Id                      int                         `json:"id"`
	Code                    string                      `json:"code"`
	PrizeDrawId             int                         `json:"prize_draw_id"`
	ProductId               int                         `json:"product_id"`
	CouponTypeApplicability *CouponTypeApplicabilityDto `json:"coupon_type_applicability"`
	UserCouponApplies       []*UserCouponApplyDto       `json:"user_coupon_aplies"`
	CreatedAt               int64                       `json:"created_at"`
	UpdatedAt               int64                       `json:"updated_at"`
}

func (c CouponDto) ToEntity() *entity.Coupon {
	userCouponApplies := make([]value_object.CouponUserApply, len(c.UserCouponApplies))

	for i, userCounpon := range c.UserCouponApplies {
		userCouponApplies[i] = *userCounpon.ToEntity()
	}

	return entity.FillCoupon(
		c.Id,
		c.Code,
		c.PrizeDrawId,
		c.ProductId,
		userCouponApplies,
		c.CouponTypeApplicability.ToEntity(),
		time.Unix(c.CreatedAt, 0),
		time.Unix(c.UpdatedAt, 0),
	)
}

func ToCouponDto(entity *entity.Coupon) (dto *CouponDto) {
	dto = new(CouponDto)

	userApplyDto := make([]*UserCouponApplyDto, len(entity.GetUserCouponApplies()))
	for i, userApply := range entity.GetUserCouponApplies() {
		userApplyDto[i] = ToUserCouponApplyDto(userApply)
	}

	dto.Id = entity.GetId()
	dto.Code = entity.GetCode()
	dto.PrizeDrawId = entity.GetPrizeDrawId()
	dto.ProductId = entity.GetProductId()
	dto.CouponTypeApplicability = ToCouponTypeApplicabilityDto(*entity.GetCouponTypeApplicability())
	dto.UserCouponApplies = userApplyDto
	dto.CreatedAt = entity.GetCreatedAt().Unix()
	dto.UpdatedAt = entity.GetUpdatedAt().Unix()

	return dto
}
