package prizedraw_repository_fixture

import (
	prizedraw_repository "getfund-api-v2/internal/domain/prizedraw/adapter/repository"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/pkg/db/schema"
	"getfund-api-v2/test/helper/db_fixture"

	"gorm.io/gorm"
)

func NewSUT() (prizedraw_contract.Repository, *gorm.DB) {
	db := db_fixture.NewMemoryDB(
		&schema.Coupon{},
		&schema.CouponTypeApplicability{},
		&schema.UserCouponApply{},
	)

	return prizedraw_repository.New(db), db
}
