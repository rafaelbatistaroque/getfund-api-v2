package validate_coupon_test

import (
	"fmt"
	coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"
	"getfund-api-v2/internal/domain/coupon/core/dto/coupon_payload"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/coupon/core/usecase/validate_coupon/validate_coupon_fixture"
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
	verify.Should(t, err.Message.Error()).Be("error on get coupon data [found null]")
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
	expectedCouponFound := &coupon_dto.CouponDto{StartAt: int64(time.Hour * 72)}
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponFound)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon validity has not start yet")
}

func Test_GivenExecute_WhenCouponFoundValidityHasEndAtLessThanNow_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	minus72Hours := time.Now().Add(-72 * time.Hour).Unix()
	minus24Hours := time.Now().Add(-24 * time.Hour).Unix()
	expectedCouponFound := &coupon_dto.CouponDto{
		StartAt: minus72Hours,
		EndAt:   &minus24Hours,
	}
	spies.RepoSpy.DefineGetCouponByCodeSuccess(expectedCouponFound)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("coupon expired")
}

func Test_GivenExecute_WhenCouponValidityIsValid_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	coupon := spies.RepoSpy.SuccessResult["GetCouponByCode"].(*coupon_dto.CouponDto)
	expectedPayload := &coupon_payload.ValidateCouponStartedPayload{
		ProductId:   coupon.ProductId,
		PrizeDrawId: coupon.PrizeDrawId,
	}

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	_, isChannelType := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:event"][0]).Be(&validate_coupon.ValidateCouponStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:payload"][0]).Be(expectedPayload)
	verify.Should(t, isChannelType).BeTrue()
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayloadAndResponse"]).Be(1)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseTimeout_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("error on get coupon data [timeout]")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
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
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [empty response]")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvalid_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
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
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [invalid response]")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseWithNullProduct_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetPrizeDrawResponse(),
		fixture.GetNullResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [invalid product]")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseWithNullPrizeDraw_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetProductResponse(),
		fixture.GetNullResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [invalid prizedraw]")
}

func Test_GivenExecute_WhenProductInactive_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetInactiveProductResponse(),
		fixture.GetPrizeDrawResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [inactive product]")
}

func Test_GivenExecute_WhenPrizeDrawhasWinner_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetCouponByCodeSuccess(fixture.GetValidCoupon())
	spies.SettingsSpy.SetTimeoutResponseEvent(2)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.RunSutWithEventResponse(
		func() {
			_, err := sut.Execute(fixture.GetInput())
			errResult <- err
		},
		fixture.GetProductResponse(),
		fixture.GetPrizeDrawWithWinnerResponse(),
	)

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("error on get coupon data [prizedraw has winner]")
}
