package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/recover_password_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenRecoverPassword_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	res, req := spies.GetHttpRequestResponse("with-body-error")

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenRecoverPassword_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetRecoverPasswordInput()
	res, req := spies.GetHttpRequestResponse(fixture.GetRecoverPasswordInputSerialized())

	// Act
	sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, spies.RecoverPasswordUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenRecoverPassword_WhenDecodeSuccess_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	res, req := spies.GetHttpRequestResponse(fixture.GetRecoverPasswordInputSerialized())

	// Act
	sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, spies.RecoverPasswordUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenRecoverPassword_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RecoverPasswordUsecaseSpy.DefineError()
	res, req := spies.GetHttpRequestResponse(fixture.GetRecoverPasswordInputSerialized())

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(spies.RecoverPasswordUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.RecoverPasswordUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenRecoverPassword_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RecoverPasswordUsecaseSpy.DefineSuccess()
	res, req := spies.GetHttpRequestResponse(fixture.GetRecoverPasswordInputSerialized())

	// Act
	result, code, _ := sut.RecoverPassword(res, req)

	// Assert
	success := result.(*recover_password.Output)
	verify.Should(t, code).Be(shared_error.SUCCESS_CODE)
	verify.Should(t, success.Message).Be(spies.RecoverPasswordUsecaseSpy.SuccessResult["Execute"].Message)
}
