package cache

import (
	"getfund-api-v2/internal/pkg/verify"
	fixture "getfund-api-v2/test/pkg/cache/rediscachefixture"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_GivenCache_WhenSetSuccess_ThenEnsureSetCorrectValues(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	uniqueKey := uuid.NewString()
	expectedValue := uuid.NewString()
	defer sut.Close()

	// Act
	sut.Set(uniqueKey, expectedValue, 1*time.Second)

	// Assert
	content, _ := sut.Get(uniqueKey)
	verify.Should(t, content).Be(expectedValue)
}

func Test_GivenCache_WhenRigisterExpired_ThenEnsureReturnEmpty(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
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
