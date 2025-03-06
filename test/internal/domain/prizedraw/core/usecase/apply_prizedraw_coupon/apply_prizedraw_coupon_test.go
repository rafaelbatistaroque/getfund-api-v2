package apply_prizedraw_coupon_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
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

func Test_GivenExecute_WhenValidInput_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
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
	_, isChannelType := spies.BusSpy.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	verify.Should(t, isChannelType).BeTrue()
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:event"][0]).Be(&apply_prizedraw_coupon.ApplyPrizeDrawCouponStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayloadAndResponse:payload"][0]).Be(expectedPayload)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayloadAndResponse"]).Be(1)
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseTimeout_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("timeout waiting for coupon apply")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, []byte(""))

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("empty response from coupon apply")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseInvalidType_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, []byte("invalid-type"))

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("invalid response from coupon apply")
}

func Test_GivenExecute_WhenEmitWithPayloadAndResponseZero_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, []byte("0"))

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.UNAVAILABLE_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("invalid purchase id")
}

func Test_GivenExecute_WhenPurchaseIdReceived_ThenEnsureCallHasherRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	done := make(chan bool, 1)

	// Act
	spies.BusSpy.Run(func() {
		sut.Execute(fixture.GetInput())
		done <- true
	}, fixture.WithValidPurchaseId())

	// Assert
	<-done
	defer close(done)
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(8)
}

func Test_GivenExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	spies.HasherSpy.DefineGetRandomCodeError()
	done := make(chan bool, 1)

	// Act
	spies.BusSpy.Run(func() {
		sut.Execute(fixture.GetInput())
		done <- true
	}, fixture.WithValidPurchaseId())

	// Assert
	<-done
	defer close(done)
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenExecute_WhenGetRandomCodeError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	spies.HasherSpy.DefineGetRandomCodeError()
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, fixture.WithValidPurchaseId())

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("erro on build lucky number")
}
func Test_GivenExecute_WhenGetRandomCodeSuccessEmprty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, fixture.WithValidPurchaseId())

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("erro on build lucky number")
}

func Test_GivenExecute_WhenGetRandomCodeSuccess_ThenEnsureCallCreateEntranceWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	validEntranceDto := spies.GetEntranceDto()
	done := make(chan bool, 1)

	// Act
	spies.BusSpy.Run(func() {
		sut.Execute(fixture.GetInput())
		done <- true
	}, fixture.WithValidPurchaseId())

	// Assert
	<-done
	defer close(done)
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
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	done := make(chan bool, 1)

	// Act
	spies.BusSpy.Run(func() {
		sut.Execute(fixture.GetInput())
		done <- true
	}, fixture.WithValidPurchaseId())

	// Assert
	<-done
	defer close(done)
	verify.Should(t, spies.RepoSpy.CallsCount["CreateEntrance"]).Be(1)
}

func Test_GivenExecute_WhenCreateEntranceError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SettingsSpy.SetTimeoutResponseEvent(1)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineCreateEntranceError()
	errResult := make(chan *result_app.ApplicationError, 1)

	// Act
	spies.BusSpy.Run(func() {
		_, err := sut.Execute(fixture.GetInput())
		errResult <- err
	}, fixture.WithValidPurchaseId())

	// Assert
	errUnwrapped := <-errResult
	defer close(errResult)
	verify.Should(t, errUnwrapped.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, errUnwrapped.Message.Error()).Be("erro on create entrance")
}
