package signin_mapper_test

import (
	"encoding/json"
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
	verify.Should(t, result.Session.IsAdmin).Be(expectedResult.IsAdmin == 1)
}

func Test_GivenSigninMapper_WhenSessionToStringSuccess_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, toSerialized := fixture.NewSut()
	expectedResult, _ := json.Marshal(toSerialized)

	// Act
	result, _ := sut.SessionToString(toSerialized)

	// Assert
	verify.Should(t, result).Be(string(expectedResult))
}

func Test_GivenSigninMapper_WhenToSessionModelMapped_ThenEnsureReturnSessionModel(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSut()
	user := fixture.GetUserModel()

	// Act
	result := sut.ToSessionModel(user)

	// Assert
	verify.Should(t, result.ID).Be(user.Id)
	verify.Should(t, result.FirstName).Be(user.FirstName)
	verify.Should(t, result.IsAdmin).Be(user.IsAdmin)
}
