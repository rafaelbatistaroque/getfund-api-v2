package user_respository_proxy_test

import (
	"bytes"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	fixture "getfund-api-v2/test/internal/domain/user/adapter/proxy/user_respository_proxy/user_respository_proxy_fixture"
	"testing"

	"github.com/google/uuid"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenCreateUser_WhenInit_ThenEnsureCallHasherMethodsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetFilledActivationUserDto()

	// Act
	sut.CreateUser(expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][0]).Be(expectedInput.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][0].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][1]).Be(expectedInput.LastName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][1].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][2]).Be(expectedInput.Email)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][2].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][3]).Be(expectedInput.MainSocialNetwork)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][3].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][4]).Be(expectedInput.RegisteredUrl)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][4].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(expectedInput.Password)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput.Username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenCreateUser_WhenHasherMethodsInvoked_ThenEnsureCallsCorrectTimes(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["Encrypt"]).Be(5)
	verify.Should(t, spies.HasherSpy.CallsCount["HashAndMerge"]).Be(1)
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenCreateUser_WhenHashWithSaltError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenCreateUser_WhenHasherMethodsSuccess_ThenEnsureCallCreateUserWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineEncryptSuccess()
	spies.HasherSpy.DefineHashAndMergeSuccess(uuid.NewString())
	spies.HasherSpy.DefineHashWithSaltSuccess(uuid.NewString())

	// Act
	sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	createUserParams := spies.RepoSpy.Params["CreateUser:user"].(*user_dto.ActivationUserDto)
	verify.Should(t, createUserParams.FirstName).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][0])
	verify.Should(t, createUserParams.LastName).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][1])
	verify.Should(t, createUserParams.Email).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][2])
	verify.Should(t, createUserParams.MainSocialNetwork).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][3])
	verify.Should(t, createUserParams.RegisteredUrl).Be(spies.HasherSpy.SuccessResultByCall["Encrypt"][4])
	verify.Should(t, createUserParams.Password).Be(spies.HasherSpy.SuccessResult["HashAndMerge"])
	verify.Should(t, createUserParams.Username).Be(spies.HasherSpy.SuccessResult["HashWithSalt"])
}

func Test_GivenCreateUser_WhenRepoCreateUserInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")

	// Act
	sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["CreateUser"]).Be(1)
}

func Test_GivenCreateUser_WhenRepoCreateUserError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")
	spies.RepoSpy.DefineCreateUserError()

	// Act
	_, err := sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["CreateUser"])
}

func Test_GivenCreateUser_WhenRepoCreateUserSuccess_ThenEnsureReturnCorrectUserDto(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess("")
	spies.RepoSpy.DefineCreateUserSuccess()

	// Act
	userReturned, _ := sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, userReturned).Be(spies.RepoSpy.SuccessResult["CreateUser"])
}

func Test_GivenUserExistsByUsername_WhenInit_ThenEnsureCallHashWithSaltWithCorrectTimes(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetFakeUsername()

	// Act
	sut.UserExistsByUsername(expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenUserExistsByUsername_WhenHashWithSaltInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.UserExistsByUsername(fixture.GetFakeUsername())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenUserExistsByUsername_WhenHashWithSaltError_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.UserExistsByUsername(fixture.GetFakeUsername())

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenUserExistsByUsername_WhenHashWithSaltSuccess_ThenEnsureCallRepoUserExistsByUsernameWithCorrectPaameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess(uuid.NewString())

	// Act
	sut.UserExistsByUsername(fixture.GetFakeUsername())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UserExistsByUsername:username"]).Be(spies.HasherSpy.SuccessResult["HashWithSalt"])
}

func Test_GivenUserExistsByUsername_WhenUserExistsByUsernameError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineUserExistsByUsernameError()

	// Act
	_, err := sut.UserExistsByUsername(fixture.GetFakeUsername())

	// Assert
	verify.Should(t, err).Be(spies.RepoSpy.ErrorResult["UserExistsByUsername"])
}
