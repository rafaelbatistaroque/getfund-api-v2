package prizedraw_dto

import (
	"getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"time"
)

type UserCouponApplyDto struct {
	UserId           int   `json:"user_id"`
	CouponId         int   `json:"coupon_id"`
	IsNewApplication bool  `json:"is_new_application"`
	CreatedAt        int64 `json:"created_at"`
	UpdatedAt        int64 `json:"updated_at"`
}

func (v UserCouponApplyDto) ToEntity() *value_object.CouponUserApply {
	return value_object.FillUserCouponApply(
		v.UserId,
		v.CouponId,
		time.Unix(v.CreatedAt, 0),
		time.Unix(v.UpdatedAt, 0),
	)
}

func ToUserCouponApplyDto(entity value_object.CouponUserApply) *UserCouponApplyDto {
	return &UserCouponApplyDto{
		UserId:           entity.GetUserId(),
		CouponId:         entity.GetUserId(),
		IsNewApplication: entity.GetIsNewApplication(),
		CreatedAt:        entity.GetCreatedAt().Unix(),
		UpdatedAt:        entity.GetUpdatedAt().Unix(),
	}
}
