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
	expectedInput := fixture.GetSigninInput()
	sut, signinSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signin(res, req)

	// Assert
	verify.Should(t, signinSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignin_WhenDecodeSuccess_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, signinSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signin(res, req)

	// Assert
	verify.Should(t, signinSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenSignin_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, signinSpy := fixture.NewSut()
	signinSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.Signin(res, req)

	// Assert
	verify.Should(t, code).Be(signinSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(signinSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignin_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, signinSpy := fixture.NewSut()
	signinSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	signed, code, _ := sut.Signin(res, req)

	// Assert
	success := signed.(*signin.Output)
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Token).Be(signinSpy.SuccessResult["Execute"].Token)
	verify.Should(t, success.Session.ID).Be(signinSpy.SuccessResult["Execute"].Session.ID)
	verify.Should(t, success.Session.FirstName).Be(signinSpy.SuccessResult["Execute"].Session.FirstName)
	verify.Should(t, success.Session.IsAdmin).Be(signinSpy.SuccessResult["Execute"].Session.IsAdmin)
}
