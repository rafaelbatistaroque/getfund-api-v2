package mail_service_test

import (
	"getfund-api-v2/pkg/verify"
	fixture "getfund-api-v2/test/internal/domain/notification/port/service/mail_service/mail_service_fixture"
	"testing"
)

func Test_GivenSendMail_WhenInvalidToParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams(fixture.WithoutTo())
	sut, _ := fixture.NewSUT()

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Be("parameter To cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidSubjectParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams(fixture.WithoutSubject())
	sut, _ := fixture.NewSUT()

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Be("parameter Subject cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidContentParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams(fixture.WithoutContent())
	sut, _ := fixture.NewSUT()

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Be("parameter Content cannot be null or empty")
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetToKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams()
	sut, dependences := fixture.NewSUT()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, dependences.Mail.GetHeader("To")[0]).Be(params.To)
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetSubjectKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams()
	sut, dependences := fixture.NewSUT()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, dependences.Mail.GetHeader("Subject")[0]).Be(params.Subject)
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetReplyToKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	params := fixture.GetFakeEmailParams()
	sut, dependences := fixture.NewSUT()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, dependences.Mail.GetHeader("Reply-To")[0]).Be(params.Attachments[0])
}
