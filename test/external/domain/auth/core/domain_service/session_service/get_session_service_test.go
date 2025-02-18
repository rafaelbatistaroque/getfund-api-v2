package session_service_test

import (
	"bytes"
	fixture "getfund-api-v2/test/external/domain/auth/core/domain_service/session_service/session_service_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetSession_WhenInvalidInput_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.GetSession(fixture.GetGetSessionInputInvalid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("get-session: parameter cannot be null or empty")
}

func Test_GivenGetSession_WhenValidInput_ThenEnsureCallCacheGetWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedParameter := fixture.GetGetSessionInputValid()

	// Act
	sut.GetSession(expectedParameter)

	// Assert
	verify.Should(t, spies.RedisCacheSpy.Params["Get:key"]).Be(expectedParameter)
}

func Test_GivenGetSession_WhenCacheGetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.RedisCacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenGetSession_WhenCacheGetError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RedisCacheSpy.DefineCacheGetError()

	// Act
	_, err := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err).Be(spies.RedisCacheSpy.ErrorResult["Get"])
}

func Test_GivenGetSession_WhenCacheGetSuccess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RedisCacheSpy.DefineCacheGetSuccess("fake-session")

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(spies.RedisCacheSpy.SuccessResult["Get"])
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["DecryptMerged:secretKey"].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenGetSession_WhenDecryptMergedInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RedisCacheSpy.DefineCacheGetSuccess("fake-session")

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["DecryptMerged"]).Be(1)
}

func Test_GivenGetSession_WhenDecryptMergedISuccess_ThenEnsureCallsResultFrom(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.RedisCacheSpy.DefineCacheGetSuccess("fake-session")
	spies.HasherSpy.DefineDecryptMergedSuccess("any-value")

	// Act
	result, _ := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, result).Be(spies.HasherSpy.SuccessResult["DecryptMerged"])
}
