package reset_password_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/dto"
	shared_error "getfund-api-v2/internal/shared/error"
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
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
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
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "RecoveryCode", 64))
}

func Test_GivenExecute_WhenInputRecoveryKeyEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithRecoveryKeyEmpty())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "RecoveryKey"))
}

func Test_GivenExecute_WhenInputRecoveryKeyInvalid_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithRecoveryKeyInvalid())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_INVALID.Error(), "RecoveryKey"))
}

func Test_GivenExecute_WhenInputPasswordEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithPasswordEmpty())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordLowerThan8Character_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidLengthPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_LENGHT_INVALID.Error(), "Password", 8))
}

func Test_GivenExecute_WhenInputPasswordMissingUpperCaseCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidUpperPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_UPPER_CHARACTER.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordMissingLowerCaseCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidLowerPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_LOWER_CHARACTER.Error(), "Password"))
}

func Test_GivenExecute_WhenInputPasswordMissingDigitCharacter_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithInvalidDigitPassword())
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
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
	verify.Should(t, err.Code).Be(shared_error.NOT_FOUND_CODE)
	verify.Should(t, err.Message.Error()).Be("recovery code not found")
}

func Test_GivenExecute_WhenExecuteFinished_ThenEnsureDeleteCacheInOrderWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	validInput := fixture.GetInput()
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Delete:key"]).Be(validInput.RecoveryKey)
	verify.Should(t, spies.CacheSpy.InvokeOrder[0]).Be("Get")
	verify.Should(t, spies.CacheSpy.InvokeOrder[1]).Be("Delete")
}

func Test_GivenExecute_WhenExecuteFinished_ThenEnsureCallDeleteCacheOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

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
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error on get recovery password data")
}

func Test_GivenExecute_WhenGetCacheSuccess_ThenEnsureCallGetAuthenticatedUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(`{"username":"fake-username"}`)
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	expectedParam := fixture.GetForgetPasswordFromGetSuccessCache(spies.CacheSpy)

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetAuthenticatedUserByUsername:username"]).Be(expectedParam.Username)
}

func Test_GivenExecute_WhenGetAuthenticatedUserByUsernameInvoked_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetAuthenticatedUserByUsername"]).Be(1)
}

func Test_GivenExecute_WhenGetAuthenticatedUserByUsernameError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.NOT_FOUND_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["GetAuthenticatedUserByUsername"])
}

func Test_GivenExecute_WhenGetAuthenticatedUserByUsernameInvoked_ThenEnsureCallUpdatePasswordWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	expectedParamPassword := fixture.GetInput()

	// Act
	sut.Execute(expectedParamPassword)

	// Assert
	authenticatedUser := spies.RepoSpy.SuccessResult["GetAuthenticatedUserByUsername"].(*dto.AuthenticatedUserDto)
	verify.Should(t, spies.RepoSpy.Params["UpdatePassword:id"]).Be(authenticatedUser.Id)
	verify.Should(t, spies.RepoSpy.Params["UpdatePassword:value"]).Be(expectedParamPassword.Password)
}

func Test_GivenExecute_WhenUpdatePasswordInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UpdatePassword"]).Be(1)
}

func Test_GivenExecute_WhenUpdatePasswordError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()
	spies.RepoSpy.DefineUpdatePasswordError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["UpdatePassword"])
}

func Test_GivenExecute_WhenResetPasswordSuccess_ThenEnsureReturnAppropriateSuccessMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess("")
	spies.RepoSpy.DefineGetAuthenticatedUserByUsernameSuccess()

	// Act
	result, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, result.Message).Be("password updated")
}
