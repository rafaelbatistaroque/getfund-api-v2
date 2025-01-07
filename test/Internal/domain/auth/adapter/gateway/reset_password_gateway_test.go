package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/reset_password_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenResetPassword_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("with-body-error")

	// Act
	_, code, err := sut.ResetPassword(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenResetPassword_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, resetPasswordUsecaseSpy := fixture.NewSut()
	expectedInput := fixture.GetResetPasswordInput()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.ResetPassword(res, req)

	// Assert
	verify.Should(t, resetPasswordUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenResetPassword_WhenDecodeSuccess_ThenEnsureCallUsecaseOnce(t *testing.T) {
	// Arrange
	sut, resetPasswordUsecaseSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.ResetPassword(res, req)

	// Assert
	verify.Should(t, resetPasswordUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenResetPassword_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, resetPasswordUsecaseSpy := fixture.NewSut()
	resetPasswordUsecaseSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.ResetPassword(res, req)

	// Assert
	verify.Should(t, code).Be(resetPasswordUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(resetPasswordUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenResetPassword_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, resetPasswordUsecaseSpy := fixture.NewSut()
	resetPasswordUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	result, code, _ := sut.ResetPassword(res, req)

	// Assert
	success := result.(*reset_password.Output)
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Message).Be(resetPasswordUsecaseSpy.SuccessResult["Execute"].Message)
}
