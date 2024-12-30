package mail_service_test

import (
	"getfund-api-v2/internal/domain/notification/adapter/domain_service/mail_service"
	"getfund-api-v2/pkg/verify"
	"testing"
)

func Test_GivenSendMail_WhenInvalidToParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut := mail_service.New()

	// Act
	err := sut.SendMail("", "", "", nil)

	// Assert
	verify.Should(t, err.Error()).Be("parameter To cannot be null or empty")
}
