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

func Test_GivenCreateUser_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.CreateUser(fixture.GetEmptyActivationUserDto())

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenCreateUser_WhenUserCreatedSuccess_ThenEnsureReturnUserDtoFilled(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	expectedUserCreated := fixture.GetFilledActivationUserDto()

	// Act
	result, err := sut.CreateUser(expectedUserCreated)

	// Assert
	userSaved := &schema.User{}
	db.Where("username = ?", expectedUserCreated.Username).First(userSaved)

	verify.Should(t, err).Nil()
	verify.Should(t, userSaved.ID).Be(uint(result.Id))
	verify.Should(t, userSaved.FirstName).Be(expectedUserCreated.FirstName)
	verify.Should(t, userSaved.LastName).Be(expectedUserCreated.LastName)
	verify.Should(t, userSaved.Password).Be(expectedUserCreated.Password)
	verify.Should(t, userSaved.IsActive).Be(expectedUserCreated.IsActive)
	verify.Should(t, userSaved.IsAdmin).Be(expectedUserCreated.IsAdmin)
	verify.Should(t, userSaved.Username).Be(expectedUserCreated.Username)
	verify.Should(t, userSaved.Username).Be(expectedUserCreated.Username)
	verify.Should(t, userSaved.CreatedAt).Be(expectedUserCreated.CreatedAt)
	verify.Should(t, userSaved.UpdatedAt).Be(expectedUserCreated.UpdatedAt)
}

func Test_GivenUserExists_WhenQueryError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	currentDb, _ := db.DB()
	currentDb.Close()

	// Act
	_, err := sut.UserExists("invalid-username")

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenUserExists_WhenNotFound_ThenEnsureReturnNuloValue(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	result, err := sut.UserExists("non-existent-username")

	// Assert
	verify.Should(t, result).Nil()
	verify.Should(t, err).Nil()
}

func Test_GivenUserExists_WhenFound_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	username := uuid.NewString()
	db.Create(&schema.User{Username: username})

	// Act
	userFound, err := sut.UserExists(username)

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, userFound).NotNil()
	verify.Should(t, userFound.Id).Be(1)
}
