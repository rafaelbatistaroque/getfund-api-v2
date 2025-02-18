package signout_test

import (
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/external/domain/auth/core/usecase/signout/signout_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignoutExecute_WhenInputInvalid_ThenEnsureReturnErro(t *testing.T) {
	// Arrange
	invalidInput, errorInput := fixture.GetInvalidInputWithError()
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(errorInput.Code)
	verify.Should(t, err.Message).Be(errorInput.Message)
}

func Test_GivenSignoutExecute_WhenValidInput_ThenEnsureCallDeleteSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetValidInput()
	sut, sessionSpy := fixture.NewSut()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, sessionSpy.Params["DeleteSession:token"]).Be(validInput.Token)
}

func Test_GivenSignoutExecute_WhenDeleteSessionInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, sessionSpy := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, sessionSpy.CallsCount["DeleteSession"]).Be(1)
}

func Test_GivenSignoutExecute_WhenDeleteSessionError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, sessionSpy := fixture.NewSut()
	sessionSpy.DefineDeleteSessionError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(sessionSpy.ErrorResult["DeleteSession"])
}

func Test_GivenSignoutExecute_WhenDeleteSessionSuccess_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, sessionSpy := fixture.NewSut()
	sessionSpy.DefineDeleteSessionSuccess()

	// Act
	result, _ := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, result.Message).Be("user disconnected")
}
