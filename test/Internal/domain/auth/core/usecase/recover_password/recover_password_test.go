package recover_password_test

import (
	"bytes"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/pkg/bus/event"
	fixture "getfund-api-v2/test/internal/domain/auth/core/usecase/recover_password/recover_password_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenRecoverPasswordExecute_WhenInputTokenInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	invalidInput, expectedError := fixture.GetInvalidInputWithError()
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(expectedError.Code)
	verify.Should(t, err.Message).Be(expectedError.Message)
}

func Test_GivenRecoverPasswordExecute_WhenValidInput_ThenEnsureCallGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInput := fixture.GetValidInput()
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(expectedInput)

	// Assert
	verify.Should(t, spies.AuthRepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(expectedInput.Username)
}

func Test_GivenRecoverPasswordExecute_WhenGetAuthenticatedUserByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.AuthRepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(1)
}

func Test_GivenRecoverPasswordExecute_WhenGetAuthenticatedUserByUsernameError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(spies.AuthRepoSpy.ErrorResult["GetAuthenticatedUserByUsername"])
}

func Test_GivenRecoverPasswordExecute_WhenGetAuthenticatedUserByUsernameSuccess_ThenEnsureCallGetRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(8)
}

func Test_GivenRecoverPasswordExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenRecoverPasswordExecute_WhenGetRandomCodeError_ThenEnsureReturnErrorFromWithServerErrorCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.HasherSpy.ErrorResult["GetRandomCode"])
}

func Test_GivenRecoverPasswordExecute_WhenGetRandomCodeSuccess_ThenEnsureCallsHashWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineGetRandomCodeSuccess()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["Hash:inputText"]).Be(spies.HasherSpy.SuccessResult["GetRandomCode"].(string))
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["Hash:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenRecoverPasswordExecute_WhenHashError_ThenEnsureReturnErrorFromWithServerErrorCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.HasherSpy.ErrorResult["Hash"])
}

func Test_GivenRecoverPasswordExecute_WhenHashSuccess_ThenEnsureCallCacheSetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetValidInput()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.HasherSpy.DefineDecryptMergedSuccess("fake-first-name")
	spies.HasherSpy.DefineHashSuccess()
	hashCode := spies.HasherSpy.SuccessResult["Hash"].(*security.Hashing).Data
	expectedKey := "recovery_password_" + hashCode
	authenticatedUser := spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_dto.AuthenticatedUserDto)
	expectedValue := auth_dto.ForgetPasswordDto{
		Username:     validInput.Username,
		FirstName:    authenticatedUser.FirstName,
		RecoveryLink: spies.SettingsSpy.GetBaseUrl() + "/new-password/" + hashCode,
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedKey)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(expectedValue)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(time.Hour)
}

func Test_GivenRecoverPasswordExecute_WhenCacheSetError_ThenEnsureReturnErrorFromWithServerErrorCode(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheSetError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.CacheSpy.ErrorResult["Set"])
}

func Test_GivenRecoverPasswordExecute_WhenCacheSetSuccess_ThenEnsureCallPublishWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashSuccess()
	expectedPaylod := "recovery_password_" + spies.HasherSpy.SuccessResult["Hash"].(*security.Hashing).Data

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.EventBusSpy.Params["EmitWithPayload:event"]).Be(&event.RecoverPasswordStarted{})
	verify.Should(t, spies.EventBusSpy.Params["EmitWithPayload:payload"]).Be(expectedPaylod)
}

func Test_GivenRecoverPasswordExecute_WhenPublishWithPayload_ThenEnsureReturnOutput(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	result, _ := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, result.Message).Be("recover password started")
}
