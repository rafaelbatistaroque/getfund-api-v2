package send_recover_password_mail_application_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail/application"
	"getfund-api-v2/internal/shared/result_app"
	inputvalidation "getfund-api-v2/pkg/input_validation"
	"getfund-api-v2/pkg/verify"
	"testing"
)

func Test_GivenExecute_WhenInvalidInput_ThenEnsureReturnApplicationErrorWithBadRequestError(t *testing.T) {
	// Arrange
	invalidInput := &send_recover_password_mail.Input{}
	sut := send_recover_password_mail_application.New()

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Be(fmt.Sprintf(inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "KeyCache"))
}
