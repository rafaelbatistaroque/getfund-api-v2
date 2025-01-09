package auth_user_repository_test

import (
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/repository/auth_repository_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"

	"github.com/google/uuid"
)

func Test_GivenGetAuthenticatedUserByUsername_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	fixture.AddUser(db, 1, false)

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("invalid-username")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenInactivedUser_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	expectedId := uuid.NewString()
	username := uuid.NewString()
	db.Create(&fixture.FakeUser{ID: expectedId, Username: username, IsActive: 0})

	// Act
	_, err := sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenQuerySuccess_ThenEnsureReturnUserFound(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	expectedId := uuid.NewString()
	username := uuid.NewString()
	db.Create(&fixture.FakeUser{ID: expectedId, Username: username, IsActive: 1})

	// Act
	authenticatedUser, _ := sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, authenticatedUser.Id).Be(expectedId)
}

func Test_GivenUpdatePassword_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	fixture.AddUser(db, 1, true)

	// Act
	err := sut.UpdatePassword("invalid-id", "invalid-password")

	// Assert
	verify.Should(t, err).NotNil()
}
