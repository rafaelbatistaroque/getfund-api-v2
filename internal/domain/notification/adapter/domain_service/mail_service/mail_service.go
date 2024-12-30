package mail_service

import (
	inputvalidation "getfund-api-v2/pkg/input_validation"

	"gopkg.in/gomail.v2"
)

type MailService interface {
	SendMail(to, subject, content string, replyTo []string) error
}

type mailService struct {
	params  inputvalidation.InputValidation
	message *gomail.Message
}

func New(message *gomail.Message) MailService {
	return &mailService{
		message: message,
	}
}

func (ms *mailService) SendMail(to, subject, content string, replyTo []string) error {
	ms.params.Required(to, "To")
	ms.params.Required(subject, "Subject")
	ms.params.Required(content, "Content")
	if ms.params.IsInvalid() {
		return ms.params.GetErrors()
	}

	ms.message.SetHeader("To", to)
	ms.message.SetHeader("Subject", subject)

	if len(replyTo) > 0 {
		ms.message.SetHeader("Reply-To", replyTo...)
	}

	return nil
}
