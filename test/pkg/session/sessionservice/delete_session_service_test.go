package sessionservice

import (
	"getfund-api-v2/internal/pkg/verify"
	fixture "getfund-api-v2/test/pkg/session/sessionservice/sessionservicefixture"
	"testing"
)

func Test_GivenDeleteSession_WhenInvalidInput_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixture.NewSut()

	// Act
	err := sut.DeleteSession(fixture.GetDeleteSessionInputInvalid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("delete-session: parameter cannot be null or empty")
}

func Test_GivenDeleteSession_WhenInputValid_ThenEnsureCallDeleteWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetDeleteSessionInputValid()
	sut, _, _, redisSpy := fixture.NewSut()

	// Act
	sut.DeleteSession(validInput)

	// Assert
	verify.Should(t, redisSpy.Params["Delete:key"]).Be(validInput)
}

func Test_GivenDeleteSession_WhenCacheInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, _, redisSpy := fixture.NewSut()

	// Act
	sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, redisSpy.CallsCount["Delete"]).Be(1)
}
