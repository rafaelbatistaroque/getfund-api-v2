package activate_user_with_coupon_confirmed_event_handler_test

import (
	fixture "getfund-api-v2/test/external/domain/coupon/event_handler/activate_user_with_coupon_confirmed_event_handler/activate_user_with_coupon_confirmed_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidActivateUserWithCouponConfirmedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	activationDataKey := "user_activation_fake_activation_code"
	expectedKey := activationDataKey + "_coupon"

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(activationDataKey))

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedKey)
}

func Test_GivenHandler_WhenGetCacheInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenHandler_WhenGetCacheError_ThenEnsureNeverCallEmitWithPayloadAndResponse(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	//Assert
	verify.Should(t, spies.ValidateCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenUnmarshalError_ThenEnsureNeverCallEmitWithPayloadAndResponse(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccessWithValue("invalid-serialized-json")

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ValidateCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenInvalidCouponData_ThenEnsureReturn(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetCacheDataWithInvalidCoupon())

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ValidateCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenValidateCouponInvoked_ThenEnsureCallsWithCorrectParamenter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetCacheDataWithValidCoupon())
	expectedInput := fixture.GetValidateCouponInput()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ValidateCouponSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenHandler_WhenValidateCouponError_ThenEnsureNeverCallApplyCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetCacheDataWithValidCoupon())
	spies.ValidateCouponSpy.DefineValidateCouponUsecaseError("")

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ApplyCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenValidateCouponSuccess_ThenEnsureCallApplyCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetCacheDataWithValidCoupon())
	spies.ValidateCouponSpy.DefineValidateCouponUsecaseSuccess()
	expectedInputApply := fixture.GetApplyCouponInput(*fixture.GetCouponData().CouponData)

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ApplyCouponSpy.Params["Execute:input"]).Be(expectedInputApply)
}

func Test_GivenHandler_WhenVApplyCouponInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetCacheDataWithValidCoupon())
	spies.ValidateCouponSpy.DefineValidateCouponUsecaseSuccess()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(""))

	// Assert
	verify.Should(t, spies.ApplyCouponSpy.CallsCount["Execute"]).Be(1)
}
