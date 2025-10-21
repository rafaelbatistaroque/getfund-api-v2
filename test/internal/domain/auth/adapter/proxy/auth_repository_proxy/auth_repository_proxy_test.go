package auth_repository_proxy_test

import (
	"bytes"
	"getfund-api-v2/internal/domain/auth/core/dto"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/proxy/auth_repository_proxy/auth_repository_proxy_fixture"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetAuthenticatedUserByUsername_WhenInit_ThenEnsureCallOnceHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()

	// Act
	sut.GetAuthenticatedUserByUsername(expectedInputText)

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInputText)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
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

func Test_GivenGetAuthenticatedUserByUsername_WhenHashWithSaltResultEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, err.Error()).Be("error on get authenticated user")
}

func Test_GivenGetAuthenticatedUserByUsername_WhenHashWithSaltSuccess_ThenEnsureCallOnceRepositoryGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedParameter := "fake-username-hashed"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess(expectedParameter)
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(1)
	verify.Should(t, spies.RepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(expectedParameter)
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameError_ThenEnsureStartFallback(t *testing.T) {
	// Arrange
	username := "fake-username"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError()

	// Act
	sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSaltLegacy"]).Be(1)
	verify.Should(t, spies.HasherSpy.Params["HashWithSaltLegacy:inputText"]).Be(username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSaltLegacy:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameNil_ThenEnsureStartFallback(t *testing.T) {
	// Arrange
	username := "fake-username"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")

	// Act
	sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSaltLegacy"]).Be(1)
	verify.Should(t, spies.HasherSpy.Params["HashWithSaltLegacy:inputText"]).Be(username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSaltLegacy:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackHashWithSaltLegacyError_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	username := "fake-username"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.HasherSpy.DefineHashWithSaltLegacyError()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, err.Error()).Be("error on get authenticated user")
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackHashWithSaltLegacySuccess_ThenEnsureCallOnceGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	usernameHashed := "fake-username-hashed-legacy"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.HasherSpy.DefineHashWithSaltLegacySuccess(usernameHashed)

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(2)
	verify.Should(t, spies.RepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(usernameHashed)
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackGetAuthenticatedUserByUsernameResultError_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError() // para a primeira e segunda chamada

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["GetAuthenticatedUserByUsername"])
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackGetAuthenticatedUserByUsernameResultSuccess_ThenEnsureCallOnceUpdateUsernameHashWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError()   //para a primeira chamada
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess() //para a segunda chamada

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.RepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*dto.AuthenticatedUserDto)
	verify.Should(t, spies.RepoSpy.CallsCount["UpdateUsernameHash"]).Be(1)
	verify.Should(t, spies.RepoSpy.Params["UpdateUsernameHash:id"]).Be(authenticatedUser.Id)
	verify.Should(t, spies.RepoSpy.Params["UpdateUsernameHash:username"]).Be(spies.HasherSpy.SuccessResult["HashWithSalt"])
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackUpdateUsernameHashError_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError() // para a primeira e segunda chamada
	spies.RepoSpy.DefineUpdateUsernameHashError()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["UpdateUsernameHash"])
}

func Test_GivenGetAuthenticatedUserByUsername_WhenFallBackUpdateUsernameHashSuccess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError()   // para a primeira
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess() //para a segunda chamada

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.RepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*dto.AuthenticatedUserDto)
	verify.Should(t, spies.HasherSpy.CallsCount["DecryptMerged"]).Be(1)
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(authenticatedUser.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["DecryptMerged:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameSuccess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.RepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*dto.AuthenticatedUserDto)
	verify.Should(t, spies.HasherSpy.CallsCount["DecryptMerged"]).Be(1)
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(authenticatedUser.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["DecryptMerged:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenGetAuthenticatedUserByUsernameSuccess_ThenEnsureReturnModelDeserialized(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.HasherSpy.DefineDecryptMergedSuccess("fake-first-name")

	// Act
	result, _ := sut.GetAuthenticatedUserByUsername("fake-username")

	// Assert
	authenticatedUser := spies.RepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*dto.AuthenticatedUserDto)
	verify.Should(t, result.Id).Be(authenticatedUser.Id)
	verify.Should(t, result.FirstName).Be(spies.HasherSpy.SuccessResult["DecryptMerged"])
	verify.Should(t, result.IsAdmin).Be(authenticatedUser.IsAdmin)
	verify.Should(t, result.Password).Be(authenticatedUser.Password)
}

func Test_GivenUpdatePassword_WhenInit_ThenEnsureCallHashAndMergeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := "fake-username"

	// Act
	sut.UpdatePassword(0, expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(expectedInput)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenUpdatePassword_WhenHashAndMergeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UpdatePassword(0, "")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashAndMerge"]).Be(1)
}

func Test_GivenUpdatePassword_WhenHashAndMergeSuccess_ThenEnsureCallRepositoryUpdatePasswordWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedParamId := 1
	spies.HasherSpy.DefineHashAndMergeSuccess("fake-password-hashed")

	// Act
	sut.UpdatePassword(expectedParamId, "")

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UpdatePassword:id"]).Be(expectedParamId)
	verify.Should(t, spies.RepoSpy.Params["UpdatePassword:value"]).Be(spies.HasherSpy.SuccessResult["HashAndMerge"].(string))
}

func Test_GivenUpdatePassword_WhenUpdatePasswordInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UpdatePassword(0, "")

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UpdatePassword"]).Be(1)
}

