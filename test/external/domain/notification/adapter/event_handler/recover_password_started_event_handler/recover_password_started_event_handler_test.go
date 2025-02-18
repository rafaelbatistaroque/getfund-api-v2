package recover_password_eventhandler_test

import (
	fixture "getfund-api-v2/test/external/domain/notification/adapter/event_handler/recover_password_started_event_handler/recover_password_started_event_handler_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"

	"github.com/google/uuid"
)

func Test_GivenHandler_WhenPayloadParseError_ThenEnsureNeverCallGetCache(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act & Assert
	sut.Handle(fixture.GetInvalidRecoverPasswordStartedEvent())

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(0)
}

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedPayload := uuid.NewString()

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(expectedPayload))

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedPayload)
}

func Test_GivenHandler_WhenGetCacheInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(uuid.NewString()))

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenHandler_WhenCacheError_ThenEnsureNeverCallUsecaseExecute(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(uuid.NewString()))

	//Assert
	verify.Should(t, spies.SendRecoverPasswordMailUsecaseSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenUnmarshalError_ThenEnsureNeverCallUsecaseExecute(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccessWithValue("invalid-serialized-json")

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(uuid.NewString()))

	// Assert
	verify.Should(t, spies.SendRecoverPasswordMailUsecaseSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenUnmarshalSuccess_ThenEnsureCallUsecaseWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccess(fixture.GetValidCacheData())
	spies.SendRecoverPasswordMailUsecaseSpy.DefineSendRecoverPasswordMailUsecaseSuccess()
	expectedInput := fixture.GetSendRecoverPasswordMailInput()

	// Act
	sut.Handle(fixture.GetValidRecoverPasswordStartedEvent(uuid.NewString()))

	//Assert
	input := spies.SendRecoverPasswordMailUsecaseSpy.Params["Execute:input"]
	verify.Should(t, input.FirstName).Be(expectedInput.FirstName)
	verify.Should(t, input.Username).Be(expectedInput.Username)
	verify.Should(t, input.RecoveryLink).Be(expectedInput.RecoveryLink)
}
