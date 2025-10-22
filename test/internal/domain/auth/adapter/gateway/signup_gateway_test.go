package auth_gateway_test

import (
	"encoding/json"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/signup_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignup_WhenDecodeError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	res, req := spies.GetHttpRequestResponse("body-with-error")

	// Act
	_, code, err := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenSignup_WhenDecodeSuccess_ThenEnsureCallOnceExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetSignupInput()
	body, _ := json.Marshal(expectedInput)
	res, req := spies.GetHttpRequestResponse(string(body))

	// Act
	sut.Signup(res, req)

	// Assert
	verify.Should(t, spies.SignupUsecaseSpy.CallsCount["Execute"]).Be(1)
	verify.Should(t, spies.SignupUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenSignup_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SignupUsecaseSpy.DefineError()
	res, req := spies.GetHttpRequestResponse("{}")

	// Act
	_, code, err := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(spies.SignupUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.SignupUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenSignup_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SignupUsecaseSpy.DefineSuccess()
	res, req := spies.GetHttpRequestResponse("{}")

	// Act
	signed, code, _ := sut.Signup(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SUCCESS_CODE)
	verify.Should(t, signed).Be(spies.SignupUsecaseSpy.SuccessResult["Execute"])
}
