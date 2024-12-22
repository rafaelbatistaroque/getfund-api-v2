package template_file_test

import (
	"getfund-api-v2/internal/domain/notification/port/template_file"
	"getfund-api-v2/pkg/verify"
	"getfund-api-v2/test/helper/settings_spy"
	"testing"
)

func Test_GivenGetRecoveryPasswordTemplate_WhenTemplateNotFound_ThenEnsureReturnCorrectError(t *testing.T) {
	// Arrange
	settingsSpy := settings_spy.New()
	sut := template_file.New(settingsSpy)

	// Act
	_, err := sut.GetRecoveryPasswordTemplate()

	// Assert
	verify.Should(t, err.Error()).Be("template does not exist")
}
