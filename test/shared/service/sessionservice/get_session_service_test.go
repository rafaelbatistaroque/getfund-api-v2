package sessionservice

import (
	"getfund-api-v2/internal/pkg/verify"
	fixture "getfund-api-v2/test/shared/service/sessionservice/sessionservicefixture"
	"testing"
)

func Test_GivenGetSession_WhenInvalidInput_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixture.NewSut()

	// Act
	_, err := sut.GetSession(fixture.GetGetSessionInputInvalid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("get-session: parameter cannot be null or empty")
}

func Test_GivenGetSession_WhenValidInput_ThenEnsureCallCacheGetWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedParameter := fixture.GetGetSessionInputValid()
	sut, _, _, cacheSpy := fixture.NewSut()

	// Act
	sut.GetSession(expectedParameter)

	// Assert
	verify.Should(t, cacheSpy.Params["Get:key"]).Be(expectedParameter)
}

func Test_GivenGetSession_WhenCacheGetInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, _, cacheSpy := fixture.NewSut()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, cacheSpy.CallsCount["Get"]).Be(1)
}

func Test_GivenGetSession_WhenCacheGetError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, cacheSpy := fixture.NewSut()
	cacheSpy.DefineCacheGetError()

	// Act
	_, err := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err).Be(cacheSpy.ErrorResult["Get"])
}

func Test_GivenGetSession_WhenCacheGetSuccess_ThenEnsureCallDecryptMergedWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, hasherSpy, settingsSpy, cacheSpy := fixture.NewSut()
	cacheSpy.DefineCacheGetSuccess()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, hasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(cacheSpy.SuccessResult["Get"])
	verify.Should(t, hasherSpy.Params["DecryptMerged:secretKey"]).Be(settingsSpy.GetSecretKey())
}

func Test_GivenGetSession_WhenDecryptMergedInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, cacheSpy := fixture.NewSut()
	cacheSpy.DefineCacheGetSuccess()

	// Act
	sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, hasherSpy.CallsCount["DecryptMerged"]).Be(1)
}

func Test_GivenGetSession_WhenDecryptMergedSuccess_ThenEnsureReturnSession(t *testing.T) {
	// Arrange
	sut, hasherSpy, _, cacheSpy := fixture.NewSut()
	cacheSpy.DefineCacheGetSuccess()
	hasherSpy.DefineDecryptMergedSuccess("valid-session")

	// Act
	result, _ := sut.GetSession(fixture.GetGetSessionInputValid())

	// Assert
	verify.Should(t, result).Be(hasherSpy.SuccessResult["DecryptMerged"])
}
