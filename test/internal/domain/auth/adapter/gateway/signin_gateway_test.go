package auth_gateway_test

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/signin_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignin_WhenDecodeError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("body-with-error")

	// Act
	_, code, err := sut.Signin(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenSignin_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetSigninInput()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signin(res, req)

	// Assert
	verify.Should(t, spies.SigninUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignin_WhenDecodeSuccess_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signin(res, req)

	// Assert
	verify.Should(t, spies.SigninUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenSignin_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SigninUsecaseSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.Signin(res, req)

	// Assert
	verify.Should(t, code).Be(spies.SigninUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.SigninUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignin_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SigninUsecaseSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	signed, code, _ := sut.Signin(res, req)

	// Assert
	success := signed.(*signin.Output)
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Token).Be(spies.SigninUsecaseSpy.SuccessResult["Execute"].Token)
	verify.Should(t, success.Session.ID).Be(spies.SigninUsecaseSpy.SuccessResult["Execute"].Session.ID)
	verify.Should(t, success.Session.FirstName).Be(spies.SigninUsecaseSpy.SuccessResult["Execute"].Session.FirstName)
	verify.Should(t, success.Session.IsAdmin).Be(spies.SigninUsecaseSpy.SuccessResult["Execute"].Session.IsAdmin)
}
