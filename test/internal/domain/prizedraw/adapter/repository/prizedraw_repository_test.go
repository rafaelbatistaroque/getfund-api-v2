package prizedraw_repository_test

import (
	fixture "getfund-api-v2/test/internal/domain/prizedraw/adapter/repository/prizedraw_repository_fixture"
	"testing"

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