func Test_GivenUpdatePassword_WhenUpdatePasswordError_ThenEnsureErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineUpdatePasswordError()

	// Act
	err := sut.UpdatePassword(0, "")

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["UpdatePassword"])
}

func Test_GivenUpdatePassword_WhenUpdatePasswordSuccess_ThenEnsureNull(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	result := sut.UpdatePassword(0, "")

	// Assert
	verify.Should(t, result).Nil()
}

func Test_GivenSignup_WhenInit_ThenEnsureCallHasherMethodsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetFilledActivationUserDto()

	// Act
	sut.Signup(expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][0]).Be(expectedInput.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][0].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][1]).Be(expectedInput.LastName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][1].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(expectedInput.Password)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput.Username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenSignup_WhenHasherMethodsInvoked_ThenEnsureCallsCorrectTimes(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["Encrypt"]).Be(2)
	verify.Should(t, spies.HasherSpy.CallsCount["HashAndMerge"]).Be(1)
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenSignup_WhenHashWithSaltError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenSignup_WhenHasherMethodsSuccess_ThenEnsureCallSignupWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineEncryptSuccess()
	spies.HasherSpy.DefineHashAndMergeSuccess(uuid.NewString())
	spies.HasherSpy.DefineHashWithSaltSuccess(uuid.NewString())

	// Act
	sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	SignupParams := spies.RepoSpy.Params["Signup:user"].(*dto.ActivationUserDto)
	verify.Should(t, SignupParams.FirstName).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][0])
	verify.Should(t, SignupParams.LastName).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][1])
	verify.Should(t, SignupParams.Password).Be(spies.HasherSpy.SuccessResult["HashAndMerge"])
	verify.Should(t, SignupParams.Username).Be(spies.HasherSpy.SuccessResult["HashWithSalt"])
}

func Test_GivenSignup_WhenRepoSignupInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")

	// Act
	sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["Signup"]).Be(1)
}

func Test_GivenSignup_WhenRepoSignupError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")
	spies.RepoSpy.DefineSignupError()

	// Act
	_, err := sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["Signup"])
}

func Test_GivenSignup_WhenRepoSignupSuccess_ThenEnsureReturnCorrectUserDto(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")
	spies.RepoSpy.DefineSignupSuccess()

	// Act
	userReturned, _ := sut.Signup(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, userReturned).Be(spies.RepoSpy.SuccessResult["Signup"])
}

func Test_GivenUserExists_WhenInit_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()

	// Act
	sut.UserExists(expectedInputText)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInputText)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenUserExists_WhenHashWithSaltError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.UserExists("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenUserExists_WhenHashWithSaltResultEmpty_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.UserExists("fake-username")

	// Assert
	verify.Should(t, err.Error()).Be("error on get authenticated user")
}

func Test_GivenUserExists_WhenHashWithSaltInvoked_ThenEnsureCallHashWithSaltOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UserExists("fake-username")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenUserExists_WhenHashWithSaltSuccess_ThenEnsureCallRepositoryUserExistsWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedParameter := "fake-username-hashed"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess(expectedParameter)

	// Act
	sut.UserExists("fake-username")

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UserExists:username"]).Be(expectedParameter)
}

func Test_GivenUserExists_WhenUserExistsInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")

	// Act
	sut.UserExists("fake-username")

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UserExists"]).Be(1)
}

func Test_GivenUserExists_WhenUserExistsError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineUserExistsError()

	// Act
	_, err := sut.UserExists("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["UserExists"])
}

func Test_GivenUserExists_WhenUserExistsFound_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")
	spies.RepoSpy.DefineUserExistsSuccessUserFound()

	// Act
	_, err := sut.UserExists("fake-username")

	// Assert
	verify.Should(t, err.Error()).Be("error on get authenticated user")
}

func Test_GivenUserExists_WhenUserExistsNotFound_ThenEnsureReturnUserFound(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("fake-username-hashed")

	// Act
	result, err := sut.UserExists("fake-username")

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, result).Nil()
}
