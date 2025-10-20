package mail_test

import (
	"testing"

	"getfund-api-v2/test/internal/shared/mail/mail_fixture"

	"github.com/rafaelbatistaroque/verify/v2"
)

func Test_GivenSendMail_WhenInvalidToParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams(mail_fixture.WithoutTo())

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Contain("To cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidSubjectParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams(mail_fixture.WithoutSubject())

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Contain("Subject cannot be null or empty")
}

func Test_GivenSendMail_WhenInvalidContentParam_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _ := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams(mail_fixture.WithoutContent())

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err.Error()).Contain("Content cannot be null or empty")
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetToKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	sut, fixture := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, fixture.MessageSpy.GetHeader("To")[0]).Be(params.To)
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetSubjectKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	sut, fixture := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, fixture.MessageSpy.GetHeader("Subject")[0]).Be(params.Subject)
}

func Test_GivenSendMail_WhenValidParams_ThenEnsureSetReplyToKeyAsCorrectParamsToSend(t *testing.T) {
	// Arrange
	sut, fixture := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams()

	// Act
	sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, fixture.MessageSpy.GetHeader("Reply-To")[0]).Be(params.ReplyTo[0])
}

func Test_GivenSendMail_WhenDialAndSendReturnsError_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, fixture := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams()
	fixture.DialerSpy.DefineError()

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err).NotNil()
	verify.Should(t, err.Error()).Be("fake-error")
}

func Test_GivenSendMail_WhenSuccess_ThenEnsureDialAndSendIsCalledOnce(t *testing.T) {
	// Arrange
	sut, fixture := mail_fixture.NewSut()
	params := mail_fixture.GetFakeEmailParams()
	fixture.DialerSpy.DefineSuccess()

	// Act
	err := sut.SendMail(params.GetParams())

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, fixture.DialerSpy.CallsCount["DialAndSend"]).Be(1)
}