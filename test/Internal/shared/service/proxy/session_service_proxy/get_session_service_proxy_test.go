package session_service_proxy_test

import (
	fixture "getfund-api-v2/test/internal/shared/service/proxy/session_service_proxy/session_service_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetSession_WhenEntry_ThenEnsureCallGetSessionWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedParameter := fixture.GetGetSessionInputValid()

	// Act
	sut.GetSession(expectedParameter)

	// Assert
	verify.Should(t, spies.SessionSpy.Params["GetSession:token"]).Be(expectedParameter)
}

func Test_GivenGetSession_WhenGetSessionInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.SessionSpy.CallsCount["GetSession"]).Be(1)
}

func Test_GivenGetSession_WhenGetSessionError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineGetSessionError()

	// Act
	_, err := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err).Be(spies.SessionSpy.ErrorResult["GetSession"])
}

func Test_GivenGetSession_WhenGetSessionSuccess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineGetSessionSuccess()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(spies.SessionSpy.SuccessResult["GetSession"])
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:secretKey"]).Be(spies.SettingsSpy.GetSecretKey())
}

func Test_GivenGetSession_WhenDecryptMergedInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineGetSessionSuccess()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["DecryptMerged"]).Be(1)
}

func Test_GivenGetSession_WhenDecryptMergedSuccess_ThenEnsureReturnSession(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineGetSessionSuccess()
	spies.HasherSpy.DefineDecryptMergedSuccess("valid-session")

	// Act
	result, _ := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, result).Be(spies.HasherSpy.SuccessResult["DecryptMerged"])
}
