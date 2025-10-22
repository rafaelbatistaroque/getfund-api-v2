package auth_gateway_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/gateway/activate_user_gateway_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenActivateUser_WhenActivationCodeUrlParamNotFound_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	res, req := spies.GetHttpRequestResponseWithUrl("/user/activate/")

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Error()).Be("activation code is required")
}

func Test_GivenActivateUser_WhenActivationCodeUrlParamFound_ThenEnsureCallUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetActivateUserInput()
	url := fmt.Sprintf("/user/activate/%s", expectedInput.ActivationCode)
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, spies.ActivateUserUsecaseSpy.Params["Execute:input"]).Be(expectedInput)
}

func Test_GivenActivateUser_WhenUsecaseInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	url := fmt.Sprintf("/user/activate/%s", "fake-activation-code")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, spies.ActivateUserUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenActivateUser_WhenUsecaseError_ThenEnsureReturnCodeAndMessageFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.ActivateUserUsecaseSpy.DefineError()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(spies.ActivateUserUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.ActivateUserUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenActivateUser_WhenActivateUserUsecaseSuccess_ThenEnsureCallSigninUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.ActivateUserUsecaseSpy.DefineSuccessWithValue()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	activateUserOutput := spies.ActivateUserUsecaseSpy.SuccessResult["Execute"]
	verify.Should(t, spies.SigninUsecaseSpy.Params["Execute:input"]).Be(&signin.Input{
		Username: activateUserOutput.Username,
		Password: activateUserOutput.Password,
	})
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccessNUll_ThenEnsureApropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Error()).Be("error to activate user")
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccess_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.ActivateUserUsecaseSpy.DefineSuccess()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, spies.SigninUsecaseSpy.CallsCount["Execute"]).Be(1)
}

func Test_GivenActivateUser_WhenSigninUsecaseError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.ActivateUserUsecaseSpy.DefineSuccess()
	spies.SigninUsecaseSpy.DefineError()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	_, code, err := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(spies.SigninUsecaseSpy.ErrorResult["Execute"].Code)
	verify.Should(t, err).Be(spies.SigninUsecaseSpy.ErrorResult["Execute"].Message)
}

func Test_GivenActivateUser_WhenSigninUsecaseSuccess_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.ActivateUserUsecaseSpy.DefineSuccess()
	spies.SigninUsecaseSpy.DefineSuccess()
	url := fmt.Sprintf("/user/activate/%s", "valid-param")
	res, req := spies.GetHttpRequestResponseWithUrl(url)

	// Act
	result, code, _ := sut.ActivateUser(res, req)

	// Assert
	verify.Should(t, code).Be(shared_error.SUCCESS_CODE)
	verify.Should(t, result).Be(spies.SigninUsecaseSpy.SuccessResult["Execute"])
}
