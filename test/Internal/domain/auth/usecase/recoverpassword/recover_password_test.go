package recoverpassword_test

import (
	"bytes"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/usecase/recoverpassword/recoverpasswordfixture"
	"testing"
)

func Test_GivenRecoverPasswordExecute_WhenInputTokenInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	invalidInput, expectedError := fixture.GetInvalidInputWithError()
	sut, _, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(expectedError.Code)
	verify.Should(t, err.Message).Be(expectedError.Message)
}

func Test_GivenRecoverPasswordExecute_WhenValidInput_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInput := fixture.GetValidInput()
	sut, hasherSpy, settingsSpy := fixture.NewSut()

	// Act
	sut.Execute(expectedInput)

	// Assert
	verify.Should(t, hasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput.Username)
	verify.Should(t, bytes.Equal(hasherSpy.Params["HashWithSalt:serverSalt"].([]byte), settingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenRecoverPasswordExecute_WhenHashWithSaltInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, hasherSpy, _ := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, hasherSpy.CallsCount["HashWithSalt"]).Be(1)
}
