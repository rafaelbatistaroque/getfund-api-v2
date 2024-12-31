package go_mail_test

import (
	"getfund-api-v2/pkg/mail"
	"getfund-api-v2/pkg/verify"
	"getfund-api-v2/test/helper/settings_spy"
	"testing"
)

func Test_GivenNew_WhenReturnMessage_ThenEnsureSetFromKeyCorrectParameter(t *testing.T) {
	// Arrange
	settings_spy := settings_spy.New()

	// Act
	sutMessage, _ := mail.New(settings_spy)

	// Assert
	verify.Should(t, sutMessage.GetHeader("From")[0]).Be(settings_spy.GetSMTPFrom())
}

func Test_GivenNew_WhenReturnDialer_ThenEnsureSetCorrectHostParameter(t *testing.T) {
	// Arrange
	settings_spy := settings_spy.New()

	// Act
	_, sutDialer := mail.New(settings_spy)

	// Assert
	verify.Should(t, sutDialer.Host).Be(settings_spy.GetSMTPHost())
}

func Test_GivenNew_WhenReturnDialer_ThenEnsureSetCorrectPortParameter(t *testing.T) {
	// Arrange
	settings_spy := settings_spy.New()

	// Act
	_, sutDialer := mail.New(settings_spy)

	// Assert
	verify.Should(t, sutDialer.Port).Be(settings_spy.GetSMTPPort())
}

func Test_GivenNew_WhenReturnDialer_ThenEnsureSetCorrectUsernameParameter(t *testing.T) {
	// Arrange
	settings_spy := settings_spy.New()

	// Act
	_, sutDialer := mail.New(settings_spy)

	// Assert
	verify.Should(t, sutDialer.Username).Be(settings_spy.GetSMTPUsername())
}

func Test_GivenNew_WhenReturnDialer_ThenEnsureSetCorrectPasswordParameter(t *testing.T) {
	// Arrange
	settings_spy := settings_spy.New()

	// Act
	_, sutDialer := mail.New(settings_spy)

	// Assert
	verify.Should(t, sutDialer.Password).Be(settings_spy.GetSMTPPassword())
}
