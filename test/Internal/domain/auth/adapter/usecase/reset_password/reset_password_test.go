package reset_password_test

import (
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/usecase/reset_password/reset_password_fixture"
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
