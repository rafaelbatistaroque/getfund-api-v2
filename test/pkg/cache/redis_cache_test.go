package cache

import (
	"getfund-api-v2/internal/pkg/verify"
	fixture "getfund-api-v2/test/pkg/cache/rediscachefixture"
	"testing"
	"time"

	"github.com/google/uuid"
)

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
	time.Sleep(2 * time.Second)

	// Assert
	content, err := sut.Get(uniqueKey)
	verify.Should(t, err).NotNil()
	verify.Should(t, content).Be("")
}
