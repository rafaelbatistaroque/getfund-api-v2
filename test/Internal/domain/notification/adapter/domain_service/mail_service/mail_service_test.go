package mail_service_test

import (
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/adapter/domain_service/mail_service/mail_service_fixture"
	"testing"
)

func Test_GivenSendMail_WhenInvalidToParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut := fixture.NewSUT()

	// Act
	err := sut.SendMail(fixture.GetFakeEmailParams(fixture.WithoutTo()))

	// Assert
	verify.Should(t, err.Error()).Be("parameter To cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidSubjectParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut := fixture.NewSUT()

	// Act
	err := sut.SendMail(fixture.GetFakeEmailParams(fixture.WithoutSubject()))

	// Assert
	verify.Should(t, err.Error()).Be("parameter Subject cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidContentParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut := fixture.NewSUT()

	// Act
	err := sut.SendMail(fixture.GetFakeEmailParams(fixture.WithoutContent()))

	// Assert
	verify.Should(t, err.Error()).Be("parameter Content cannot be null or empty")
}
