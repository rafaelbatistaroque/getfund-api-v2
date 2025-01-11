package session_service_proxy_test

import (
	fixture "getfund-api-v2/test/internal/shared/service/proxy/session_service_proxy/session_service_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenDeleteSession_WhenEntry_ThenEnsureCallDeleteSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	validInput := fixture.GetDeleteSessionInputValid()

	// Act
	sut.DeleteSession(validInput)

	// Assert
	verify.Should(t, spies.SessionSpy.Params["DeleteSession:token"]).Be(validInput)
}

func Test_GivenDeleteSession_WhenDeleteSessionInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, spies.SessionSpy.CallsCount["DeleteSession"]).Be(1)
}

func Test_GivenDeleteSession_WhenDeleteSessionError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	err := sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, spies.SessionSpy.ErrorResult["Delete"]).Be(err)
}

func Test_GivenDeleteSession_WhenDeleteSessionSuccess_ThenEnsureReturnNullError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	err := sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, err).Nil()
}
