package auth_parser_test

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/port/parser/signout_parser_fixture"
	"testing"
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
	expectedInput := fixture.GetSignoutInput()
	sut, signoutSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.Signout(res, req)

	// Assert
	verify.Should(t, signoutSpy.Params["Execute:input"]).Be(expectedInput)
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
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, success.Message).Be(signoutSpy.SuccessResult["Execute"].Message)
}
