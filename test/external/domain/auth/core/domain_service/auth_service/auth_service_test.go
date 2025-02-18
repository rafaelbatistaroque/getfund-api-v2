package auth_service_test

import (
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/external/domain/auth/core/domain_service/auth_service/auth_service_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenAuthenticate_WhenInit_ThenEnsureCallsGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Authenticate(expectedInputText, "")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(expectedInputText)
}

func Test_GivenAuthenticate_WhenGetAuthenticatedUserByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, spies.AuthRepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(1)
}

func Test_GivenAuthenticate_WhenGetAuthenticatedUserByUsernameError_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameError()

	// Act
	_, err := sut.Authenticate("fake-username", "")

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, spies.AuthRepoSpy.ErrorResult["GetAuthenticatedUserByUsername"]).Be(err.Message)
	verify.Should(t, result_app.UNAUTHORIZED_CODE).Be(err.Code)
}

func Test_GivenAuthenticate_WhenGetAuthenticatedUserByUsernameSuccess_ThenCallIsMatchWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedPassword := "fake-password"
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Authenticate("fake-username", expectedPassword)

	// Assert
	authenticatedUser := spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_dto.AuthenticatedUserDto)
	verify.Should(t, spies.HasherSpy.Params["IsMatch:inputHashed"]).Be(authenticatedUser.Password)
	verify.Should(t, spies.HasherSpy.Params["IsMatch:inputText"]).Be(expectedPassword)
}

func Test_GivenAuthenticate_WhenIsMatchInvoked_ThenCallIsMatchOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["IsMatch"]).Be(1)
}

func Test_GivenAuthenticate_WhenIsMatchFalse_ThenEnsureReturnUnauthorizedError(t *testing.T) {
	// Arrange
	expectedPasswordError := errors.New("invalid password")
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
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
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.HasherSpy.DefineIsMatchSuccess()

	// Act
	sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, spies.MapperSpy.Params["ToSessionModel:authenticatedUser"]).Be(spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"])
}

func Test_GivenAuthenticate_WhenSuccess_ThenEnsureReturnoSessionModelFilled(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.AuthRepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.HasherSpy.DefineIsMatchSuccess()
	authenticatedUser := spies.AuthRepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*auth_dto.AuthenticatedUserDto)
	spies.MapperSpy.DefineToSessionModelSuccess(authenticatedUser)

	// Act
	result, _ := sut.Authenticate("fake-username", "fake-password")

	// Assert
	verify.Should(t, result.ID).Be(authenticatedUser.Id)
	verify.Should(t, result.IsAdmin).Be(authenticatedUser.IsAdmin)
}
