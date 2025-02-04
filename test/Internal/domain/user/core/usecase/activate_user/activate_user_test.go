package activate_user_test

import (
	"fmt"
	"getfund-api-v2/internal/shared/result_app"
	fixture "getfund-api-v2/test/internal/domain/user/core/usecase/activate_user/activate_user_fixture"
	"testing"

	"github.com/rafaelbatistaroque/validation"
	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenExecute_WhenActivationCodeEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeEmpty := fixture.GetInput(fixture.WithEmptyActivationCode())

	// Act
	_, err := sut.Execute(inputWithActivationCodeEmpty)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_NOT_EMPTY.Error(), "ActivationCode"))
}

func Test_GivenExecute_WhenActivationCodeInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	inputWithActivationCodeInvalid := fixture.GetInput(fixture.WithInvalidActivationCodeLength())

	// Act
	_, err := sut.Execute(inputWithActivationCodeInvalid)

	// Assert
	verify.Should(t, err.Code).Be(result_app.UNAUTHORIZED_CODE)
	verify.Should(t, err.Message.Error()).Contain(fmt.Sprintf(validation.Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), "ActivationCode", 20))
}

func Test_GivenExecute_WhenInputValid_ThenEnsureCallCacheGetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetInput()
	expectedParam := "user_activation_" + validInput.ActivationCode

	// Act
	sut.Execute(validInput)

	// Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedParam)
}

func Test_GivenExecute_WhenCacheGetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Execute(fixture.GetInput())

	// Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}
