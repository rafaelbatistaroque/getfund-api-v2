package auth_user_repository_test

import (
	"getfund-api-v2/pkg/db/schema"
	fixture "getfund-api-v2/test/internal/domain/auth/adapter/repository/auth_repository_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"

	"github.com/google/uuid"
)

func Test_GivenGetAuthenticatedUserByUsername_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.GetAuthenticatedUserByUsername("invalid-username")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenInactivedUser_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	username := uuid.NewString()
	db.Create(&schema.User{Username: username, IsActive: false})

	// Act
	_, err := sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetAuthenticatedUserByUsername_WhenQuerySuccess_ThenEnsureReturnUserFound(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	username := uuid.NewString()
	db.Create(&schema.User{Username: username, IsActive: true})

	// Act
	authenticatedUser, _ := sut.GetAuthenticatedUserByUsername(username)

	// Assert
	verify.Should(t, authenticatedUser.Id).Be(1)
}

func Test_GivenUpdatePassword_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	err := sut.UpdatePassword(1, "")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenUpdatePassword_WhenSuccess_ThenEnsureNull(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	newPassword := uuid.NewString()
	db.Create(&schema.User{Password: uuid.NewString()})

	// Act
	result := sut.UpdatePassword(1, newPassword)

	// Assert
	user := &schema.User{}
	db.Where("id = ?", 1).First(user)

	verify.Should(t, result).Nil()
	verify.Should(t, user.Password).Be(newPassword)
}
