package session_service_proxy_test

import (
	"bytes"
	fixture "getfund-api-v2/test/internal/shared/service/proxy/session_service_proxy/session_service_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSaveSession_WhenEntry_ThenEnsureCallEncryptWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInputValue := fixture.GetSaveSessionInputValid()

	// Act
	sut.SaveSession(expectedInputValue)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["Encrypt:input"]).Be(expectedInputValue)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["Encrypt:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenSaveSession_WhenEncryptInvoked_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["Encrypt"]).Be(1)
}

func Test_GivenSaveSession_WhenEncryptSuccess_ThenEnsureCallHashAndMergeWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineEncryptSuccess()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(spies.HasherSpy.SuccessResult["Encrypt"])
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenSaveSession_WhenHashAndMergeInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashAndMerge"]).Be(1)
}

func Test_GivenSaveSession_WhenHashAndMergeSuccess_ThenEnsureCallSaveSessionWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineEncryptSuccess()
	spies.HasherSpy.DefineHashAndMergeSuccess("fake-token")
	expectedValueConcat := spies.HasherSpy.SuccessResult["HashAndMerge"].(string) + "@" + spies.HasherSpy.SuccessResult["Encrypt"].(string)

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.SessionSpy.Params["SaveSession:session"]).Be(expectedValueConcat)
}

func Test_GivenSaveSession_WhenSaveSessionInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.SessionSpy.CallsCount["SaveSession"]).Be(1)
}

func Test_GivenSaveSession_WhenSaveSessionError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineSaveSessionError()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, err).Be(spies.SessionSpy.ErrorResult["SaveSession"])
}

func Test_GivenSaveSession_WhenSaveSessionSuccessThenEnsureReturnSuccess(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.SessionSpy.DefineSaveSessionSuccess()

	// Act
	result, _ := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, result).Be(spies.SessionSpy.SuccessResult["SaveSession"])
}
