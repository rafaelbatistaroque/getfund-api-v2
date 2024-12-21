package recover_password_eventhandler_test

import (
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/port/event_handler/recover_password_started_event_handler/recover_password_started_event_handler_fixture"
	"testing"

	"github.com/google/uuid"
)

func Test_GivenHandler_WhenUnmarshalError_ThenEnsurePanic(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act & Assert
	verify.Should(t, nil).Panic(func() { sut.Handle(fixture.GetInvalidRecoverPasswordStartedEvent()) })
}

func Test_GivenHandler_WhenUnmarshalSuccess_ThenEnsureCAllUSecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedPayload := uuid.NewString()
	usecaseInput := fixture.GetSendRecoverPasswordMailInput(expectedPayload)
	sut, usecaseSpy := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(expectedPayload))

	//Assert
	verify.Should(t, usecaseSpy.Params["Execute:input"]).Be(usecaseInput)
}
