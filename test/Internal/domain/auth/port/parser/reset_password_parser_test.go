package auth_parser_test

import (
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/port/parser/reset_password_parser_fixture"
	"testing"
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
