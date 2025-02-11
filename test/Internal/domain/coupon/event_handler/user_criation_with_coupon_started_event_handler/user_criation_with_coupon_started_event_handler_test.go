package user_criation_with_coupon_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/coupon/event_handler/user_criation_with_coupon_started_event_handler/user_criation_with_coupon_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidUserCriationWithCouponStarted())

	//Assert
	verify.Should(t, spies.UseCaseSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallValidadeCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetValidateCouponInput()

	// Act
	sut.Handle(fixture.GetValidUserCriationWithCouponStarted())

	//Assert
	verify.Should(t, spies.UseCaseSpy.Params["Execute:input"]).Be(expectedInput)
}
