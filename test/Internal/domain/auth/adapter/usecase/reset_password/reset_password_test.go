package reset_password_test

import (
	"fmt"
	reset_password_application "getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password/application"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/usecase/reset_password/reset_password_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenInputRecoveryCodeEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithRecoveryCodeEmpty())
	sut := reset_password_application.New()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "RecoveryCode"))
}

func Test_GivenExecute_WhenInputRecoveryCodeLengthNotExactly64_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidOptions := []fixture.Option{fixture.WithRecoveryCodeEmpty(), fixture.WithRecoveryCodeInvalidLength()}
	invalidInput := fixture.GetInput(invalidOptions...)
	sut := reset_password_application.New()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "RecoveryCode", 64))
}

func Test_GivenExecute_WhenInputPasswordEmpty_ThenEnsureReturnErrorFromValidate(t *testing.T) {
	// Arrange
	invalidInput := fixture.GetInput(fixture.WithPasswordEmpty())
	sut := reset_password_application.New()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message).Be(fmt.Errorf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}
