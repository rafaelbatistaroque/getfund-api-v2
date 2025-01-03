package template_file_service_test

import (
	fixture "getfund-api-v2/test/internal/domain/notification/port/service/template_file_service/template_file_service_fixture"
	"os"
	"testing"

	"github.com/rafaelbatistaroque/verify"
)

func Test_GivenGetRecoveryPasswordTemplate_WhenTemplateNotFound_ThenEnsureReturnCorrectError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetRecoveryPasswordTemplate()

	// Assert
	verify.Should(t, err.Error()).Be("template does not exist")
}

func Test_GivenGetRecoveryPasswordTemplate_WhenTemplateFound_ThenEnsureReturnCorrectTemplate(t *testing.T) {
	// Arrange
	templateDir, _ := os.Getwd()
	filePath := templateDir + "/recovery_password_template.html"
	templateContent := "fake-template-content"
	file, _ := os.Create(filePath)
	file.WriteString(templateContent)

	defer os.Remove(filePath)
	defer file.Close()

	sut, spies := fixture.NewSUT()
	spies.SettingsSpy.SetTemplateDir(templateDir)

	// Act
	result, _ := sut.GetRecoveryPasswordTemplate()

	// Assert
	verify.Should(t, result).Be(templateContent)
}
