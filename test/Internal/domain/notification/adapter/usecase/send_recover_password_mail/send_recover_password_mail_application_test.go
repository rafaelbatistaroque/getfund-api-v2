package send_recover_password_mail_application_test

import (
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	inputvalidation "getfund-api-v2/pkg/input_validation"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/adapter/usecase/send_recover_password_mail/send_recover_password_mail_fixture"
	"testing"
)

func Test_GivenExecute_WhenInvalidInput_ThenEnsureReturnApplicationErrorWithBadRequestError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.Execute(fixture.GetInvalidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Be(fmt.Sprintf(inputvalidation.Err_Msg_PARAMETER_NOT_EMPTY.Error(), "KeyCache"))
}

func Test_GivenExecute_WhenValidInput_ThenEnsureCallCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetValidInput()
	sut, spies := fixture.NewSUT()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(validInput.KeyCache)
}
