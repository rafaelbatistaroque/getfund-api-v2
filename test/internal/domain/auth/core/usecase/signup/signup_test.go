package signup_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup/event"
	shared_error "getfund-api-v2/internal/shared/error"
	fixture "getfund-api-v2/test/internal/domain/auth/core/usecase/signup/signup_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSignupExecute_WhenInputFirstNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyFirstName())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "FirstName"))
}

func Test_GivenSignupExecute_WhenInputFirstNameInvalidLength_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidFirstNameLength())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_LENGHT_MAX_INVALID.Error(), "FirstName", 50))
}

func Test_GivenSignupExecute_WhenInputLastNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyLastName())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "LastName"))
}

func Test_GivenSignupExecute_WhenInputLastNameInvalidLength_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidLastNameLength())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_LENGHT_MAX_INVALID.Error(), "LastName", 50))
}

func Test_GivenSignupExecute_WhenInputUsernameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyUsername())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Username"))
}

func Test_GivenSignupExecute_WhenInputUsernameInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidUsername())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_EMAIL_INVALID.Error(), "Username"))
}

func Test_GivenSignupExecute_WhenInputPasswordEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyPassword())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

func Test_GivenSignupExecute_WhenInputPasswordInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidPasswords := map[string]string{
		"fake_p":               fmt.Sprintf(validation.Err_PARAMETER_LENGHT_INVALID.Error(), "Password", 8),           //invalid min length
		"FAKE_STRONG_PASSWORD": fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_LOWER_CHARACTER.Error(), "Password"), //invalid required lower case
		"fake_strong_password": fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_UPPER_CHARACTER.Error(), "Password"), //invalid required upper case
		"fake_Strong_Password": fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_DIGIT_CHARACTER.Error(), "Password"), //invalid required upper case
	}

	for invalidPassword, messageError := range invalidPasswords {

		invalidInput := fixture.GetInput(fixture.WithInvalidPassword(invalidPassword))

		// Act
		_, err := sut.Execute(invalidInput)

		// Assert
		verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
		verify.Should(t, err.Message.Error()).Contain(messageError)
	}
}

func Test_GivenSignupExecute_WhenInputPasswordConfirmationEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyPasswordConfirmation())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "PasswordConfirmation"))
}

func Test_GivenSignupExecute_WhenInputPasswordConfirmationDifferentFromPassword_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidPasswordConfirmation("diff-password"))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_INVALID.Error(), "PasswordConfirmation"))
}

func Test_GivenSignupExecute_WhenInputCouponCodeInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidCouponCode("invalid"))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(shared_error.UNPROCESSABLE_CONTENT_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "CouponCode", 8))
}

func Test_GivenSignupExecute_WhenInputValid_ThenEnsureCallUserExistsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.RepoSpy.Params["UserExists:username"]).Be(validInput.Username)
}
func Test_GivenSignupExecute_WhenUserExistsInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["UserExists"]).Be(1)
}

func Test_GivenSignupExecute_WhenUserExistsError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineUserExistsError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["UserExists"])
}

func Test_GivenSignupExecute_WhenUserExistsFound_ThenEnsureReturnDuplicateEntryError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineUserExistsSuccessUserFound()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.DUPLICATED_ENTRY_CODE)
	verify.Should(t, err.Message.Error()).Be("user already exists")
}

func Test_GivenSignupExecute_WhenUserExistsNotFound_ThenEnsureCallGetRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineUserExistsSuccess()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(20)
}

func Test_GivenSignupExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenSignupExecute_WhenGetRandomCodeError_ThenEnsureReturnInternalServerErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to save user")
}

func Test_GivenSignupExecute_WhenGetRandomCodeSuccess_ThenEnsureCallSetCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	expectedKeyCache := "user_activation_" + spies.HasherSpy.SuccessResult["GetRandomCode"].(string)
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Set:key"]).Be(expectedKeyCache)
	verify.Should(t, spies.CacheSpy.Params["Set:value"]).Be(validInput)
	verify.Should(t, spies.CacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenSignupExecute_WhenCacheSetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Set"]).Be(1)
}

func Test_GivenSignupExecute_WhenCacheSetError_ThenEnsureReturnErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheSetError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(shared_error.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to save user")
}

func Test_GivenSignupExecute_WhenCacheSuccess_ThenEnsureCallEmitWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput(fixture.WithEmptyCouponCode())
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	activationCode := spies.HasherSpy.SuccessResult["GetRandomCode"].(string)
	payload := &event.SignupStartedPayload{
		FirstName:      validInput.FirstName,
		Email:          validInput.Username,
		ActivationCode: activationCode,
		ActivationLink: spies.SettingsSpy.GetBaseUrl() + "/user-activation/" + activationCode,
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:event"][0]).Be(&event.SignupStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:payload"][0]).Be(payload)
}

func Test_GivenSignupExecute_WhenEmitWithPayloadInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput(fixture.WithEmptyCouponCode()))

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(1)
}

func Test_GivenSignupExecute_WhenSuccess_ThenEnsureReturnSuccessWithAppropiateMessage(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	result, _ := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, result.Message).Be("user creation started")
}
