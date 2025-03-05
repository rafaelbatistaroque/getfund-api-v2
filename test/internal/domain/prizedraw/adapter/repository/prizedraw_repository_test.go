package prizedraw_repository_test

import (
	"getfund-api-v2/pkg/db/schema"
	fixture "getfund-api-v2/test/internal/domain/prizedraw/adapter/repository/prizedraw_repository_fixture"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetCouponByCode_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetCouponByCode("invalid-coupon-code")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetCouponByCode_WhenNotFound_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetCouponByCode("non-existent-username")

	// Assert
	verify.Should(t, err.Error()).Be("coupon not found")
}

func Test_GivenGetCouponByCode_WhenFound_ThenEnsureReturnSuccess(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	couponCode := uuid.NewString()
	couponType := &schema.CouponTypeApplicability{}
	db.Create(couponType)
	coupon := &schema.Coupon{Code: couponCode, CouponTypeApplicabilityID: couponType.ID}
	db.Create(coupon)
	userCouponApply := &schema.UserCouponApply{CouponID: int(coupon.ID)}
	db.Create(userCouponApply)

	// Act
	couponFound, err := sut.GetCouponByCode(couponCode)

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, couponFound).NotNil()
	verify.Should(t, couponFound.Id).Be(1)
	verify.Should(t, couponFound.Code).Be(couponCode)
	verify.Should(t, couponFound.CouponTypeApplicability.Id).Be(couponType.ID)
	verify.Should(t, len(couponFound.UserCouponApplies)).Be(1)
	verify.Should(t, couponFound.UserCouponApplies[0].CouponId).Be(int(coupon.ID))
}

func Test_GivenGetPrizeDrawById_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetPrizeDrawById(0)

	// Assert
	verify.Should(t, err).NotNil()
}
