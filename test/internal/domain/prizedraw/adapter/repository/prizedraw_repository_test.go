package prizedraw_repository_test

import (
	"getfund-api-v2/internal/infra/db/schema"
	fixture "getfund-api-v2/test/internal/domain/prizedraw/adapter/repository/prizedraw_repository_fixture"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetCouponByCode_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB.DB()
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
	currentDb, _ := db.DB.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetPrizeDrawById(0)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetPrizeDrawById_WhenNotFound_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetPrizeDrawById(0)

	// Assert
	verify.Should(t, err.Error()).Be("prize draw not found")
}

func Test_GivenGetPrizeDrawById_WhenFound_ThenEnsureReturnSuccess(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	fakeWinner := 450
	prizeDraw := &schema.PrizeDraw{
		Name:                "fake-name",
		DetailedDescription: "fake-description",
		PrizeDescription:    "fake-prizedraw-description",
		ExpectedAmount:      200,
		StartAt:             int(time.Now().Unix()),
		WinnerEntranceID:    &fakeWinner,
	}
	db.Create(prizeDraw)

	// Act
	prizeDrawFound, err := sut.GetPrizeDrawById(int(prizeDraw.ID))

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, prizeDrawFound).NotNil()
	verify.Should(t, prizeDrawFound.Id).Be(int(prizeDraw.ID))
	verify.Should(t, prizeDrawFound.WinnerEntranceId).Be(prizeDraw.WinnerEntranceID)
}

func Test_GivenGetCouponById_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetCouponById(0)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetCouponById_WhenNotFound_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetCouponById(40)

	// Assert
	verify.Should(t, err.Error()).Be("coupon not found")
}

func Test_GivenGetCouponById_WhenFound_ThenEnsureReturnSuccess(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	couponType := &schema.CouponTypeApplicability{}
	db.Create(couponType)
	coupon := &schema.Coupon{CouponTypeApplicabilityID: couponType.ID}
	db.Create(coupon)
	userCouponApply := &schema.UserCouponApply{CouponID: int(coupon.ID)}
	db.Create(userCouponApply)

	// Act
	couponFound, err := sut.GetCouponById(1)

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, couponFound).NotNil()
	verify.Should(t, couponFound.Id).Be(1)
	verify.Should(t, couponFound.CouponTypeApplicability.Id).Be(couponType.ID)
	verify.Should(t, len(couponFound.UserCouponApplies)).Be(1)
	verify.Should(t, couponFound.UserCouponApplies[0].CouponId).Be(int(coupon.ID))
}

func Test_GivenSaveEntranceWithCouponApplied_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB.DB()
	currentDb.Close()

	// Act
	err := sut.SaveEntranceWithCouponApplied(nil, nil)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenSaveEntranceWithCouponApplied_WhenEntranceError_ThenEnsureReturnErrorAndRoolback(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	invalidEntrance := fixture.GetInvalidEntranceDto()

	// Act
	err := sut.SaveEntranceWithCouponApplied(invalidEntrance, nil)

	// Assert
	verify.Should(t, err).NotNil()
	var entrance = &schema.Entrance{}
	notFoundError := db.Select("code").Where("id = ?", 1).First(entrance).Error
	verify.Should(t, notFoundError.Error()).Be("record not found")
}

func Test_GivenSaveEntranceWithCouponApplied_WhenApplyCouponError_ThenEnsureReturnErrorAndRoolback(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	validEntrance := fixture.GetEntranceDto()
	invalidCoupon := fixture.GetCoupon()

	// Act
	err := sut.SaveEntranceWithCouponApplied(validEntrance, invalidCoupon)

	// Assert
	verify.Should(t, err).NotNil()
	var entrance = &schema.Entrance{}
	notFoundEntrance := db.Select("code").Where("id = ?", 1).First(entrance).Error
	verify.Should(t, notFoundEntrance.Error()).Be("record not found")
	var coupon = &schema.Coupon{}
	notFoundCoupon := db.Select("code").Where("id = ?", 1).First(coupon).Error
	verify.Should(t, notFoundCoupon.Error()).Be("record not found")
}
