package auth_service_test

import (
	"bytes"
	"errors"
	authmodel "getfund-api-v2/internal/domain/auth/adapter/model"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/domain_service/auth_service/auth_service_fixture"
	"testing"
)

func Test_GivenAuthenticate_WhenInit_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, settingsSpy, _, hasherSpy, _ := fixture.NewSut()
	hasherSpy.DefineHashWithSaltError()

	// Act
	sut.Authenticate(expectedInputText, "")

	// Assert
	verify.Should(t, hasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInputText)
	verify.Should(t, bytes.Equal(hasherSpy.Params["HashWithSalt:serverSalt"].([]byte), settingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenAuthenticate_WhenInit_ThenEnsureCallHashWithSaltOnce(t *testing.T) {
	// Arrange
	sut, _, _, hasherSpy, _ := fixture.NewSut()
	hasherSpy.DefineHashWithSaltError()

	// Act
	sut.Authenticate("", "")

	// Assert
	verify.Should(t, hasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenAuthenticate_WhenHashWithSaltError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, hasherSpy, _ := fixture.NewSut()
	hasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.Authenticate("", "")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenAuthenticate_WhenHashWithSaltSuccess_ThenEnsureCallsGetByUserNameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedValue := "fake-username-hashed"
	sut, _, userRepo, hasherSpy, _ := fixture.NewSut()
	hasherSpy.DefineHashWithSaltSuccess(expectedValue)
	userRepo.DefineSuccess()

	// Act
	sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, userRepo.Params["GetByUserName:username"]).Be(expectedValue)
}

func Test_GivenAuthenticate_WhenHashWithSaltSuccess_ThenEnsureCallsGetByUserNameOnce(t *testing.T) {
	// Arrange
	sut, _, userRepo, _, _ := fixture.NewSut()
	userRepo.DefineSuccess()

	// Act
	sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, userRepo.CallsCount["GetByUserName"]).Be(1)
}

func Test_GivenAuthenticate_WhenGetByUserNameError_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	sut, _, userRepo, _, _ := fixture.NewSut()
	userRepo.DefineError()

	// Act
	_, err := sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, userRepo.ErrorResult["GetByUserName"]).Be(err.Message)
	verify.Should(t, result_app.UNAUTHORIZED_CODE).Be(err.Code)
}

func Test_GivenAuthenticate_WhenGetByUserNameSuccess_ThenCallIsMatchWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedPassword := "fake-password"
	sut, settingsSpy, userRepo, hasherSpy, _ := fixture.NewSut()
	userRepo.DefineSuccess()

	// Act
	sut.Authenticate("fake-username", expectedPassword)

	// Assert
	user := userRepo.SuccessResult["GetByUserName"].(*authmodel.UserModel)
	verify.Should(t, hasherSpy.Params["IsMatch:inputHashed"]).Be(user.Password)
	verify.Should(t, hasherSpy.Params["IsMatch:inputText"]).Be(expectedPassword)
	verify.Should(t, bytes.Equal(hasherSpy.Params["HashWithSalt:serverSalt"].([]byte), settingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenAuthenticate_WhenGetByUserNameSuccess_ThenCallIsMatchOnce(t *testing.T) {
	// Arrange
	sut, _, userRepo, hasherSpy, _ := fixture.NewSut()
	userRepo.DefineSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, hasherSpy.CallsCount["IsMatch"]).Be(1)
}

func Test_GivenAuthenticate_WhenIsMatchFalse_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	expectedPasswordError := errors.New("invalid password")
	sut, _, userRepo, hasherSpy, _ := fixture.NewSut()
	userRepo.DefineSuccess()
	hasherSpy.DefineIsMatchError()

	// Act
	_, err := sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message).Be(expectedPasswordError)
}

func Test_GivenAuthenticate_WhenIsMatchSuccess_ThenEnsureCallToSessionModelWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, _, userRepo, hasherSpy, mapperSpy := fixture.NewSut()
	userRepo.DefineSuccess()
	hasherSpy.DefineIsMatchSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, mapperSpy.Params["ToSessionModel:user"]).Be(userRepo.SuccessResult["GetByUserName"])
}

func Test_GivenAuthenticate_WhenSuccess_ThenEnsureReturnoSessionModelFilled(t *testing.T) {
	// Arrange
	sut, _, userRepo, hasherSpy, mapperSpy := fixture.NewSut()
	userRepo.DefineSuccess()
	hasherSpy.DefineIsMatchSuccess()
	user := userRepo.SuccessResult["GetByUserName"].(*authmodel.UserModel)
	mapperSpy.DefineToSessionModelSuccess(user)

	// Act
	result, _ := sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, result.ID).Be(user.Id)
	verify.Should(t, result.IsAdmin).Be(user.IsAdmin)
}
