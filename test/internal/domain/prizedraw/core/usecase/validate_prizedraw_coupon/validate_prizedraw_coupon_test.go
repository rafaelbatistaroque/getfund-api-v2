package validate_prizedraw_coupon_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
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

func Test_GivenExecute_WhenInputValid_ThenEnsureCallGetCouponByCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetCouponByCode:couponCode"]).Be(validInput.CouponCode)
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

func Test_GivenExecute_WhenCouponFoundValidityNotStartYet_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedCouponFound := &prizedraw_dto.CouponDto{StartAt: int64(time.Hour * 72)}
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponFound)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon validity has not start yet")
}

func Test_GivenExecute_WhenCouponTypeIsUniqueApplicationApplied_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithApplication(validInput.UserId, 1)
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
	couponAppliedByUser := fixture.GetValidCouponWithApplicationReached(5, 2)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon already applied")
}

func Test_GivenExecute_WhenCouponAlreadyAppliedByUser_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	couponAppliedByUser := fixture.GetValidCouponWithApplication(validInput.UserId, 0)
	spies.RepoSpy.DefineGetCouponByCodeSuccess(couponAppliedByUser)

	// Act
	_, err := sut.Execute(validInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon already applied by user")
}

func Test_GivenExecute_WhenCouponFoundValidityHasEndAtLessThanNow_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	minus72Hours := time.Now().Add(-72 * time.Hour).Unix()
	minus24Hours := time.Now().Add(-24 * time.Hour).Unix()
	expectedCouponFound := &prizedraw_dto.CouponDto{
		StartAt:           minus72Hours,
		EndAt:             &minus24Hours,
		TypeApplicability: 3,
	}
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponFound)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon expired")
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
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(&prizedraw_dto.PrizeDrawDto{Id: 5})

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("prizedraw is not valid for this coupon")
}

func Test_GivenExecute_WhenPrizeDrawIsValid_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	coupon := spies.RepoSpy.SuccessResult["GetCouponByCode"].(*prizedraw_dto.CouponDto)
	expectedPayload := &prizedraw_payload.ValidateCouponStartedPayload{
		ProductId: coupon.ProductId,
	}

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	_, isChannelType := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	verify.Should(t, isChannelType).BeTrue()
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:event"][0]).Be(&validate_prizedraw_coupon.ValidateCouponStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:payload"][0]).Be(expectedPayload)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayloadAndResponse"]).Be(1)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseTimeout_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("timeout waiting for coupon validation")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		[]byte(""),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("empty response from coupon validation")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvalid_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		[]byte("invalid-value"),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("invalid response from coupon validation")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseWithNullProduct_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetNullResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("invalid product data")
}

func Test_GivenExecute_WhenProductInactive_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetInactiveProductResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("inactive product")
}

func Test_GivenExecute_WhenCouponProductDifferentFromSelectedProduct_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput(fixture.WithSelectedProductId(2)))
			errResult <- err
		},
		fixture.GetProductResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("coupon is not valid for this product")
}

func Test_GivenExecute_WhenValidateCouponSuccess_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineGetPrizeDrawByIdSuccess(fixture.GetValidPrizeDraw())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	resultChannel := make(chan *validate_prizedraw_coupon.Output, 1)

	// Act
	spies.BusSpy.RunSutWithEventResponse(
		func() {
			result, _ := sut.Execute(fixture.GetInput())
			resultChannel <- result
		},
		fixture.GetProductResponse(),
	)

	// Assert
	resultUnwrapped := <-resultChannel
	defer close(resultChannel)
	verify.Should(t, resultUnwrapped.Message).Be("coupon is valid")
}
