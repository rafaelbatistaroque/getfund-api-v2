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
