package user_criation_with_coupon_code_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/coupon/event_handler/user_criation_with_coupon_code_started_event_handler/user_criation_with_coupon_code_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCouponCodeUsecase(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidUserCriationWithCouponCodeStarted())

	//Assert
	verify.Should(t, spies.UseCaseSpy.CallsCount["Execute"]).Be(0)
}
