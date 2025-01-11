package session_service_test

import (
	fixture "getfund-api-v2/test/internal/shared/service/session_service/session_service_fixture"
	"strings"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSaveSession_WhenInputEmpty_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputEmpty())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("save-session: parameter cannot be null or empty")
}

func Test_GivenSaveSession_WhenInvalidInput_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()

	// Act
	_, err := sut.SaveSession(fixture.GetSaveSessionInputInvalid())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("save-session: parameter invalid")
}

func Test_GivenSaveSession_WhenValidInput_ThenEnsureCallsRedisSetWithCorrectParams(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedSession := fixture.GetSaveSessionInputValid()
	expectedValue := strings.Split(expectedSession, "@")

	// Act
	sut.SaveSession(expectedSession)

	// Assert
	verify.Should(t, spies.RedisCacheSpy.Params["Set:key"]).Be(expectedValue[0])
	verify.Should(t, spies.RedisCacheSpy.Params["Set:value"]).Be(expectedValue[1])
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
	sut, _ := fixture.NewSut()
	input := fixture.GetSaveSessionInputValid()
	expectedResult := strings.Split(input, "@")

	// Act
	result, _ := sut.SaveSession(input)

	// Assert
	verify.Should(t, result).Be(expectedResult[0])
}
