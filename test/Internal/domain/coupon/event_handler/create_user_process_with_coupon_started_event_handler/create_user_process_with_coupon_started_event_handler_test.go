package create_user_process_with_coupon_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/coupon/event_handler/create_user_process_with_coupon_started_event_handler/create_user_process_with_coupon_started_event_handler_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.UseCaseSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallValidadeCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetValidateCouponInput()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.UseCaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenHandler_WhenValidadeCouponInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.UseCaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenHandler_WhenValidadeCouponErrorWithStatus_ThenEnsureCallCacheSetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedMessage := "any-message-status"
	spies.UseCaseSpy.DefineValidateCouponUsecaseError("status:" + expectedMessage)
	expectedCacheValue := fixture.GetCreateUserProcessWithCouponPayload(expectedMessage)
	expectedCacheKey := expectedCacheValue.ActivationDataKey + "_coupon"

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedCacheKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedCacheValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenHandler_WhenValidadeCouponSuccess_ThenEnsureCallCacheSetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.UseCaseSpy.DefineValidateCouponUsecaseSuccess()
	expectedCacheValue := fixture.GetCreateUserProcessWithCouponPayload("")
	expectedCacheKey := expectedCacheValue.ActivationDataKey + "_coupon"

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessWithCouponStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedCacheKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedCacheValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}
