package create_user_test

import (
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/core/usecase/create_user/create_user_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenCreateUserExecute_WhenInputFirstNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithFirstNameEmpty())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "FirstName"))
}

func Test_GivenCreateUserExecute_WhenInputFirstNameInvalidLength_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithFirstNameLengthInvalid())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_LENGHT_MAX_INVALID.Error(), "FirstName", 50))
}

func Test_GivenCreateUserExecute_WhenInputLastNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithLastNameEmpty())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "LastName"))
}
