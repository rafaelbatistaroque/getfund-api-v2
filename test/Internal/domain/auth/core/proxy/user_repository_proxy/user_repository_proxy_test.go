package user_repository_proxy_test

import (
	"bytes"
	fixture "getfund-api-v2/test/internal/domain/auth/core/proxy/user_repository_proxy/user_repository_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetByUserName_WhenInit_ThenEnsureCallHashWithSaltWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedInputText := "fake-username"
	sut, spies := fixture.NewSut()

	// Act
	sut.GetByUserName(expectedInputText)

	// Assert
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInputText)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}

func Test_GivenGetByUserName_WhenInit_ThenEnsureCallHashWithSaltOnce(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()

	// Act
	sut.GetByUserName("fake-username")

	// Assert
	verify.Should(t, spies.HasherSpy.CallsCount["HashWithSalt"]).Be(1)
}

func Test_GivenGetByUserName_WhenHashWithSaltError_ThenEnsureReturnServerError(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltError()

	// Act
	_, err := sut.GetByUserName("fake-username")

	// Assert
	verify.Should(t, err).Be(spies.HasherSpy.ErrorResult["HashWithSalt"])
}

func Test_GivenGetByUserName_WhenHashWithSaltSuccess_ThenEnsureCallRepositoryGetByUserNameWithCorrectParameter(t *testing.T) {
	// Arrange
	expectedParameter := "fake-username-hashed"
	sut, spies := fixture.NewSut()
	spies.HasherSpy.DefineHashWithSaltSuccess(expectedParameter)

	// Act
	sut.GetByUserName("fake-username")

	// Assert
	verify.Should(t, spies.UserRepoSpy.Params["GetByUserName:username"]).Be(expectedParameter)
}
