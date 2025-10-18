package prizedraw_repository_fixture

import (
	prizedraw_repository "getfund-api-v2/internal/domain/prizedraw/adapter/repository"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/infra/db"
	"getfund-api-v2/internal/infra/db/schema"
	"getfund-api-v2/test/helper/db_fixture"
	"time"
)

func NewSUT() (prizedraw_contract.Repository, *db.GetFund) {
	get_fund_db := db_fixture.NewMemoryDB(
		&schema.Coupon{},
		&schema.CouponTypeApplicability{},
		&schema.UserCouponApply{},
		&schema.Entrance{},
	)

	return prizedraw_repository.New(get_fund_db), get_fund_db
}

func GetEntranceDto() *prizedraw_dto.EntranceDto {
	return &prizedraw_dto.EntranceDto{
		LuckyCode:   "TS8LS31O",
		UserId:      1,
		PrizeDrawId: 1,
		PurchaseId:  40,
		IsDonation:  false,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

func GetInvalidEntranceDto() *prizedraw_dto.EntranceDto {
	return &prizedraw_dto.EntranceDto{
		LuckyCode:   "TS8LS31OA1AHHAQPFT2439d6",
		UserId:      1,
		PrizeDrawId: 1,
		PurchaseId:  9999999999999,
		IsDonation:  false,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

func GetCoupon() *prizedraw_dto.CouponDto {
	less72Hours := time.Now().Add(-24 * time.Hour).Unix()
	more24Hours := time.Now().Add(24 * time.Hour).Unix()
	email := "fake@mail.com"

	couponDto := &prizedraw_dto.CouponDto{
		ProductId:   10,
		PrizeDrawId: 5,
		Id:          1,
		CouponTypeApplicability: &prizedraw_dto.CouponTypeApplicabilityDto{
			StartAt:     less72Hours,
			EndAt:       &more24Hours,
			LinkedEmail: &email,
		},
	}

	return couponDto
}

func ApplyCoupon(couponDto *prizedraw_dto.CouponDto) {
	couponDto.UserCouponApplies = make([]*prizedraw_dto.UserCouponApplyDto, 1)
	couponDto.UserCouponApplies = append(couponDto.UserCouponApplies, &prizedraw_dto.UserCouponApplyDto{CouponId: 1, UserId: 1, IsNewApplication: true})
	couponDto.UserCouponApplies = append(couponDto.UserCouponApplies, &prizedraw_dto.UserCouponApplyDto{CouponId: 2, UserId: 2})
}
