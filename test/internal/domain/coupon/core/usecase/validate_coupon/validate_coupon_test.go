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
