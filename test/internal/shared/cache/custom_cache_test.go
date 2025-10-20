package custom_cache_test

import (
	fixture "getfund-api-v2/test/internal/shared/cache/custom_cache_fixture"
	"testing"
	"time"

	"github.com/rafaelbatistaroque/verify/v2"

	"github.com/google/uuid"
)

func Test_GivenCacheSet_WhenMarshalError_ThenEnsureSetCorrectValues(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	invalidValue := map[any]any{123: "invalid-value"}
	defer sut.Close()

	// Act
	err := sut.Set(uuid.NewString(), invalidValue, 1*time.Second)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenCacheSet_WhenSuccess_ThenEnsureSetCorrectValues(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	uniqueKey := uuid.NewString()
	expectedValue := uuid.NewString()
	defer sut.Close()

	// Act
	sut.Set(uniqueKey, expectedValue, 1*time.Second)

	// Assert
	content, _ := sut.Get(uniqueKey)
	verify.Should(t, content).Be(expectedValue)
}

func Test_GivenCacheDelete_WhenError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, cancelContext := fixture.NewSut()
	cancelContext()
	defer sut.Close()

	// Act
	err := sut.Delete(uuid.NewString())

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenCacheDelete_WhenSuccess_ThenEnsureReturnNull(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	uniqueKey := uuid.NewString()
	sut.Set(uniqueKey, uniqueKey, 2*time.Second)
	defer sut.Close()

	// Act
	err := sut.Delete(uniqueKey)

	// Assert
	verify.Should(t, err).Nil()
}

func Test_GivenCache_WhenRigisterExpired_ThenEnsureReturnEmpty(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	uniqueKey := uuid.NewString()
	expectedValue := uuid.NewString()
	defer sut.Close()

	// Act
	sut.Set(uniqueKey, expectedValue, 1*time.Second)
	time.Sleep(1500 * time.Millisecond)

	// Assert
	content, err := sut.Get(uniqueKey)
	verify.Should(t, err).NotNil()
	verify.Should(t, content).Be("")
}
