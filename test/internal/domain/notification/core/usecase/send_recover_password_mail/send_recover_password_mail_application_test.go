package send_recover_password_mail_application_test

import (
	"fmt"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/notification/core/usecase/send_recover_password_mail/send_recover_password_mail_fixture"
	"testing"

	inputvalidation "github.com/rafaelbatistaroque/validation"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenInvalidInput_ThenEnsureReturnApplicationErrorWithBadRequestError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(inputvalidation.Err_PARAMETER_NOT_EMPTY.Error(), "Username"))
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(inputvalidation.Err_PARAMETER_NOT_EMPTY.Error(), "FirstName"))
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(inputvalidation.Err_PARAMETER_NOT_EMPTY.Error(), "RecoveryLink"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallGetRecoveryPasswordTemplateOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.TemplateFileSpy.CallsCount["GetRecoveryPasswordTemplate"]).Be(1)
}

func Test_GivenExecute_WhenGetRecoveryPasswordTemplateInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.TemplateFileSpy.CallsCount["GetRecoveryPasswordTemplate"]).Be(1)
}

func Test_GivenExecute_WhenGetRecoveryPasswordTemplateError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.TemplateFileSpy.DefineGetRecoveryPasswordTemplateError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.TemplateFileSpy.ErrorResult["GetRecoveryPasswordTemplate"])
}

func Test_GivenExecute_WhenGotTemplateAndRecoveryPasswordModel_ThenEnsureSendMailWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	validInput := fixture.GetValidInput()
	spies.TemplateFileSpy.DefineGetRecoveryPasswordTemplateSuccess()
	templateReplaced := spies.TemplateFileSpy.GetRecoveryPasswordTemplateReplaced(validInput.FirstName, validInput.RecoveryLink)

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.MailSpy.Params["SendMail:to"]).Be(validInput.Username)
	verify.Should(t, spies.MailSpy.Params["SendMail:subject"]).Be("Password Recovery")
	verify.Should(t, spies.MailSpy.Params["SendMail:content"]).Be(templateReplaced)
}

func Test_GivenExecute_WhenSendMailInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()

	// Act
	sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, spies.MailSpy.CallsCount["SendMail"]).Be(1)
}

func Test_GivenExecute_WhenSendMailError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.MailSpy.DefineSendMailError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.MailSpy.ErrorResult["SendMail"])
}

func Test_GivenExecute_WhenSuccess_ThenEnsureReturnOutputWithCorrectMessage(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	result, _ := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, result.Messagem).Be("Email sent successfully")
}
