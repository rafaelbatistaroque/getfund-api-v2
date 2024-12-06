package signinmapper_test

import (
	"bytes"
	"encoding/json"
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/auth/main/signinmapper/signinmapperfixture"
	"testing"

	"github.com/google/uuid"
)

func Test_GivenSigninMapper_WhenToOutput_ThenEnsureCorrectMapToSigninOutput(t *testing.T) {
	// Arrange
	expectedToken := "fake-token"
	sut, expectedResult, _, _ := fixture.NewSut()

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
	sut, toSerialized, _, _ := fixture.NewSut()
	expectedResult, _ := json.Marshal(toSerialized)

	// Act
	result, _ := sut.SessionToString(toSerialized)

	// Assert
	verify.Should(t, result).Be(string(expectedResult))
}

func Test_GivenSigninMapper_WhenToSessionModelMap_ThenEnsureCallDecryptWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, _, hasherSpy, settingsSpy := fixture.NewSut()
	user := fixture.GetUserModel()

	// Act
	sut.ToSessionModel(user)

	// Assert
	verify.Should(t, hasherSpy.Params["DecryptMerged:mergedEncryptedData"]).Be(user.FirstName)
	verify.Should(t, bytes.Equal(hasherSpy.Params["DecryptMerged:secretKey"].([]byte), settingsSpy.GetSecretKey())).BeTrue()
}

func Test_GivenSigninMapper_WhenToSessionModelMapped_ThenEnsureReturnSessionModel(t *testing.T) {
	// Arrange
	sut, _, hasherSpy, _ := fixture.NewSut()
	user := fixture.GetUserModel()
	hasherSpy.DefineDecryptMergedSuccess(uuid.NewString())

	// Act
	result := sut.ToSessionModel(user)

	// Assert
	verify.Should(t, result.ID).Be(user.Id)
	verify.Should(t, result.IsAdmin).Be(user.IsAdmin)
	verify.Should(t, result.FirstName).Be(hasherSpy.SuccessResult["DecryptMerged"])
}
