package user_respository_proxy_test

import (
	"bytes"
	fixture "getfund-api-v2/test/internal/domain/user/adapter/proxy/user_respository_proxy/user_respository_proxy_fixture"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenCreateUser_WhenInit_ThenEnsureCallHasherMethodsWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, spies := fixture.NewSut()
	expectedInput := fixture.GetFilledActivationUserDto()

	// Act
	sut.CreateUser(expectedInput)

	// Assert
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][0]).Be(expectedInput.FirstName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][0].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][1]).Be(expectedInput.LastName)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][1].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][2]).Be(expectedInput.Email)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][2].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][3]).Be(expectedInput.MainSocialNetwork)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][3].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.ParamsByCall["Encrypt:input"][4]).Be(expectedInput.RegisteredUrl)
	verify.Should(t, bytes.Equal(spies.HasherSpy.ParamsByCall["Encrypt:secretKey"][4].([]byte), spies.SettingsSpy.GetSecretKey())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashAndMerge:input"]).Be(expectedInput.Password)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashAndMerge:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
	verify.Should(t, spies.HasherSpy.Params["HashWithSalt:inputText"]).Be(expectedInput.Username)
	verify.Should(t, bytes.Equal(spies.HasherSpy.Params["HashWithSalt:serverSalt"].([]byte), spies.SettingsSpy.GetServerSalt())).BeTrue()
}
