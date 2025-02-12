package signin_mapper_test

import (
	fixture "getfund-api-v2/test/internal/domain/auth/main/signin_mapper/signin_mapper_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenSigninMapper_WhenToOutput_ThenEnsureCorrectMapToSigninOutput(t *testing.T) {
	// Arrange
	expectedToken := "fake-token"
	sut, expectedResult := fixture.NewSut()

	// Act
	result := sut.ToOutput(expectedToken, expectedResult)

	// Assert
	verify.Should(t, result).NotNil()
	verify.Should(t, result.Token).Be(expectedToken)
	verify.Should(t, result.Session.ID).Be(expectedResult.ID)
	verify.Should(t, result.Session.FirstName).Be(expectedResult.FirstName)
	verify.Should(t, result.Session.IsAdmin).Be(expectedResult.IsAdmin)
}

func Test_GivenSigninMapper_WhenToSessionModelMapped_ThenEnsureReturnSessionModel(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	authenticatedUser := fixture.GetauthenticatedUser()

	// Act
	result := sut.ToSessionModel(authenticatedUser)

	// Assert
	verify.Should(t, result.ID).Be(authenticatedUser.Id)
	verify.Should(t, result.FirstName).Be(authenticatedUser.FirstName)
	verify.Should(t, result.IsAdmin).Be(authenticatedUser.IsAdmin)
}
