package send_activation_account_mail_application_test

import (
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/notification/core/usecase/send_activation_account_mail/send_activation_account_mail_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenFirstNameEmpty_ThenEnsureReturnUnprocessableError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "FirstName"))
}

func Test_GivenExecute_WhenEmailEmpty_ThenEnsureReturnUnprocessableError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Email"))
}

func Test_GivenExecute_WhenEmailInvalid_ThenEnsureReturnUnprocessableError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_EMAIL_INVALID.Error(), "Email"))
}

func Test_GivenExecute_WhenActivationLinkEmpty_ThenEnsureReturnUnprocessableError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "ActivationLink"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallGetActivationAccountTemplateOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.TemplateFileSpy.CallsCount["GetActivationAccountTemplate"]).Be(1)
}

func Test_GivenExecute_WhenGetActivationAccountTemplateError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.TemplateFileSpy.DefineGetActivationAccountTemplateError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.TemplateFileSpy.ErrorResult["GetActivationAccountTemplate"])
}
