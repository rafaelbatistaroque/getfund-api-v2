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
