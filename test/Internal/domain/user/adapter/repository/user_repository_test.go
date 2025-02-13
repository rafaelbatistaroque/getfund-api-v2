package user_repository_test

import (
	"getfund-api-v2/pkg/db/schema"
	fixture "getfund-api-v2/test/internal/domain/user/adapter/repository/user_repository_fixture"
	"testing"

	"github.com/google/uuid"
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

func Test_GivenUserExistsByUsername_WhenFound_ThenEnsureReturnSuccess(t *testing.T) {
	// Arrange
	sut, db := fixture.NewSUT()
	username := uuid.NewString()
	db.Create(&schema.User{Username: username})

	// Act
	userFound, err := sut.UserExistsByUsername(username)

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, userFound).NotNil()
	verify.Should(t, userFound.Id).Be(1)
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
	db.Where("username = ?", expectedUserCreated.Email).First(userSaved)

	verify.Should(t, err).Nil()
	verify.Should(t, userSaved.ID).Be(uint(result.Id))
	verify.Should(t, userSaved.FirstName).Be(expectedUserCreated.FirstName)
	verify.Should(t, userSaved.LastName).Be(expectedUserCreated.LastName)
	verify.Should(t, userSaved.CountryID).Be(uint(expectedUserCreated.CountryId))
	verify.Should(t, userSaved.UserCategoryID).Be(uint(expectedUserCreated.UserCategoryId))
	verify.Should(t, userSaved.Password).Be(expectedUserCreated.Password)
	verify.Should(t, userSaved.Gender).Be(expectedUserCreated.Gender)
	verify.Should(t, userSaved.IsActive).Be(expectedUserCreated.IsActive)
	verify.Should(t, userSaved.IsAdmin).Be(expectedUserCreated.IsAdmin)
	verify.Should(t, userSaved.MainSocialNetwork).Be(expectedUserCreated.MainSocialNetwork)
	verify.Should(t, userSaved.RegisteredUrl).Be(expectedUserCreated.RegisteredUrl)
	verify.Should(t, userSaved.Email).Be(expectedUserCreated.Email)
	verify.Should(t, userSaved.Username).Be(expectedUserCreated.Username)
	verify.Should(t, userSaved.CreatedAt).NotNil()
}
