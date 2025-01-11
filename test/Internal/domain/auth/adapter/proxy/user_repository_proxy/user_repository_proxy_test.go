package user_repository_proxy_test

import (
	"bytes"
	auth_model "getfund-api-v2/internal/domain/auth/core/auth_dto"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/proxy/user_repository_proxy/auth_repository_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetAuthenticatedUserByUsername_WhenInit_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()

	// Act
	sut.GetAuthenticatedUserByUsername(expectedInputText)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInputText)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenHashWithSaltInvoked_ThenEnsureCallHashWithSaltOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenGetAuthenticatedUserByUsername_WhenHashWithSaltError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenGetAuthenticatedUserByUsername_WhenHashWithSaltSuccess_ThenEnsureCallRepositoryGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedParameter := "fake-username-hashed"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess(expectedParameter)

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(expectedParameter)
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(1)
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameError()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.AuthRepoSpy.ErrorResult["GetAuthenticatedUserByUsername"])
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameSucess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_model.AuthenticatedUserDto)
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(authenticatedUser.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["DecryptMerged:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameSucess_ThenEnsureReturnModelDeserialized(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.HasherSpy.DefineDecryptMergedSuccess("fake-first-name")

	// Act
	result, _ := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_model.AuthenticatedUserDto)
	verify.Should(t, result.Id).Be(authenticatedUser.Id)
	verify.Should(t, result.FirstName).Be(spies.HasherSpy.SuccessResult["DecryptMerged"])
	verify.Should(t, result.IsAdmin).Be(authenticatedUser.IsAdmin)
	verify.Should(t, result.Password).Be(authenticatedUser.Password)
}

func Test_GivenUpdatePassword_WhenInit_ThenEnsureCallHashAndMergeWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInput := "fake-username"
	sut, spies := fixture.NewSut()

	// Act
	sut.UpdatePassword("", expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(expectedInput)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenUpdatePassword_WhenHashAndMergeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UpdatePassword("", "")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashAndMerge"]).Be(1)
}

func Test_GivenUpdatePassword_WhenHashAndMergeSuccess_ThenEnsureCallRepositoryUpdatePasswordWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedParamId := "fake-id"
	spies.HasherSpy.DefineHashAndMergeSuccess("fake-password-hashed")

	// Act
	sut.UpdatePassword(expectedParamId, "")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.Params["UpdatePassword:id"]).Be(expectedParamId)
	verify.Should(t, spies.AuthRepoSpy.Params["UpdatePassword:value"]).Be(spies.HasherSpy.SuccessResult["HashAndMerge"].(string))
}

func Test_GivenUpdatePassword_WhenUpdatePasswordInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UpdatePassword("", "")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.CallsCount["UpdatePassword"]).Be(1)
}

func Test_GivenUpdatePassword_WhenUpdatePasswordError_ThenEnsureErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineUpdatePasswordError()

	// Act
	err := sut.UpdatePassword("", "")

	// Assert
	verify.Should(t, err).Be(spies.AuthRepoSpy.ErrorResult["UpdatePassword"])
}

func Test_GivenUpdatePassword_WhenUpdatePasswordSuccess_ThenEnsureNull(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	result := sut.UpdatePassword("", "")

	// Assert
	verify.Should(t, result).Nil()
}
