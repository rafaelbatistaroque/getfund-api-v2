package signin_application_test

import (
	"errors"
	auth_dto "getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/shared/result_app"
	fixtures "getfund-api-v2/test/internal/domain/auth/core/usecase/signin/signin_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSigninExecute_WhenSigninInputUserNameInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithUserNameInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Message.Error()).Be(errorInput.Message.Error())
	verify.Should(t, err.Code).Be(errorInput.Code)
}

func Test_GivenSigninExecute_WhenSigninInputPasswordInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithPasswordInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Message.Error()).Be(errorInput.Message.Error())
	verify.Should(t, err.Code).Be(errorInput.Code)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsWithCorrectParameters(t *testing.T) {
	// Arrange
	correctParam := fixtures.GetValidInput()
	sut, authService, _, _ := fixtures.NewSut()

	// Act
	sut.Execute(correctParam)

	// Assert
	verify.Should(t, authService.Params["password"]).Be(correctParam.Password)
	verify.Should(t, authService.Params["username"]).Be(correctParam.UserName)
	verify.Should(t, authService.Params).Len(2)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	correctParam := fixtures.GetValidInput()
	sut, authService, _, _ := fixtures.NewSut()

	// Act
	sut.Execute(correctParam)

	// Assert
	verify.Should(t, authService.CallsCount).Be(1)
}

func Test_GivenSigninExecute_WhenAuthenticateError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, authService, _, _ := fixtures.NewSut()
	anyError := result_app.New(result_app.UNAUTHORIZED_CODE, errors.New("fake-message"))
	authService.DefineNotAuthenticate(anyError.Code, anyError.Message)

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(anyError.Code)
	verify.Should(t, err.Message.Error()).Be(anyError.Message.Error())
}

func Test_GivenSigninExecute_WhenAuthenticateSuccess_ThenEnsureCallsMapperSessionToStringWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, authService, _, mapper := fixtures.NewSut()
	authService.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert

	verify.
		Should(t, mapper.Params["SessionToString:session"].(*auth_dto.SessionDto)).
		Be(authService.SuccessResult)
}

func Test_GivenSigninExecute_WhenMapperSessionToStringInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, _, mapper := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, mapper.CallsCount["SessionToString"]).Be(1)
}

func Test_GivenSigninExecute_WhenMapperSessionToStringError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, _, _, mapper := fixtures.NewSut()
	mapper.DefineError()

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).NotNil()
}

func Test_GivenSigninExecute_WhenSessionToStringSuccess_ThenEnsureCallsSaveSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, _, sessionService, mapper := fixtures.NewSut()
	mapper.DefineSuccess()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.
		Should(t, sessionService.Params["SaveSession:session"]).
		Be(mapper.SuccessResult["SessionToString"])
}

func Test_GivenSigninExecute_WhenSaveSessionInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, sessionService.CallsCount["SaveSession"]).Be(1)
}

func Test_GivenSigninExecute_WhenSaveSessionError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()
	sessionService.DefineSaveSessionError()

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).NotNil()
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperWithCorrectParameters(t *testing.T) {
	// Arrange
	sut, authService, sessionService, mapper := fixtures.NewSut()
	sessionService.DefineSaveSessionSuccess()
	authService.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, mapper.Params["ToOutput:token"]).Be(sessionService.SuccessResult["SaveSession"])
	verify.Should(t, mapper.Params["ToOutput:session"]).Be(authService.SuccessResult)
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperToOutputOnce(t *testing.T) {
	// Arrange
	sut, _, _, mapper := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, mapper.CallsCount["ToOutput"]).Be(1)
}

func Test_GivenSigninExecute_WhenMapperInvoke_ThenEnsureReturnOutputWithSession(t *testing.T) {
	// Arrange
	sut, authService, sessionService, mapper := fixtures.NewSut()
	mapper.ForceReturn = false
	authService.DefineAuthenticate()
	sessionService.DefineSaveSessionSuccess()

	// Act
	result, _ := sut.Execute(fixtures.GetValidInput())

	// Assert
	verify.Should(t, result).NotNil()
	verify.Should(t, result.Token).Be(sessionService.SuccessResult["SaveSession"])
	verify.Should(t, result.Session.ID).Be(authService.SuccessResult.ID)
	verify.Should(t, result.Session.FirstName).Be(authService.SuccessResult.FirstName)
	verify.Should(t, result.Session.IsAdmin).Be(authService.SuccessResult.IsAdmin == 1)
}
