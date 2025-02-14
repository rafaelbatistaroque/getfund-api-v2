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

func Test_GivenHandler_WhenPayloadParseSuccess_ThenEnsureCallGetCacheWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedPayload := "user_activation_fake_activation_code"

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessStartedEvent(expectedPayload))

	//Assert
	verify.Should(t, spies.CacheSpy.Params["Get:key"]).Be(expectedPayload)
}

func Test_GivenHandler_WhenCacheGetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessStartedEvent(""))

	//Assert
	verify.Should(t, spies.CacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenHandler_WhenCacheError_ThenEnsureNeverCallUsecaseExecute(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetError()

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessStartedEvent(""))

	//Assert
	verify.Should(t, spies.SendActivationAccountMailSpy.CallsCount["Execute"]).Be(0)
}

func Test_GivenHandler_WhenUnmarshalError_ThenEnsureNeverCallUsecaseExecute(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.CacheSpy.DefineCacheGetSuccessWithValue("invalid-serialized-json")

	// Act
	sut.Handle(fixture.GetValidCreateUserProcessStartedEvent(""))

	// Assert
	verify.Should(t, spies.SendActivationAccountMailSpy.CallsCount["Execute"]).Be(0)
}
