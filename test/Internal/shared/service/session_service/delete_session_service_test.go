package session_service_test

import (
	fixture "getfund-api-v2/test/internal/shared/service/session_service/session_service_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
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

func Test_GivenDeleteSession_WhenValidInput_ThenEnsureCallCacheDeleteWithCorrectParameter(t *testing.T) {
	// Arrange
	validInput := fixture.GetDeleteSessionInputValid()
	sut, _, _, redisSpy := fixture.NewSut()

	// Act
	sut.DeleteSession(validInput)

	// Assert
	verify.Should(t, redisSpy.Params["Delete:key"]).Be(validInput)
}

func Test_GivenDeleteSession_WhenCacheDeleteInvoked_ThenEnsureCallsOnce(t *testing.T) {
	// Arrange
	sut, _, _, redisSpy := fixture.NewSut()

	// Act
	sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, redisSpy.CallsCount["Delete"]).Be(1)
}

func Test_GivenDeleteSession_WhenCacheDeleteError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _, redisSpy := fixture.NewSut()
	redisSpy.DefineCacheDeleteError()

	// Act
	err := sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, redisSpy.ErrorResult["Delete"]).Be(err)
}

func Test_GivenDeleteSession_WhenCacheDeleteSuccess_ThenEnsureReturnNullError(t *testing.T) {
	// Arrange
	sut, _, _, _ := fixture.NewSut()

	// Act
	err := sut.DeleteSession(fixture.GetDeleteSessionInputValid())

	// Assert
	verify.Should(t, err).Nil()
}
