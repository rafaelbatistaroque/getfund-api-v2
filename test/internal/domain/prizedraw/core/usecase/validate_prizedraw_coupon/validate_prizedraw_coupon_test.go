package validate_prizedraw_coupon_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	fixture "getfund-api-v2/test/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/validate_prizedraw_coupon_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenCouponCodeEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyCouponCode())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "CouponCode"))
}

func Test_GivenExecute_WhenCouponCodeInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidCouponCode())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "CouponCode", 8))
}

func Test_GivenExecute_WhenEmailEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyEmail())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Email"))
}

func Test_GivenExecute_WhenEmailInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidEmail())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_EMAIL_INVALID.Error(), "Email"))
}

func Test_GivenExecute_WhenInputValid_ThenEnsureCallGetCouponByCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetCouponByCode:couponCode"]).Be(validInput.CouponCode)
}

func Test_GivenExecute_WhenGetCouponByCodeError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["GetCouponByCode"])
}

func Test_GivenExecute_WhenGetCouponByCodeWithSuccessNull_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon null")
}

func Test_GivenExecute_WhenGetCouponByCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetCouponByCode"]).Be(1)
}

func Test_GivenExecute_WhenCouponFoundValidityNotStartYet_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedCouponFound := fixture.GetCouponNotStartYet(time.Hour * 72)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponFound)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon validity has not start yet")
}

func Test_GivenExecute_WhenCouponTypeIsUniqueApplicationByEmail_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithEmailLinked()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon not applicable to this email")
}

func Test_GivenExecute_WhenCouponTypeIsUniqueApplicationApplied_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithApplication(validInput.UserId, entity.UNIQUE_APPLICATION_TYPE)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon already applied")
}

func Test_GivenExecute_WhenCouponTypeIsByApplicationLimitReached_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithApplicationReached(5, entity.LIMIT_APPLICATION_TYPE)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon application limit reached")
}

func Test_GivenExecute_WhenCouponFoundValidityHasEndAtLessThanNow_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expiredCoupon := fixture.GetExpiredCoupon()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expiredCoupon)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon expired")
}

func Test_GivenExecute_WhenCouponAlreadyAppliedByUser_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithApplication(validInput.UserId, entity.LIMIT_APPLICATION_TYPE)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon already applied by user")
}

func Test_GivenExecute_WhenCouponValidityIsValidWithoutPrizeDrawLinked_ThenEnsureCallGetPrizeDrawByIdWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetCouponWithoutPrizeDrawLinked(0))
	expectedPrizedDrawId := 40

	// Act
	sut.Execute(fixture.GetInput(fixture.WithSelectedPrizeDrawId(expectedPrizedDrawId)))

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetPrizeDrawById:id"]).Be(expectedPrizedDrawId)
}

func Test_GivenExecute_WhenCouponValidityIsValid_ThenEnsureCallGetPrizeDrawByIdWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedCouponResponse := fixture.GetValidCoupon()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponResponse)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetPrizeDrawById:id"]).Be(expectedCouponResponse.PrizeDrawId)
}

func Test_GivenExecute_WhenGetPrizeDrawByIdError_ThenEnsureCallApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message.Error()).Be("prizedraw not found")
}

func Test_GivenExecute_WhenGetPrizeDrawByIdInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetPrizeDrawById"]).Be(1)
}

func Test_GivenExecute_WhenGetPrizeDrawByIdWithSuccessNull_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("invalid prizedraw data")
}

func Test_GivenExecute_WhenGetPrizeDrawByIdWithSuccessHasWinner_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	winner := 1
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(&prizedraw_dto.PrizeDrawDto{WinnerEntranceId: &winner})

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("prizedraw has winner")
}

func Test_GivenExecute_WhenGetPrizeDrawByIdSuccessDifferentFromSelectedPrizeDraw_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(&prizedraw_dto.PrizeDrawDto{Id: 8})

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("prizedraw is not valid for this coupon")
}

func Test_GivenExecute_WhenPrizeDrawIsValid_ThenEnsureCallEmitAndWaitPromiseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	coupon := spies.RepoSpy.SuccessResult["GetCouponByCode"].(*prizedraw_dto.CouponDto)
	expectedPayload := &prizedraw_payload.ValidatePrizeDrawCouponStartedPayload{
		ProductId: coupon.ProductId,
	}

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:event"][0]).Be(&validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:payload"][0]).Be(expectedPayload)
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:result"][0]).NotNil()
}

func Test_GivenExecute_WhenEmitAndWaitPromiseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitAndWaitPromise"]).Be(1)
}

func Test_GivenExecute_WhenEmitAndWaitPromiseError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	expectedError := "[coupon validate] " + spies.BusSpy.SuccessResult["EmitAndWaitPromise"].(*bus.Promise).ErrorToString()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be(expectedError)
}

func Test_GivenExecute_WhenEmitAndWaitPromiseWithNullProduct_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.BusSpy.DefineEmitAndWaitPromiseErrorNull()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("[coupon validate] invalid product data")
}

func Test_GivenExecute_WhenProductInactive_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.BusSpy.DefineResult(fixture.GetInactiveProductResponse())

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert

	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("inactive product")
}

func Test_GivenExecute_WhenCouponProductDifferentFromSelectedProduct_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.BusSpy.DefineResult(fixture.GetProductWithDiferentIdResponse())

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon is not valid for this product")
}

func Test_GivenExecute_WhenValidateCouponSuccess_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validCoupon := fixture.GetValidCoupon()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(validCoupon)
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.BusSpy.DefineResult(fixture.GetProductResponse())

	// Act
	result, _ := sut.Execute(fixture.GetInput())

	// Assert

	verify.Should(t, result.Message).Be("coupon is valid")
	verify.Should(t, result.CouponId).Be(validCoupon.Id)
	verify.Should(t, result.PrizeDrawId).Be(validCoupon.PrizeDrawId)
	verify.Should(t, result.ProductId).Be(validCoupon.ProductId)
}
