package signin_application_test

import (
	"errors"
	shared_error "getfund-api-v2/internal/shared/error"
	fixtures "getfund-api-v2/test/internal/domain/auth/core/usecase/signin/signin_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSigninExecute_WhenSigninInputUserNameInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithUserNameInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Message.Error()).Be(errorInput.Message.Error())
	verify.Should(t, err.Code).Be(errorInput.Code)
}

func Test_GivenSigninExecute_WhenSigninInputPasswordInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithPasswordInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Message.Error()).Be(errorInput.Message.Error())
	verify.Should(t, err.Code).Be(errorInput.Code)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsWithCorrectParameters(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	correctParam := fixtures.GetValidInput()

	// Act
	sut.Execute(correctParam)

	// Assert
	verify.Should(t, spies.AuthServiceSpy.Params["password"]).Be(correctParam.Password)
	verify.Should(t, spies.AuthServiceSpy.Params["username"]).Be(correctParam.Username)
	verify.Should(t, spies.AuthServiceSpy.Params).Len(2)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	correctParam := fixtures.GetValidInput()

	// Act
	sut.Execute(correctParam)

	// Assert
	verify.Should(t, spies.AuthServiceSpy.CallsCount).Be(1)
}

func Test_GivenSigninExecute_WhenAuthenticateError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	anyError := shared_error.New(shared_error.UNAUTHORIZED_CODE, errors.New("fake-message"))
	spies.AuthServiceSpy.DefineNotAuthenticate(anyError.Code, anyError.Message)

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(anyError.Code)
	verify.Should(t, err.Message.Error()).Be(anyError.Message.Error())
}

func Test_GivenSigninExecute_WhenAuthenticateSuccess_ThenEnsureCallsSaveSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	spies.AuthServiceSpy.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, spies.SessionSpy.Params["SaveSession:session"]).Be(spies.AuthServiceSpy.SuccessResult)
}

func Test_GivenSigninExecute_WhenSaveSessionInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, spies.SessionSpy.CallsCount["SaveSession"]).Be(1)
}

func Test_GivenSigninExecute_WhenSaveSessionError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	spies.SessionSpy.DefineSaveSessionError()

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).NotNil()
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperWithCorrectParameters(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	spies.SessionSpy.DefineSaveSessionSuccess()
	spies.AuthServiceSpy.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, spies.MapperSpy.Params["ToOutput:token"]).Be(spies.SessionSpy.SuccessResult["SaveSession"])
	verify.Should(t, spies.MapperSpy.Params["ToOutput:session"]).Be(spies.AuthServiceSpy.SuccessResult)
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperToOutputOnce(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, spies.MapperSpy.CallsCount["ToOutput"]).Be(1)
}

func Test_GivenSigninExecute_WhenMapperInvoke_ThenEnsureReturnOutputWithSession(t *testing.T) {
	// Arrange
	sut, spies := fixtures.NewSut()
	spies.MapperSpy.ForceReturn = false
	spies.AuthServiceSpy.DefineAuthenticate()
	spies.SessionSpy.DefineSaveSessionSuccess()

	// Act
	result, _ := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, result).NotNil()
	verify.Should(t, result.Token).Be(spies.SessionSpy.SuccessResult["SaveSession"])
	verify.Should(t, result.Session.ID).Be(spies.AuthServiceSpy.SuccessResult.ID)
	verify.Should(t, result.Session.FirstName).Be(spies.AuthServiceSpy.SuccessResult.FirstName)
	verify.Should(t, result.Session.IsAdmin).Be(spies.AuthServiceSpy.SuccessResult.IsAdmin)
}
