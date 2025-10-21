package apply_prizedraw_coupon_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/prizedraw/core/dto"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_error "getfund-api-v2/internal/shared/error"
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
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "CouponId"))
}

func Test_GivenExecute_WhenPrizeDrawIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithPrizeDrawId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "PrizeDrawId"))
}

func Test_GivenExecute_WhenProductIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithProductId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "ProductId"))
}

func Test_GivenExecute_WhenUserIdZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithUserId(0))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO.Error(), "UserId"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallEmitAndWaitPromiseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseError()
	validInput := fixture.GetInput()
	expectedPayload := &event.ApplyPrizeDrawCouponStartedPayload{
		UserId:       validInput.UserId,
		ProductId:    validInput.ProductId,
		PrizeDrawId:  validInput.PrizeDrawId,
		CouponId:     validInput.CouponId,
		ItemQuantity: 1,
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:event"][0]).Be(&event.ApplyPrizeDrawCouponStartedEvent{})
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
	expectedError := "[coupon apply] " + spies.BusSpy.SuccessResult["EmitAndWaitPromise"].(*shared_bus.Promise).ErrorToString()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be(expectedError)
}

func Test_GivenExecute_WhenEmitAndWaitPromiseZero_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefineEmitAndWaitPromiseErrorNull()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("[coupon apply] invalid purchase id")
}

func Test_GivenExecute_WhenPurchaseIdReceived_ThenEnsureCallHasherRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(8)
}

func Test_GivenExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenExecute_WhenGetRandomCodeError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on build lucky number")
}

func Test_GivenExecute_WhenGetRandomCodeSuccessEmprty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on build lucky number")
}

func Test_GivenExecute_WhenGetRandomCodeSuccess_ThenEnsureCallGetCouponByIdWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	expectedParams := fixture.GetInput()

	// Act
	sut.Execute(expectedParams)

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetCouponById:couponId"]).Be(expectedParams.CouponId)
}

func Test_GivenExecute_WhenGetCouponByIdInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetCouponById"]).Be(1)
}

func Test_GivenExecute_WhenGetCouponByIdError_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineGetCouponByIdError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNAVAILABLE_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["GetCouponById"])
}

func Test_GivenExecute_WhenGetCouponByIdSuccess_ThenEnsureCallWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	validInput := fixture.GetInput()
	expectedCouponDto := fixture.GetValidCoupon()
	fixture.ApplyCoupon(expectedCouponDto, validInput.UserId, validInput.CouponId)
	spies.RepoSpy.DefineGetCouponByIdSuccess(expectedCouponDto)
	expectedEntranceParam := spies.GetEntranceDto()

	// Act
	sut.Execute(validInput)

	// Assert
	saveEntranceWithCouponAppliedParamEntrance := spies.RepoSpy.Params["SaveEntranceWithCouponApplied:entrance"].(*dto.EntranceDto)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.LuckyCode).Be(expectedEntranceParam.LuckyCode)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.PrizeDrawId).Be(expectedEntranceParam.PrizeDrawId)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.PurchaseId).Be(expectedEntranceParam.PurchaseId)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.UserId).Be(expectedEntranceParam.UserId)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.IsDonation).Be(expectedEntranceParam.IsDonation)
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.CreatedAt).NotNil()
	verify.Should(t, saveEntranceWithCouponAppliedParamEntrance.UpdatedAt).NotNil()
	saveEntranceWithCouponAppliedParamCoupon := spies.RepoSpy.Params["SaveEntranceWithCouponApplied:coupon"].(*dto.CouponDto)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.UserCouponApplies[0].UserId).Be(validInput.UserId)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.UserCouponApplies[0].CouponId).Be(expectedCouponDto.Id)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.PrizeDrawId).Be(expectedCouponDto.PrizeDrawId)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.ProductId).Be(expectedCouponDto.ProductId)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.Code).Be(expectedCouponDto.Code)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.CouponTypeApplicability.Id).Be(expectedCouponDto.CouponTypeApplicability.Id)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.CouponTypeApplicability.CouponTypeCode).Be(expectedCouponDto.CouponTypeApplicability.CouponTypeCode)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.CouponTypeApplicability.LimitApplication).Be(expectedCouponDto.CouponTypeApplicability.LimitApplication)
	verify.Should(t, saveEntranceWithCouponAppliedParamCoupon.CouponTypeApplicability.LinkedEmail).Be(expectedCouponDto.CouponTypeApplicability.LinkedEmail)
}

func Test_GivenExecute_WhenSaveEntranceWithCouponAppliedInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, 2)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineGetCouponByIdSuccess(fixture.GetValidCoupon())

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["SaveEntranceWithCouponApplied"]).Be(1)
}

func Test_GivenExecute_WhenSaveEntranceWithCouponAppliedError_ThenEnsureCallEmitAndWaitPromiseWithCorrectParameterANdReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedPurchaseId := 7
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, expectedPurchaseId)
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponFailedEvent{}, false)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineGetCouponByIdSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineSaveEntranceWithCouponAppliedError()
	expectedPayload := &event.ApplyPrizeDrawCouponFailedPayload{
		PurchaseId: expectedPurchaseId,
	}

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:event"][1]).Be(&event.ApplyPrizeDrawCouponFailedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:payload"][1]).Be(expectedPayload)
	verify.Should(t, spies.BusSpy.Params["EmitAndWaitPromise:result"][1]).NotNil()
	verify.Should(t, err.Code).Be(shared_error.UNAVAILABLE_CODE)
	verify.Should(t, err.Message.Error()).Be("erro on apply coupon")
}

func Test_GivenExecute_WhenSaveEntranceWithCouponAppliedSuccess_ThenEnsureReturnSuccessOutput(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedPurchaseId := 7
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponStartedEvent{}, expectedPurchaseId)
	spies.BusSpy.DefinePromiseResult(&event.ApplyPrizeDrawCouponFailedEvent{}, false)
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	spies.RepoSpy.DefineGetCouponByIdSuccess(fixture.GetValidCoupon())
	spies.RepoSpy.DefineSaveEntranceWithCouponAppliedSuccess()

	// Act
	result, _ := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, result.Message).Be("coupon successfully applied")
}
