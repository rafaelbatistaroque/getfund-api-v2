package signinapplication

import (
	"errors"
	appErr "getfund-api-v2/internal/pkg/applicationerror"
	"getfund-api-v2/internal/pkg/expect"
	appCode "getfund-api-v2/internal/pkg/helpers/applicationcode"
	fixtures "getfund-api-v2/internal/test/domain/auth/application/signinapplication/signinfixture"
	"testing"
)

func Test_GivenSigninExecute_WhenSigninInputUserNameInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithUserNameInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	expect.Equal(t, err.Message.Error(), errorInput.Message.Error())
	expect.Equal(t, err.Code, errorInput.Code)
}

func Test_GivenSigninExecute_WhenSigninInputPasswordInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixtures.NewSut()
	invalidInput, errorInput := fixtures.GetInputWithPasswordInvalid()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	expect.Equal(t, err.Message.Error(), errorInput.Message.Error())
	expect.Equal(t, err.Code, errorInput.Code)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsWithCorrectParameters(t *testing.T) {
	// Arrange
	correctParam := fixtures.GetValidInput()
	sut, authService, _, _ := fixtures.NewSut()

	// Act
	sut.Execute(correctParam)

	// Assert
	expect.Equal(t, authService.Params["password"], correctParam.Password)
	expect.Equal(t, authService.Params["username"], correctParam.UserName)
	expect.Len(t, authService.Params, 2)
}

func Test_GivenSigninExecute_WhenAuthenticateInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	correctParam := fixtures.GetValidInput()
	sut, authService, _, _ := fixtures.NewSut()

	// Act
	sut.Execute(correctParam)

	// Assert
	expect.Equal(t, authService.CallsCount, 1)
}

func Test_GivenSigninExecute_WhenAuthenticateError_ThenEnsureReturnErrorFrom(t *testing.T) {
	// Arrange
	sut, authService, _, _ := fixtures.NewSut()
	anyError := appErr.New(appCode.CODE_UNAUTHORIZED, errors.New("fake-message"))
	authService.DefineNotAuthenticate(anyError.Code, anyError.Message)

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.NotNil(t, err)
	expect.Equal(t, err.Code, anyError.Code)
	expect.Equal(t, err.Message.Error(), anyError.Message.Error())
}

func Test_GivenSigninExecute_WhenAuthenticateSuccess_ThenEnsureCallsBuildTokenWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, authService, sessionService, _ := fixtures.NewSut()
	authService.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.StrictEqual(t, sessionService.BuildTokenParam, authService.SuccessResult)
}

func Test_GivenSigninExecute_WhenBuildTokenInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.Equal(t, sessionService.BuildTokenCallsCount, 1)
}

func Test_GivenSigninExecute_WhenBuildTokenError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()
	sessionService.DefineBuildTokenError()

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.NotNil(t, err)
	expect.Equal(t, err.Code, appCode.CODE_SERVER_ERROR)
	expect.NotNil(t, err.Message)
}

func Test_GivenSigninExecute_WhenBuildTokenSuccess_ThenEnsureCallsSaveSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, authService, sessionService, _ := fixtures.NewSut()
	authService.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.StrictEqual(t, sessionService.SaveSessionParam, authService.SuccessResult)
}

func Test_GivenSigninExecute_WhenSaveSessionInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.Equal(t, sessionService.SaveSessionCallsCount, 1)
}

func Test_GivenSigninExecute_WhenSaveSessionError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, _, sessionService, _ := fixtures.NewSut()
	sessionService.DefineSaveSessionError()

	// Act
	_, err := sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.NotNil(t, err)
	expect.Equal(t, err.Code, appCode.CODE_SERVER_ERROR)
	expect.NotNil(t, err.Message)
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperWithCorrectParameters(t *testing.T) {
	// Arrange
	sut, authService, _, mapper := fixtures.NewSut()
	authService.DefineAuthenticate()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.StrictEqual(t, mapper.Params["session"], authService.SuccessResult)
}

func Test_GivenSigninExecute_WhenSaveSessionSuccess_ThenEnsureCallsMapperOnce(t *testing.T) {
	// Arrange
	sut, _, _, mapper := fixtures.NewSut()

	// Act
	sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.Equal(t, mapper.CallsCount, 1)
}

func Test_GivenSigninExecute_WhenMapperInvoke_ThenEnsureReturnOutputWithSession(t *testing.T) {
	// Arrange
	sut, authService, _, mapper := fixtures.NewSut()
	mapper.ForceReturn = false
	authService.DefineAuthenticate()

	// Act
	result, _ := sut.Execute(fixtures.GetValidInput())

	// Assert
	expect.NotNil(t, result)
	expect.Equal(t, result.Token, authService.SuccessResult.GetToken())
	expect.Equal(t, result.Session.Id, authService.SuccessResult.GetID())
	expect.Equal(t, result.Session.FirstName, authService.SuccessResult.GetFirstName())
	expect.Equal(t, result.Session.IsAdmin, authService.SuccessResult.GetIsAdmin())
}
