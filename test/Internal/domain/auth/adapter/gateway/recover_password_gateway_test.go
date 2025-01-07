package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/recover_password_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenRecoverPassword_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("with-body-error")

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenRecoverPassword_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, recoverPasswordSpy := fixture.NewSut()
	expectedInput := fixture.GetRecoverPasswordInput()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, recoverPasswordSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenRecoverPassword_WhenDecodeSuccess_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, recoverPasswordSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, recoverPasswordSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenRecoverPassword_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, recoverPasswordSpy := fixture.NewSut()
	recoverPasswordSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(recoverPasswordSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(recoverPasswordSpy.ErrorResult["Execute"].Message)
}

func Test_GivenRecoverPassword_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, recoverPasswordSpy := fixture.NewSut()
	recoverPasswordSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	result, code, _ := sut.RecoverPassword(res, req)

	// Assert
	success := result.(*recover_password.Output)
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Message).Be(recoverPasswordSpy.SuccessResult["Execute"].Message)
}
