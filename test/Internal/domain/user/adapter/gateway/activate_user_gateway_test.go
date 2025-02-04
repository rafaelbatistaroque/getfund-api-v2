package user_gateway_test

import (
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/adapter/gateway/activate_user_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenActivateUser_WhenActivationCodeUrlParamNotFound_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Error()).Be("activation code is required")
}

func Test_GivenActivateUser_WhenActivationCodeUrlParamFound_ThenEnsureCallUseCAseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	expectedInput := fixture.GetActivateUserInput()
	res, req := fixture.GetHttpRequestResponse(expectedInput.ActivationCode)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, usecase.ActivateUserUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenActivateUser_WhenAUsecaseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("fake-activation-code")

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, usecase.ActivateUserUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenActivateUser_WhenUsecaseError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(usecase.ActivateUserUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(usecase.ActivateUserUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenActivateUser_WhenUsecaseSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	output, code, _ := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, output).Be(usecase.ActivateUserUsecaseSpy.SuccessResult["Execute"])
}
