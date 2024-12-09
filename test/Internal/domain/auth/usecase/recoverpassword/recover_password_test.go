package recoverpassword_test

import (
	"bytes"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/usecase/recoverpassword/recoverpasswordfixture"
	"testing"
)

func Test_GivenRecoverPasswordExecute_WhenInputTokenInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	invalidInput, expectedError := fixture.GetInvalidInputWithError()
	sut, _, _, _ := fixture.NewSut()

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
	sut, hasherSpy, settingsSpy, _ := fixture.NewSut()

	// Act
	sut.Execute(expectedInput)

	// Assert
	verify.Should(t, hasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput.Username)
	verify.Should(t, bytes.Equal(hasherSpy.Params["HashWithSalt:serverSalt"].([]byte), settingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenRecoverPasswordExecute_WhenHashWithSaltInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, _ := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, hasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenRecoverPasswordExecute_WhenHashWithSaltError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, _ := fixture.NewSut()
	hasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(resultapp.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(hasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenRecoverPasswordExecute_WhenHashWithSaltSuccess_ThenEnsureCallGetByUserNameWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, userRepoSpy := fixture.NewSut()
	hasherSpy.DefineHashWithSaltSuccess("fake-success-result")

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, userRepoSpy.Params["GetByUserName:username"]).Be(hasherSpy.SuccessResult["HashWithSalt"])
}

func Test_GivenRecoverPasswordExecute_WhenGetByUserNameInvoked_ThenEnsureCalssOnce(t *testing.T) {
	// Arrange
	sut, _, _, userRepoSpy := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, userRepoSpy.CallsCount["GetByUserName"]).Be(1)
}

func Test_GivenRecoverPasswordExecute_WhenGetByUserNameError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, _, _, userRepoSpy := fixture.NewSut()
	userRepoSpy.DefineError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(resultapp.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(userRepoSpy.ErrorResult["GetByUserName"])
}
