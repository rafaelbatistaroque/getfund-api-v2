package signout

import (
	"getfund-api-v2/internal/pkg/verify"
	"getfund-api-v2/internal/shared/resultapp"
	fixture "getfund-api-v2/test/domain/auth/usecase/signout/signoutfixture"
	"testing"
)

func Test_GivenSignoutExecute_WhenInputInvalid_ThenEnsureReturnErro(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInvalidInput()
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(resultapp.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message).Be(invalidInput.GetErrors())
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
	verify.Should(t, err.Code).Be(resultapp.SERVER_ERROR_CODE)
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
