package mail_service

import (
	contract "getfund-api-v2/internal/domain/notification/adapter/contract"
	inputvalidation "getfund-api-v2/pkg/input_validation"

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
	params  inputvalidation.InputValidation
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
	ms.params.Required(to, _TO)
	ms.params.Required(subject, _SUBJECT)
	ms.params.Required(content, _CONTENT)
	if ms.params.IsInvalid() {
		return ms.params.GetErrors()
	}

	ms.message.SetHeader(_TO, to)
	ms.message.SetHeader(_SUBJECT, subject)

	if len(replyTo) > 0 {
		ms.message.SetHeader(_REPLY_TO, replyTo...)
	}

	ms.message.SetBody(_CONTENT_TYPE, content)

	ms.dialer.DialAndSend(ms.message)

	return nil
}
