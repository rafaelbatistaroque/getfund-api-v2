package auth_gateway_test

import (
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/signup_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignup_WhenDecodeError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpWithBody("body-with-error")

	// Act
	_, code, err := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenSignup_WhenDecodeSuccess_ThenEnsureCallOnceExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, SignupSpy := fixture.NewSut()
	expectedInput := fixture.GetSignupInput()
	res, req := fixture.GetHttpWithBody(expectedInput)

	// Act
	sut.Signup(res, req)

	// Assert
	verify.Should(t, SignupSpy.CallsCount["Execute"]).Be(1)
	verify.Should(t, SignupSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignup_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, SignupSpy := fixture.NewSut()
	SignupSpy.DefineError()
	res, req := fixture.GetHttpDefault()

	// Act
	_, code, err := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(SignupSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(SignupSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignup_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, SignupSpy := fixture.NewSut()
	SignupSpy.DefineSuccess()
	res, req := fixture.GetHttpDefault()

	// Act
	signed, code, _ := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SUCCESS_CODE)
	verify.Should(t, signed).Be(SignupSpy.SuccessResult["Execute"])
}
