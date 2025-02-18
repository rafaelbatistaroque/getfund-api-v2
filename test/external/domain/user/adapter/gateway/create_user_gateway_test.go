package user_gateway_test

import (
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/external/domain/user/adapter/gateway/create_user_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenCreateUser_WhenDecodeError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("body-with-error")

	// Act
	_, code, err := sut.CreateUser(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err).NotNil()
}

func Test_GivenCreateUser_WhenDecodeSuccess_ThenEnsureCallExecuteWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, createUserSpy := fixture.NewSut()
	expectedInput := fixture.GetCreateUserInput()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.CreateUser(res, req)

	// Assert
	verify.Should(t, createUserSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenCreateUser_WhenDecodeSuccess_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, createUserSpy := fixture.NewSut()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	sut.CreateUser(res, req)

	// Assert
	verify.Should(t, createUserSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenCreateUser_WhenExecuteError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, createUserSpy := fixture.NewSut()
	createUserSpy.DefineError()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	_, code, err := sut.CreateUser(res, req)

	// Assert
	verify.Should(t, code).Be(createUserSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(createUserSpy.ErrorResult["Execute"].Message)
}

func Test_GivenCreateUser_WhenExecuteSuccess_ThenEnsureReturnOutputWithSuccessCode(t *testing.T) {
	// Arrange
	sut, createUserSpy := fixture.NewSut()
	createUserSpy.DefineSuccess()
	res, req := fixture.GetHttpRequestResponse("")

	// Act
	signed, code, _ := sut.CreateUser(res, req)

	// Assert
	verify.Should(t, code).Be(result_app.SUCCESS_CODE)
	verify.Should(t, signed).Be(createUserSpy.SuccessResult["Execute"])
}
