package activate_user_with_coupon_confirmed_event_handler_test

import (
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	fixture "getfund-api-v2/test/internal/domain/prizedraw/adapter/event_handler/activate_user_with_coupon_confirmed_event_handler/activate_user_with_coupon_confirmed_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallValidadeCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidActivateUserWithCouponConfirmedEvent())

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetCouponByCode"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetCouponByCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedCouponCode := "fake-coupon-code"

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent(expectedCouponCode, "", 1))

	//Assert
	verify.Should(t, spies.RepoSpy.Params["GetCouponByCode:couponCode"]).Be(expectedCouponCode)
}

func Test_GivenHandler_WhenGetCouponByCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	//Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetCouponByCode"]).Be(1)
}

func Test_GivenHandler_WhenGetCouponByCodeError_ThenEnsureNeverCallEmitWithPayloadAndResponse(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeError()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	//Assert
	verify.Should(t, spies.ValidatePrizeDrawCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenGetCouponByCodeWithNullSuccess_ThenEnsureCallValidateCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	// Assert
	verify.Should(t, spies.ValidatePrizeDrawCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenGetCouponByCodeSuccess_ThenEnsureCallValidateCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	expectedInput := fixture.GetValidateCouponInput()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("fake-coupon-code", "fake-email", 1))

	// Assert
	verify.Should(t, spies.ValidatePrizeDrawCouponSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenHandler_WhenValidateCouponError_ThenEnsureNeverCallApplyCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.ValidatePrizeDrawCouponSpy.DefineValidateCouponUsecaseError("")

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	// Assert
	verify.Should(t, spies.ApplyPrizeDrawCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenValidateCouponOutputNull_ThenEnsureNeverCallApplyCoupon(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.ValidatePrizeDrawCouponSpy.DefineValidateCouponUsecaseSuccessNull()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	// Assert
	verify.Should(t, spies.ApplyPrizeDrawCouponSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenValidateCouponSuccess_ThenEnsureCallApplyCouponWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validCoupon := fixture.GetValidCoupon()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(validCoupon)
	expectedUserId := 3
	spies.ValidatePrizeDrawCouponSpy.DefineValidateCouponUsecaseSuccessWithOutput(&validate_prizedraw_coupon.Output{
		CouponId:    validCoupon.Id,
		PrizeDrawId: validCoupon.PrizeDrawId,
		ProductId:   validCoupon.ProductId,
	})
	expectedInputApply := fixture.GetApplyCouponInput(validCoupon, expectedUserId)

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", expectedUserId))

	// Assert
	verify.Should(t, spies.ApplyPrizeDrawCouponSpy.Params["Execute:input"]).Be(expectedInputApply)
}

func Test_GivenHandler_WhenApplyCouponInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.ValidatePrizeDrawCouponSpy.DefineValidateCouponUsecaseSuccess()

	// Act
	sut.Handle(fixture.GetValidActivateUserWithCouponConfirmedEvent("", "", 1))

	// Assert
	verify.Should(t, spies.ApplyPrizeDrawCouponSpy.CallsCount["Execute"]).Be(1)
}
