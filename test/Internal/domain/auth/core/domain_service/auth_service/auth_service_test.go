package auth_service_test

import (
	"errors"
	authmodel "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/core/domain_service/auth_service/auth_service_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenAuthenticate_WhenInit_ThenEnsureCallsGetByUserNameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()

	// Act
	sut.Authenticate(expectedInputText, "")

	// Assert
	verify.Should(t, spies.UserRepoSpy.Params["GetByUserName:username"]).Be(expectedInputText)
}

func Test_GivenAuthenticate_WhenGetByUserNameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, spies.UserRepoSpy.CallsCount["GetByUserName"]).Be(1)
}

func Test_GivenAuthenticate_WhenGetByUserNameError_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameError()

	// Act
	_, err := sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, spies.UserRepoSpy.ErrorResult["GetByUserName"]).Be(err.Message)
	verify.Should(t, result_app.UNAUTHORIZED_CODE).Be(err.Code)
}

func Test_GivenAuthenticate_WhenGetByUserNameSuccess_ThenCallIsMatchWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedPassword := "fake-password"
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()

	// Act
	sut.Authenticate("fake-username", expectedPassword)

	// Assert
	user := spies.UserRepoSpy.SuccessResult["GetByUserName"].(*authmodel.UserModel)
	verify.Should(t, spies.HasherSpy.Params["IsMatch:inputHashed"]).Be(user.Password)
	verify.Should(t, spies.HasherSpy.Params["IsMatch:inputText"]).Be(expectedPassword)
}

func Test_GivenAuthenticate_WhenIsMatchInvoked_ThenCallIsMatchOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["IsMatch"]).Be(1)
}

func Test_GivenAuthenticate_WhenIsMatchFalse_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	expectedPasswordError := errors.New("invalid password")
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()
	spies.HasherSpy.DefineIsMatchError()

	// Act
	_, err := sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message).Be(expectedPasswordError)
}

func Test_GivenAuthenticate_WhenIsMatchSuccess_ThenEnsureCallToSessionModelWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()
	spies.HasherSpy.DefineIsMatchSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, spies.MapperSpy.Params["ToSessionModel:user"]).Be(spies.UserRepoSpy.SuccessResult["GetByUserName"])
}

func Test_GivenAuthenticate_WhenSuccess_ThenEnsureReturnoSessionModelFilled(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.UserRepoSpy.DefineGetByUserNameSuccess()
	spies.HasherSpy.DefineIsMatchSuccess()
	user := spies.UserRepoSpy.SuccessResult["GetByUserName"].(*authmodel.UserModel)
	spies.MapperSpy.DefineToSessionModelSuccess(user)

	// Act
	result, _ := sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, result.ID).Be(user.Id)
	verify.Should(t, result.IsAdmin).Be(user.IsAdmin)
}
