package template_file_test

import (
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/port/template_file/template_file_fixture"
	"testing"
)

func Test_GivenGetRecoveryPasswordTemplate_WhenTemplateNotFound_ThenEnsureReturnCorrectError(t *testing.T) {
	// Arrange
	sut, _ := fixture.NewSUT()

	// Act
	_, err := sut.GetRecoveryPasswordTemplate()

	// Assert
	verify.Should(t, err.Error()).Be("template does not exist")
}
