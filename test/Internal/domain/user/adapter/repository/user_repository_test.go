package user_repository_test

import (
	fixture "getfund-api-v2/test/internal/domain/user/adapter/repository/user_repository_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenUserExistsByUsername_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.UserExistsByUsername("invalid-username")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenUserExistsByUsername_WhenNotFound_ThenEnsureReturnApropriateError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.UserExistsByUsername("non-existent-username")

	// Assert
	verify.Should(t, err.Error()).Be("user not found")
}
