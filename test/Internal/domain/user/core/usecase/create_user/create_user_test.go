package create_user_test

import (
	"fmt"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus/event"
	fixture "getfund-api-v2/test/internal/domain/user/core/usecase/create_user/create_user_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenCreateUserExecute_WhenInputFirstNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyFirstName())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "FirstName"))
}

func Test_GivenCreateUserExecute_WhenInputFirstNameInvalidLength_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidFirstNameLength())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_LENGHT_MAX_INVALID.Error(), "FirstName", 50))
}

func Test_GivenCreateUserExecute_WhenInputLastNameEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyLastName())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "LastName"))
}

func Test_GivenCreateUserExecute_WhenInputLastNameInvalidLength_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidLastNameLength())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_LENGHT_MAX_INVALID.Error(), "LastName", 50))
}

func Test_GivenCreateUserExecute_WhenInputEmailInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidEmail())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_EMAIL_INVALID.Error(), "Email"))
}

func Test_GivenCreateUserExecute_WhenInputGenderEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyGender())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Gender"))
}

func Test_GivenCreateUserExecute_WhenInputGenderInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidGender())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_WITHIN_LIST.Error(), "Gender", "[f m u nb]"))
}

func Test_GivenCreateUserExecute_WhenInputPasswordEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyPassword())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "Password"))
}

func Test_GivenCreateUserExecute_WhenInputPasswordInvalid_ThenEnsureReturnError(t *testing.T) {
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
		verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
		verify.Should(t, err.Message.Error()).Contain(messageError)
	}
}

func Test_GivenCreateUserExecute_WhenInputMainSocialNetworkEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyMainSocialNetwork())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "MainSocialNetwork"))
}

func Test_GivenCreateUserExecute_WhenInputMainSocialNetworkEInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	defaultError := fmt.Sprintf(validation.Err_PARAMETER_SHOULD_BE_SOCIAL_URL_INVALID.Error(), "MainSocialNetwork")
	sut, _ := fixture.NewSut()
	invalidMainNetworks := map[string]string{
		// Invalid protocol
		"http://example.com":         defaultError, // Protocol http is not allowed
		"ftp://example.com":          defaultError, // Protocol ftp is not allowed
		"file://example.com":         defaultError, // Protocol file is not allowed
		"mailto:user@example.com":    defaultError, // Protocol mailto is not allowed
		"data:text/plain;base64,...": defaultError, // Protocol data is not allowed

		// Invalid domain
		"https://-example.com": defaultError, // Domain starts with an invalid character (-)
		"https://example-.com": defaultError, // Domain ends with an invalid character (-)
		"https://example..com": defaultError, // Domain contains consecutive dots
		"https://example_com":  defaultError, // Underscore (_) is not allowed in the domain

		// Invalid path
		"https://example.com/pa th": defaultError, // Path contains a space
		"https://example.com/<tag>": defaultError, // Path contains invalid characters (< >)
		"https://example.com//path": defaultError, // Path contains consecutive slashes
		"https://example.com/|path": defaultError, // Path contains an invalid character (|)
	}

	for invalidMainNetwork, messageError := range invalidMainNetworks {

		invalidInput := fixture.GetInput(fixture.WithInvalidMainSocialNetwork(invalidMainNetwork))

		// Act
		_, err := sut.Execute(invalidInput)

		// Assert
		verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
		verify.Should(t, err.Message.Error()).Contain(messageError)
	}
}

func Test_GivenCreateUserExecute_WhenInputRegisteredUrlEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithEmptyRegisteredUrl())

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "RegisteredUrl"))
}

func Test_GivenCreateUserExecute_WhenInputCouponCodeInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidInput := fixture.GetInput(fixture.WithInvalidCouponCode("invalid"))

	// Act
	_, err := sut.Execute(invalidInput)

	// Assert
	verify.Should(t, err.Code).Be(result_app.BAD_REQUEST_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "CouponCode", 8))
}

func Test_GivenCreateUserExecute_WhenInputValid_ThenEnsureCallGetUserByUsernameWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.RepoSpy.Params["GetUserByUsername:username"]).Be(validInput.Email)
}
func Test_GivenCreateUserExecute_WhenGetUserByUsernameInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.RepoSpy.CallsCount["GetUserByUsername"]).Be(1)
}

func Test_GivenCreateUserExecute_WhenGetUserByUsernameError_ThenEnsureReturnNotFoundError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetUserByUsernameError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message).Be(spies.RepoSpy.ErrorResult["GetUserByUsername"])
}

func Test_GivenCreateUserExecute_WhenGetUserByUsernameFound_ThenEnsureReturnDuplicateEntryError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetUserByUsernameSuccess()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.DUPLICATED_ENTRY_CODE)
	verify.Should(t, err.Message.Error()).Be("user already exists")
}

func Test_GivenCreateUserExecute_WhenGetUserByUsernameNotFound_ThenEnsureCallGetRandomCodeWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RepoSpy.DefineGetUserByUsernameSuccessNotFound()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["GetRandomCode:length"]).Be(20)
}

func Test_GivenCreateUserExecute_WhenGetRandomCodeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["GetRandomCode"]).Be(1)
}

func Test_GivenCreateUserExecute_WhenGetRandomCodeError_ThenEnsureReturnInternalServerErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineGetRandomCodeError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to save user")
}

func Test_GivenCreateUserExecute_WhenGetRandomCodeSuccess_ThenEnsureCallSetCacheWithCorrectParameter(t *testing.T) {
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

func Test_GivenCreateUserExecute_WhenCacheSetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Set"]).Be(1)
}

func Test_GivenCreateUserExecute_WhenCacheSetError_ThenEnsureReturnErrorWithAppropriateMessage(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheSetError()

	// Act
	_, err := sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, err.Code).Be(result_app.SERVER_ERROR_CODE)
	verify.Should(t, err.Message.Error()).Be("error to save user")
}

func Test_GivenCreateUserExecute_WhenCacheSuccess_ThenEnsureCallEmitWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput(fixture.WithEmptyCouponCode())
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	payload := &user_dto.UserCriationStartedDto{
		ActivationCode: "user_activation_" + spies.HasherSpy.SuccessResult["GetRandomCode"].(string),
		FirstName:      validInput.FirstName,
		Email:          validInput.Email,
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:event"]).Be(&event.UserCriationStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:payload"]).Be(payload)
}

func Test_GivenCreateUserExecute_WhenEmitWithPayloadInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput(fixture.WithEmptyCouponCode()))

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(1)
}

func Test_GivenCreateUserExecute_WhenHasCouponCode_ThenEnsureCallEmitWithPayloadWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	spies.HasherSpy.DefineGetRandomCodeSuccess()
	payload := &user_dto.UserCriationWithCouponStardedDto{
		CouponCode:     validInput.CouponCode,
		ActivationCode: "user_activation_" + spies.HasherSpy.SuccessResult["GetRandomCode"].(string),
	}

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:event"]).Be(&event.UserCriationWithCouponCodeStartedEvent{})
	verify.Should(t, spies.BusSpy.Params["EmitWithPayload:payload"]).Be(payload)
}

func Test_GivenCreateUserExecute_WhenHasCouponCodeAndEmitWithPayloadInvoked_ThenEnsureCallsTwice(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.BusSpy.CallsCount["EmitWithPayload"]).Be(2)
}
