package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/external/domain/auth/adapter/gateway/signout_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignout_WhenSessionNotFound_ThenEnsureReturnServerErrorWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("not-found")

	// Act
	_, code, err := sut.Signout(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenSignout_WhenTokenFound_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetSignoutInput()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signout(res, req)

	// Assert
	verify.Should(t, spies.SignoutUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignout_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SignoutUsecaseSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.Signout(res, req)

	// Assert
	verify.Should(t, code).Be(spies.SignoutUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.SignoutUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignout_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SignoutUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	result, code, _ := sut.Signout(res, req)

	// Assert
	success := result.(*signout.Output)
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Message).Be(spies.SignoutUsecaseSpy.SuccessResult["Execute"].Message)
}
