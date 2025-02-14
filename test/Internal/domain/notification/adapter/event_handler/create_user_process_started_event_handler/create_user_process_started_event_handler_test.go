package create_user_process_started_event_handler_test

import (
	fixture "getfund-api-v2/test/internal/domain/notification/adapter/event_handler/create_user_process_started_event_handler/create_user_process_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallGetCache(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetInvalidCreateUserProcessStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(0)
}
