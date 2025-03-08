package apply_prizedraw_coupon_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	fixture "getfund-api-v2/test/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/apply_prizedraw_coupon_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenCouponIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithCouponId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "CouponId"))
}

func Test_GivenExecute_WhenPrizeDrawIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithPrizeDrawId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "PrizeDrawId"))
}

func Test_GivenExecute_WhenProductIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithProductId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "ProductId"))
}

func Test_GivenExecute_WhenUserIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithUserId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "UserId"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallEmitAndWaitPromiseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	validInput := fixture.GetInput()
	expectedPayload := &prizedraw_payload.ApplyPrizeDrawCouponStartedPayload{
		UserId:       validInput.UserId,
		ProductId:    validInput.ProductId,
		PrizeDrawId:  validInput.PrizeDrawId,
		CouponId:     validInput.CouponId,
		ItemQuantity: 1,
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:event"][0]).Be(&apply_prizedraw_coupon.ApplyPrizeDrawCouponStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:payload"][0]).Be(expectedPayload)
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:result"][0]).NotNil()
}

func Test_GivenExecute_WhenEmitAndWaitPromiseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitAndWaitPromise"]).Be(1)
}

func Test_GivenExecute_WhenEmitAndWaitPromiseEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	expectedError := "[coupon apply] " + spies.BusSpy.SuccessResult["EmitAndWaitPromise"].(*bus.Promise).ErrorToString()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be(expectedError)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseZero_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseErrorNull()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("[coupon apply] invalid purchase id")
}

func Test_GivenExecute_WhenPurchaseIdReceived_ThenEnsureCallHasherRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(8)
}

func Test_GivenExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenExecute_WhenGetRandomCodeError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on build lucky number")
}
func Test_GivenExecute_WhenGetRandomCodeSuccessEmprty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on build lucky number")
}

func Test_GivenExecute_WhenGetRandomCodeSuccess_ThenEnsureCallCreateEntranceWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	validEntranceDto := spies.GetEntranceDto()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	entranceParams := spies.RepoSpy.Params["CreateEntrance:entrance"].(*prizedraw_dto.EntranceDto)
	verify.Should(t, entranceParams.LuckyCode).Be(validEntranceDto.LuckyCode)
	verify.Should(t, entranceParams.UserId).Be(validEntranceDto.UserId)
	verify.Should(t, entranceParams.PrizeDrawId).Be(validEntranceDto.PrizeDrawId)
	verify.Should(t, entranceParams.PurchaseId).Be(validEntranceDto.PurchaseId)
	verify.Should(t, entranceParams.IsDonation).Be(validEntranceDto.IsDonation)
}

func Test_GivenExecute_WhenCreateEntranceInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["CreateEntrance"]).Be(1)
}

func Test_GivenExecute_WhenCreateEntranceError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineResult(2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineCreateEntranceError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on create entrance")
}
