package recover_password_eventhandler_test

import (
	"getfund-api-v2/pkg/eventbus/event"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/port/event_handler/recover_password_event_handler/recover_password_event_handler_fixture"
	"testing"
)

func Test_GivenHandler_WhenUnmarshalError_ThenEnsurePanic(t *testing.T) {
	// Arrange
	invalidEvent := &event.RecoverPasswordStarted{}
	sut, _ := fixture.NewSut()

	// Act & Assert
	verify.Should(t, nil).Panic(func() { sut.Handle(invalidEvent) })
}
