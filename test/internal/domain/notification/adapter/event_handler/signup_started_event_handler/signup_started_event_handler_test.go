package signup_process_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/notification/adapter/event_handler/signup_started_event_handler/signup_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallUsecase(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidSignupProcessStartedEvent())

	//Assert
	verify.Should(t, spies.SendActivationAccountMailSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenUnmarshalSuccess_ThenEnsureCallUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetSendActivationAccountMailInput()

	// Act
	sut.Handle(fixture.GetValidSignupProcessStartedEvent())

	//Assert
	inputParams := spies.SendActivationAccountMailSpy.Params["Execute:input"]
	verify.Should(t, inputParams.FirstName).Be(expectedInput.FirstName)
	verify.Should(t, inputParams.Email).Be(expectedInput.Email)
	verify.Should(t, inputParams.ActivationLink).Be(expectedInput.ActivationLink)
}

func Test_GivenHandler_WhenHandleSuccess_ThenEnsureCallUsecaseOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidSignupProcessStartedEvent())

	//Assert
	verify.Should(t, spies.SendActivationAccountMailSpy.CallsCount["Execute"]).Be(1)
}
