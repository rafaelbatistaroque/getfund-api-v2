package session_service_test

import (
	"bytes"
	fixture "getfund-api-v2/test/internal/domain/auth/core/domain_service/session_service/session_service_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify/v2"
)

func Test_GivenSaveSession_WhenInputNull_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputNull())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("save-session: session cannot be null or empty")
}

func Test_GivenSaveSession_WhenMarshalSuccess_ThenEnsureCallEncryptWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["Encrypt:input"]).Be(fixture.GetSaveSessionInputValidSerialized())
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["Encrypt:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenSaveSession_WhenEncryptInvoked_ThenEnsureCallsOnce(t *testing.T) {
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
	expectedSession := fixture.GetSaveSessionInputValid()

	// Act
	sut.SaveSession(expectedSession)

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

func Test_GivenSaveSession_WhenHashAndMergeSuccess_ThenEnsureCallRedisSetWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineEncryptSuccess()
	spies.HasherSpy.DefineHashAndMergeSuccess("fake-token")

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.RedisCacheSpy.Params["Set:key"]).Be(spies.HasherSpy.SuccessResult["HashAndMerge"])
	verify.Should(t, spies.RedisCacheSpy.Params["Set:value"]).Be(spies.HasherSpy.SuccessResult["Encrypt"])
	verify.Should(t, spies.RedisCacheSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenSaveSession_WhenRedisSetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, spies.RedisCacheSpy.CallsCount["Set"]).Be(1)
}

func Test_GivenSaveSession_WhenRedisSetError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RedisCacheSpy.DefineCacheSetError()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, err).Be(spies.RedisCacheSpy.ErrorResult["Set"])
}

func Test_GivenSaveSession_WhenRedisSetSuccess_ThenEnsureReturnToken(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashAndMergeSuccess("fake-token")

	// Act
	result, _ := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, result).Be(spies.HasherSpy.SuccessResult["HashAndMerge"])
}
