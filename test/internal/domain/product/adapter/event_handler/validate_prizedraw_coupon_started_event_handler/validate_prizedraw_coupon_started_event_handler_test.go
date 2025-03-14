package validate_prizedraw_coupon_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/product/adapter/event_handler/validate_prizedraw_coupon_started_event_handler/validate_prizedraw_coupon_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallGetProductById(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidValidatePrizeDrawCouponStartedEvent())

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetProductById"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetProductByIdWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedProductId := 20

	// Act
	sut.Handle(fixture.GetValidValidatePrizeDrawCouponStartedEvent(expectedProductId))

	//Assert
	verify.Should(t, spies.RepoSpy.Params["GetProductById:productId"]).Be(expectedProductId)
}

func Test_GivenHandler_WhenGetProductByIdInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidValidatePrizeDrawCouponStartedEvent(1))

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetProductById"]).Be(1)
}
