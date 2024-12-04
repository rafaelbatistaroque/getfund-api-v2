package authadapter

import (
	"getfund-api-v2/internal/pkg/verify"
	"getfund-api-v2/internal/shared/resultapp"
	fixture "getfund-api-v2/test/domain/auth/port/adapter/recoverpasswordfixture"
	"testing"
)

func Test_GivenRecoverPassword_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("with-body-error")

	// Act
	_, code, err := sut.RecoverPassword(res, req)

	// Assert
	verify.Should(t, code).Be(resultapp.BAD_REQUEST)
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
