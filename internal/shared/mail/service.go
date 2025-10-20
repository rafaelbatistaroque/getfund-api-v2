package mail

import (
	"github.com/rafaelbatistaroque/validation"
)

var (
	_TO           = "To"
	_SUBJECT      = "Subject"
	_CONTENT      = "Content"
	_REPLY_TO     = "Reply-To"
	_CONTENT_TYPE = "text/html"
)

// Message defines the interface for an email message.
type Message interface {
	SetHeader(key string, value ...string)
	SetBody(contentType string, body string)
}

// Dialer defines the interface for a mail dialer.
type Dialer interface {
	DialAndSend(m ...Message) error
}

// Contract defines the interface for a mail sending service.
type Contract interface {
	SendMail(to, subject, content string, replyTo []string) error
}

type mailService struct {
	rules validation.Rule
	message Message
	dialer  Dialer
}

// New creates a new mail service instance.
func New(message Message, dialer Dialer) Contract {
	return &mailService{
		message: message,
		dialer:  dialer,
	}
}

func (ms *mailService) SendMail(to, subject, content string, replyTo []string) error {
	ms.rules.
		ApplyRules(to, _TO, &validation.RequiredRule{}).
		ApplyRules(subject, _SUBJECT, &validation.RequiredRule{}).
		ApplyRules(content, _CONTENT, &validation.RequiredRule{})

	if ms.rules.IsInvalid() {
		return ms.rules.GetErrors()
	}

	ms.message.SetHeader(_TO, to)
	ms.message.SetHeader(_SUBJECT, subject)

	if len(replyTo) > 0 {
		ms.message.SetHeader(_REPLY_TO, replyTo...)
	}

	ms.message.SetBody(_CONTENT_TYPE, content)

	return ms.dialer.DialAndSend(ms.message)
}