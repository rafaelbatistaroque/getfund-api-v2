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

func Test_GivenExecute_WhenCacheError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.CacheSpy.ErrorResult["Get"])
}

func Test_GivenExecute_WhenUnmarshalError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.CacheSpy.DefineCacheGetSuccessWithValue("invalid-serialized-json")

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to unmarshal data")
}

func Test_GivenExecute_WhenGetRecoveryPasswordTemplateError_ThenEnsureReturnApplicationErrorWithServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSUT()
	spies.CacheSpy.DefineCacheGetSuccessWithValue("{}")
	spies.TemplateFileSpy.DefineGetRecoveryPasswordTemplateError()

	// Act
	_, err := sut.Execute(fixture.GetValidInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.TemplateFileSpy.ErrorResult["GetRecoveryPasswordTemplate"])
}
