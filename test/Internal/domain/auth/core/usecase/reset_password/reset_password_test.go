package reset_password_test

import (
	"bytes"
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/core/usecase/reset_password/reset_password_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenInputRecoveryCodeEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithRecoveryCodeEmpty())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "RecoveryCode"))
}

func Test_GivenExecute_WhenInputRecoveryCodeLengthNotExactly64_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidOptions := []fixture.Option{fixture.WithRecoveryCodeEmpty(), fixture.WithRecoveryCodeInvalidLength()}
	invalidInput := fixture.GetInput(invalidOptions...)
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "RecoveryCode", 64))
}

func Test_GivenExecute_WhenInputPasswordEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithPasswordEmpty())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordLowerThan8Character_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidLengthPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_LENGHT_INVALID.Error(), "Password", 8))
}

func Test_GivenExecute_WhenInputPasswordMissingUpperCaseCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidUpperPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_UPPER_CHARACTER.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordMissingLowerCaseCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidLowerPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_LOWER_CHARACTER.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordMissingDigitCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidDigitPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_DIGIT_CHARACTER.Error(), "Password"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallGetCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetInput()
	expectedCacheParameter := "recovery_password_" + validInput.RecoveryCode
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedCacheParameter)
}

func Test_GivenExecute_WhenGetCacheInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenExecute_WhenGetCacheError_ThenEnsureReturnNotFoundWithErrorFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(spies.CacheSpy.ErrorResult["Get"])
}

func Test_GivenExecute_WhenExecuteFinished_ThenEnsureDeleteCacheInOrderWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetInput()
	expectedCacheParameter := "recovery_password_" + validInput.RecoveryCode
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Delete:key"]).Be(expectedCacheParameter)
	verify.Should(t, spies.CacheSpy.InvokeOrder[0]).Be("Get")
	verify.Should(t, spies.CacheSpy.InvokeOrder[1]).Be("Delete")
}

func Test_GivenExecute_WhenExecuteFinished_ThenEnsureCallDeleteCacheOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Delete"]).Be(1)
}

func Test_GivenExecute_WhenUnmarshalError_ThenEnsureReturnAppropriateError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(`{any-body}`)

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to unmarshal data")
}

func Test_GivenExecute_WhenGetCacheSuccess_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(`{"username":"fake-username"}`)
	expectedParam := fixture.GetForgetPasswordFromGetSuccessCache(spies.CacheSpy)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedParam.Username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenExecute_WhenHashWithSaltInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenExecute_WhenHashWithSaltError_ThenEnsureReturnInternalError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenExecute_WhenHashWithSaltSuccess_ThenEnsureCallGetByUserNameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedValue := "fake-username-hashed"
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.HasherSpy.DefineHashWithSaltSuccess(expectedValue)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.UserRepoSpy.Params["GetByUserName:username"]).Be(expectedValue)
}

func Test_GivenExecute_WhenGetByUserNameInvoked_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.UserRepoSpy.CallsCount["GetByUserName"]).Be(1)
}

func Test_GivenExecute_WhenGetByUserNameError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.UserRepoSpy.DefineGetByUserNameError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(spies.UserRepoSpy.ErrorResult["GetByUserName"])
}
