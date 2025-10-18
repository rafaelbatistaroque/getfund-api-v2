package mail

import (
	"github.com/rafaelbatistaroque/validation"

	"gopkg.in/gomail.v2"
)

var (
	_TO           = "To"
	_SUBJECT      = "Subject"
	_CONTENT      = "Content"
	_REPLY_TO     = "Reply-To"
	_CONTENT_TYPE = "text/html"
)

// Contract defines the interface for a mail sending service.
// It provides a high-level method for sending emails.
type Contract interface {
	// SendMail sends an email with the specified details.
	// The content should be an HTML string.
	// It validates that 'to', 'subject', and 'content' are not empty.
	SendMail(to, subject, content string, replyTo []string) error
}

type customMail struct {
	rules validation.Rule
	*gomail.Message
	*gomail.Dialer
}

// New creates a new mail service instance.
// It takes a gomail.Message and a gomail.Dialer as dependencies.
func New(message *gomail.Message, dialer *gomail.Dialer) Contract {
	return &customMail{
		Message: message,
		Dialer:  dialer,
	}
}

func (ms *customMail) SendMail(to, subject, content string, replyTo []string) error {
	ms.rules.
		ApplyRules(to, _TO, &validation.RequiredRule{}).
		ApplyRules(subject, _SUBJECT, &validation.RequiredRule{}).
		ApplyRules(content, _CONTENT, &validation.RequiredRule{})

	if ms.rules.IsInvalid() {
		return ms.rules.GetErrors()
	}

	ms.Message.SetHeader(_TO, to)
	ms.Message.SetHeader(_SUBJECT, subject)

	if len(replyTo) > 0 {
		ms.Message.SetHeader(_REPLY_TO, replyTo...)
	}

	ms.Message.SetBody(_CONTENT_TYPE, content)

	ms.Dialer.DialAndSend(ms.Message)

	return nil
}
