package authadapter

import (
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/pkg/verify"
	"getfund-api-v2/internal/shared/resultapp"
	fixture "getfund-api-v2/test/domain/auth/port/adapter/signoutfixture"
	"net/http"
	"testing"
)

func Test_GivenSignout_WhenDecodeError_ThenEnsureReturnStatusBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("body-with-error")

	// Act
	_, code, err := sut.Signout(res, req)

	// Assert
	verify.Should(t, code).Be(http.StatusBadRequest)
	verify.Should(t, err).NotNil()
}

func Test_GivenSignout_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInput := fixture.GetSignoutInput()
	sut, signoutSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signout(res, req)

	// Assert
	verify.Should(t, signoutSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignout_WhenDecodeSuccess_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, signoutSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signout(res, req)

	// Assert
	verify.Should(t, signoutSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenSignout_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, signoutSpy := fixture.NewSut()
	signoutSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.Signout(res, req)

	// Assert
	verify.Should(t, code).Be(signoutSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(signoutSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignout_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, signoutSpy := fixture.NewSut()
	signoutSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	result, code, _ := sut.Signout(res, req)

	// Assert
	success := result.(*signout.Output)
	verify.Should(t, code).Be(resultapp.CODE_SUCCESS)
	verify.Should(t, success.Message).Be(signoutSpy.SuccessResult["Execute"].Message)
}
