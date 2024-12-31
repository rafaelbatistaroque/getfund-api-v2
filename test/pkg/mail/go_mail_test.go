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
