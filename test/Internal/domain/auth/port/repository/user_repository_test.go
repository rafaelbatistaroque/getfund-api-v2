package auth_user_repository_test

import (
	fixture "getfund-api-v2/test/internal/domain/auth/port/repository/user_repository_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"

	"github.com/google/uuid"
)

func Test_GivenGetByUserName_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	fixture.AddUser(db, 1, false)

	// Act
	_, err := sut.GetByUserName("invalid-username")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetByUserName_WhenInactivedUser_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	expectedId := uuid.NewString()
	username := uuid.NewString()
	db.Create(&fixture.FakeUser{ID: expectedId, Username: username, IsActive: 0})

	// Act
	_, err := sut.GetByUserName(username)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetByUserName_WhenQuerySuccess_ThenEnsureReturnUserFound(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	expectedId := uuid.NewString()
	username := uuid.NewString()
	db.Create(&fixture.FakeUser{ID: expectedId, Username: username, IsActive: 1})

	// Act
	user, _ := sut.GetByUserName(username)

	// Assert
	verify.Should(t, user.Id).Be(expectedId)
}
