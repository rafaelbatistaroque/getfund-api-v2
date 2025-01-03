package session_service_test

import (
	"bytes"
	fixture "getfund-api-v2/test/internal/shared/service/session_service/session_service_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSaveSession_WhenInvalidInput_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixture.NewSut()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputInvalid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("save-session: parameter cannot be null or empty")
}

func Test_GivenSaveSession_WhenEncryptInvoked_ThenEnsureCallWithCorrectParams(t *testing.T) {
	// Arrange
	sut, hasherSpy, settingsSpy, _ := fixture.NewSut()
	expectedInputValue := fixture.GetSaveSessionInputValid()

	// Act
	sut.SaveSession(expectedInputValue)

	// Assert
	verify.Should(t, hasherSpy.Params["Encrypt:input"]).Be(expectedInputValue)
	verify.Should(t, bytes.Equal(hasherSpy.Params["Encrypt:secretKey"].([]byte), settingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenSaveSession_WhenEncryptInvoked_ThenEnsureCallOnce(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, _ := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, hasherSpy.CallsCount["Encrypt"]).Be(1)
}

func Test_GivenSaveSession_WhenEncryptSuccess_ThenEnsureCallHashAndMergeWithCorrectParams(t *testing.T) {
	// Arrange
	sut, hasherSpy, settingsSpy, _ := fixture.NewSut()
	hasherSpy.DefineEncryptSuccess()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, hasherSpy.Params["HashAndMerge:input"]).Be(hasherSpy.SuccessResult["Encrypt"])
	verify.Should(t, bytes.Equal(hasherSpy.Params["HashAndMerge:serverSalt"].([]byte), settingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenSaveSession_WhenHashAndMergeInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, _ := fixture.NewSut()
	hasherSpy.DefineEncryptSuccess()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, hasherSpy.CallsCount["HashAndMerge"]).Be(1)
}

func Test_GivenSaveSession_WhenHashAndMergeSuccess_ThenEnsureCallsRedisSetWithCorrectParams(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, redisSpy := fixture.NewSut()
	hasherSpy.DefineEncryptSuccess()
	hasherSpy.DefineHashAndMergeSuccess("any-token")

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, redisSpy.Params["Set:key"]).Be(hasherSpy.SuccessResult["HashAndMerge"])
	verify.Should(t, redisSpy.Params["Set:value"]).Be(hasherSpy.SuccessResult["Encrypt"])
	verify.Should(t, redisSpy.Params["Set:time"]).Be(24 * time.Hour)
}

func Test_GivenSaveSession_WhenRedisSetInvoke_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, _, redisSpy := fixture.NewSut()

	// Act
	sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, redisSpy.CallsCount["Set"]).Be(1)
}

func Test_GivenSaveSession_WhenRedisSetError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, redisSpy := fixture.NewSut()
	redisSpy.DefineCacheSetError()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, err).Be(redisSpy.ErrorResult["Set"])
}

func Test_GivenSaveSession_WhenRedisSetSuccess_ThenEnsureReturnToken(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, _ := fixture.NewSut()
	hasherSpy.DefineHashAndMergeSuccess("any-token")

	// Act
	result, _ := sut.SaveSession(fixture.GetSaveSessionInputValid())

	// Assert
	verify.Should(t, result).Be(hasherSpy.SuccessResult["HashAndMerge"])
}
