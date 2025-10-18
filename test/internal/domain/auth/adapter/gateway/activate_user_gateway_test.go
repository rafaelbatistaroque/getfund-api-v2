package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/activate_user_gateway_fixture"
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
	verify.Should(t, code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Error()).Be("activation code is required")
}

func Test_GivenActivateUser_WhenActivationCodeUrlParamFound_ThenEnsureCallUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	expectedInput := fixture.GetActivateUserInput()
	res, req := fixture.GetHttpRequestResponse(expectedInput.ActivationCode)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, usecase.ActivateUserUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenActivateUser_WhenUsecaseInvoked_ThenEnsureCallsOnce(t *testing.T) {
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

func Test_GivenActivateUser_WhenActivateUserUsecaseSuccess_ThenEnsureCallSigninUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineSuccessWithValue()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	sut.ActivateUser(res, req)

	// Assert
	activateUserOutput := usecase.ActivateUserUsecaseSpy.SuccessResult["Execute"]
	verify.Should(t, usecase.SigninUsecaseSpy.Params["Execute:input"]).Be(&signin.Input{
		Username: activateUserOutput.Username,
		Password: activateUserOutput.Password,
	})
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccessNUll_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Error()).Be("error to activate user")
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccess_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, usecase.SigninUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenActivateUser_WhenSigninUsecaseError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineSuccess()
	usecase.SigninUsecaseSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(usecase.SigninUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(usecase.SigninUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccess_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, usecase := fixture.NewSut()
	usecase.ActivateUserUsecaseSpy.DefineSuccess()
	usecase.SigninUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("valid-param")

	// Act
	result, code, _ := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SUCCESS_CODE)
	verify.Should(t, result).Be(usecase.SigninUsecaseSpy.SuccessResult["Execute"])
}
