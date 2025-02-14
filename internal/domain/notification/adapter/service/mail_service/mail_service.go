package mail_service

import (
	contract "getfund-api-v2/internal/domain/notification/core/contract"

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

type mailService struct {
	rules   validation.Rule
	message *gomail.Message
	dialer  *gomail.Dialer
}

func New(message *gomail.Message, dialer *gomail.Dialer) contract.MailService {
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

	//ms.dialer.DialAndSend(ms.message)

	return nil
}
