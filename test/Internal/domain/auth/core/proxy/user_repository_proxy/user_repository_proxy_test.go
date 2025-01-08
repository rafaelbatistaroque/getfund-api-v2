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
